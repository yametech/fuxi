/*
Copyright (c) 2018 The Helm Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helm

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/ghodss/yaml"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	helmrepo "k8s.io/helm/pkg/repo"
)

const (
	defaultTimeoutSeconds = 10
	additionalCAFile      = "/usr/local/share/ca-certificates/ca.crt"
)

type importChartFilesJob struct {
	Name         string
	Repo         *Repo
	ChartVersion ChartVersion
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var NetClient httpClient = &http.Client{}

func parseRepoURL(repoURL string) (*url.URL, error) {
	repoURL = strings.TrimSpace(repoURL)
	return url.ParseRequestURI(repoURL)
}

func init() {
	var err error
	NetClient, err = initNetClient(additionalCAFile)
	if err != nil {
		log.Fatal(err)
	}
}

type assetManager interface {
	Delete(repo Repo) error
	Sync(repo Repo, charts []Chart) error
	RepoAlreadyProcessed(repo Repo, checksum string) bool
	UpdateLastCheck(repoNamespace, repoName, checksum string, now time.Time) error
	Init() error
	Close() error
	InvalidateCache() error
	updateIcon(repo Repo, data []byte, contentType, ID string) error
	filesExist(repo Repo, chartFilesID, digest string) bool
	insertFiles(chartID string, files ChartFiles) error
}

//
//func newManager(config datastore.Config, kubeappsNamespace string) (assetManager, error) {
//	return newPGManager(config, kubeappsNamespace)
//}

var (
	version          = "devel"
	userAgentComment string
)

func userAgent() string {
	ua := "syncer/" + version
	if userAgentComment != "" {
		ua = fmt.Sprintf("%s (%s)", ua, userAgentComment)
	}
	return ua
}

func getSha256(src []byte) (string, error) {
	f := bytes.NewReader(src)
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func GetRepo(namespace, name, repoURL, authorizationHeader string) (*RepoInternal, []byte, error) {
	url, err := parseRepoURL(repoURL)
	if err != nil {
		log.WithFields(log.Fields{"url": repoURL}).WithError(err).Error("failed to parse URL")
		return nil, []byte{}, err
	}

	repoBytes, err := FetchRepoIndex(url.String(), authorizationHeader)
	if err != nil {
		return nil, []byte{}, err
	}

	repoChecksum, err := getSha256(repoBytes)
	if err != nil {
		return nil, []byte{}, err
	}

	return &RepoInternal{Namespace: namespace, Name: name, URL: url.String(), Checksum: repoChecksum, AuthorizationHeader: authorizationHeader}, repoBytes, nil
}

func FetchRepoIndex(url, authHeader string) ([]byte, error) {
	indexURL, err := parseRepoURL(url)
	if err != nil {
		log.WithFields(log.Fields{"url": url}).WithError(err).Error("failed to parse URL")
		return nil, err
	}
	indexURL.Path = path.Join(indexURL.Path, "index.yaml")
	req, err := http.NewRequest("GET", indexURL.String(), nil)
	if err != nil {
		log.WithFields(log.Fields{"url": req.URL.String()}).WithError(err).Error("could not build repo index request")
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent())
	if len(authHeader) > 0 {
		req.Header.Set("Authorization", authHeader)
	}

	res, err := NetClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		log.WithFields(log.Fields{"url": req.URL.String()}).WithError(err).Error("error requesting repo index")
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{"url": req.URL.String(), "status": res.StatusCode}).Error("error requesting repo index, are you sure this is a chart repository?")
		return nil, errors.New("repo index request failed")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ParseRepoIndex(body []byte) (*helmrepo.IndexFile, error) {
	var index helmrepo.IndexFile
	err := yaml.Unmarshal(body, &index)
	if err != nil {
		return nil, err
	}
	index.SortEntries()
	return &index, nil
}

func ChartsFromIndex(index *helmrepo.IndexFile, r *Repo) []Chart {
	var charts []Chart
	for _, entry := range index.Entries {
		if entry[0].GetDeprecated() {
			log.WithFields(log.Fields{"name": entry[0].GetName()}).Info("skipping deprecated chart")
			continue
		}
		charts = append(charts, newChart(entry, r))
	}
	sort.Slice(charts, func(i, j int) bool { return charts[i].ID < charts[j].ID })
	return charts
}

// Takes an entry from the index and constructs a database representation of the
// object.
func newChart(entry helmrepo.ChartVersions, r *Repo) Chart {
	var c Chart
	copier.Copy(&c, entry[0])
	copier.Copy(&c.ChartVersions, entry)
	c.Repo = r
	c.ID = fmt.Sprintf("%s/%s", r.Name, c.Name)
	c.Category = entry[0].Annotations["category"]
	return c
}

func extractFilesFromTarball(filenames []string, tarf *tar.Reader) (map[string]string, error) {
	ret := make(map[string]string)
	for {
		header, err := tarf.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return ret, err
		}

		for _, f := range filenames {
			if strings.EqualFold(header.Name, f) {
				var b bytes.Buffer
				io.Copy(&b, tarf)
				ret[f] = string(b.Bytes())
				break
			}
		}
	}
	return ret, nil
}

func chartTarballURL(r *RepoInternal, cv ChartVersion) string {
	source := cv.URLs[0]
	if _, err := parseRepoURL(source); err != nil {
		// If the chart URL is not absolute, join with repo URL. It's fine if the
		// URL we build here is invalid as we can catch this error when actually
		// making the request
		u, _ := url.Parse(r.URL)
		u.Path = path.Join(u.Path, source)
		return u.String()
	}
	return source
}

func initNetClient(additionalCA string) (*http.Client, error) {
	// Get the SystemCertPool, continue with an empty pool on error
	caCertPool, _ := x509.SystemCertPool()
	if caCertPool == nil {
		caCertPool = x509.NewCertPool()
	}

	// If additionalCA exists, load it
	if _, err := os.Stat(additionalCA); !os.IsNotExist(err) {
		certs, err := ioutil.ReadFile(additionalCA)
		if err != nil {
			return nil, fmt.Errorf("Failed to append %s to RootCAs: %v", additionalCA, err)
		}

		// Append our cert to the system pool
		if ok := caCertPool.AppendCertsFromPEM(certs); !ok {
			return nil, fmt.Errorf("Failed to append %s to RootCAs", additionalCA)
		}
	}

	// Return Transport for testing purposes
	return &http.Client{
		Timeout: time.Second * defaultTimeoutSeconds,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
			Proxy: http.ProxyFromEnvironment,
		},
	}, nil
}

type fileImporter struct {
	manager assetManager
}

func (f *fileImporter) fetchFiles(charts []Chart, r *RepoInternal) {
	// Process 10 charts at a time
	numWorkers := 10
	iconJobs := make(chan Chart, numWorkers)
	chartFilesJobs := make(chan importChartFilesJob, numWorkers)
	var wg sync.WaitGroup

	log.Debugf("starting %d workers", numWorkers)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go f.importWorker(&wg, iconJobs, chartFilesJobs, r)
	}

	// Enqueue jobs to process chart icons
	for _, c := range charts {
		iconJobs <- c
	}
	// Close the iconJobs channel to signal the worker pools to move on to the
	// chart files jobs
	close(iconJobs)

	// Iterate through the list of charts and enqueue the latest chart version to
	// be processed. Append the rest of the chart versions to a list to be
	// enqueued later
	var toEnqueue []importChartFilesJob
	for _, c := range charts {
		chartFilesJobs <- importChartFilesJob{c.Name, c.Repo, c.ChartVersions[0]}
		for _, cv := range c.ChartVersions[1:] {
			toEnqueue = append(toEnqueue, importChartFilesJob{c.Name, c.Repo, cv})
		}
	}

	// Enqueue all the remaining chart versions
	for _, cfj := range toEnqueue {
		chartFilesJobs <- cfj
	}
	// Close the chartFilesJobs channel to signal the worker pools that there are
	// no more jobs to process
	close(chartFilesJobs)

	// Wait for the worker pools to finish processing
	wg.Wait()
}

func (f *fileImporter) importWorker(wg *sync.WaitGroup, icons <-chan Chart, chartFiles <-chan importChartFilesJob, r *RepoInternal) {
	defer wg.Done()
	for c := range icons {
		log.WithFields(log.Fields{"name": c.Name}).Debug("importing icon")
		if err := f.fetchAndImportIcon(c, r); err != nil {
			log.WithFields(log.Fields{"name": c.Name}).WithError(err).Error("failed to import icon")
		}
	}
	for j := range chartFiles {
		log.WithFields(log.Fields{"name": j.Name, "version": j.ChartVersion.Version}).Debug("importing readme and values")
		if err := f.fetchAndImportFiles(j.Name, r, j.ChartVersion); err != nil {
			log.WithFields(log.Fields{"name": j.Name, "version": j.ChartVersion.Version}).WithError(err).Error("failed to import files")
		}
	}
}

func (f *fileImporter) fetchAndImportIcon(c Chart, r *RepoInternal) error {
	if c.Icon == "" {
		log.WithFields(log.Fields{"name": c.Name}).Info("icon not found")
		return nil
	}

	req, err := http.NewRequest("GET", c.Icon, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent())
	if len(r.AuthorizationHeader) > 0 {
		req.Header.Set("Authorization", r.AuthorizationHeader)
	}

	res, err := NetClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%d %s", res.StatusCode, c.Icon)
	}

	b := []byte{}
	contentType := ""
	var img image.Image
	// if the icon is in any other format try to convert it to PNG
	if strings.Contains(res.Header.Get("Content-Type"), "image/svg") {
		// if the icon is an SVG, it requires special processing
		icon, err := oksvg.ReadIconStream(res.Body)
		if err != nil {
			log.WithFields(log.Fields{"name": c.Name}).WithError(err).Error("failed to decode icon")
			return err
		}
		w, h := int(icon.ViewBox.W), int(icon.ViewBox.H)
		icon.SetTarget(0, 0, float64(w), float64(h))
		rgba := image.NewNRGBA(image.Rect(0, 0, w, h))
		icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)
		img = rgba
	} else {
		img, err = imaging.Decode(res.Body)
		if err != nil {
			log.WithFields(log.Fields{"name": c.Name}).WithError(err).Error("failed to decode icon")
			return err
		}
	}

	// TODO: make this configurable?
	resizedImg := imaging.Fit(img, 160, 160, imaging.Lanczos)
	var buf bytes.Buffer
	imaging.Encode(&buf, resizedImg, imaging.PNG)
	b = buf.Bytes()
	contentType = "image/png"
	_ = b
	_ = contentType
	//log.Println(b)
	//log.Println(contentType)
	return nil
	//return f.manager.updateIcon(Repo{Namespace: r.Namespace, Name: r.Name}, b, contentType, c.ID)
}

func (f *fileImporter) fetchAndImportFiles(name string, r *RepoInternal, cv ChartVersion) error {
	chartID := fmt.Sprintf("%s/%s", r.Name, name)
	chartFilesID := fmt.Sprintf("%s-%s", chartID, cv.Version)

	// Check if we already have indexed files for this chart version and digest
	//TODO:fixed should check it
	//if f.manager.filesExist(Repo{Namespace: r.Namespace, Name: r.Name}, chartFilesID, cv.Digest) {
	//	log.WithFields(log.Fields{"name": name, "version": cv.Version}).Debug("skipping existing files")
	//	return nil
	//}
	log.WithFields(log.Fields{"name": name, "version": cv.Version}).Debug("fetching files")

	url := chartTarballURL(r, cv)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent())
	if len(r.AuthorizationHeader) > 0 {
		req.Header.Set("Authorization", r.AuthorizationHeader)
	}

	res, err := NetClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// We read the whole chart into memory, this should be okay since the chart
	// tarball needs to be small enough to fit into a GRPC call (Tiller
	// requirement)
	gzf, err := gzip.NewReader(res.Body)
	if err != nil {
		return err
	}
	defer gzf.Close()

	tarf := tar.NewReader(gzf)

	readmeFileName := name + "/README.md"
	valuesFileName := name + "/values.yaml"
	schemaFileName := name + "/values.schema.json"
	filenames := []string{valuesFileName, readmeFileName, schemaFileName}

	files, err := extractFilesFromTarball(filenames, tarf)
	if err != nil {
		return err
	}

	chartFiles := ChartFiles{ID: chartFilesID, Repo: &Repo{Name: r.Name, Namespace: r.Namespace, URL: r.URL}, Digest: cv.Digest}
	if v, ok := files[readmeFileName]; ok {
		chartFiles.Readme = v
	} else {
		log.WithFields(log.Fields{"name": name, "version": cv.Version}).Info("README.md not found")
	}
	if v, ok := files[valuesFileName]; ok {
		chartFiles.Values = v
	} else {
		log.WithFields(log.Fields{"name": name, "version": cv.Version}).Info("values.yaml not found")
	}
	if v, ok := files[schemaFileName]; ok {
		chartFiles.Schema = v
	} else {
		log.WithFields(log.Fields{"name": name, "version": cv.Version}).Info("values.schema.json not found")
	}

	// inserts the chart files if not already indexed, or updates the existing
	// entry if digest has changed
	//return f.manager.insertFiles(chartID, chartFiles)
	return nil
}
