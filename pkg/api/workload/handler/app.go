package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ghodss/yaml"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	k8sjson "k8s.io/apimachinery/pkg/util/json"

	"k8s.io/apimachinery/pkg/runtime"

	cmdutil "k8s.io/kubectl/pkg/cmd/util"

	"helm.sh/helm/v3/pkg/action"

	"k8s.io/kubectl/pkg/scheme"

	"github.com/yametech/fuxi/pkg/api/common"
	"github.com/yametech/fuxi/pkg/app/helm"
	helmrepo "k8s.io/helm/pkg/repo"

	"github.com/gin-gonic/gin"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/releaseutil"
	helmtime "helm.sh/helm/v3/pkg/time"
	helmmeta "k8s.io/helm/pkg/proto/hapi/chart"
)

type HelmCharts struct {
	APIVersion  string    `json:"apiVersion"`
	AppVersion  string    `json:"appVersion"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	Digest      string    `json:"digest"`
	Home        string    `json:"home"`
	Icon        string    `json:"icon"`
	Keywords    []string  `json:"keywords"`
	Maintainers []*helmmeta.Maintainer
	Name        string   `json:"name"`
	Sources     []string `json:"sources"`
	Urls        []string `json:"urls"`
	Version     string   `json:"version"`
	Repo        string   `json:"repo"`
}

const (
	RepoName = "compass-harbor"
)

//ListCharts List Chart
func (w *WorkloadsAPI) ListCharts(g *gin.Context) {

	_, contentResult, err := helm.GetRepo("", "", w.HarborAddress, "")
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	index, err := helm.ParseRepoIndex(contentResult)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	for key, value := range index.Entries {
		v := value[0]
		cvs := helmrepo.ChartVersions{}
		cvs = append(cvs, v)
		index.Entries[key] = cvs
	}
	respMap := make(map[string]HelmCharts)
	for key, value := range index.Entries {
		chart := new(HelmCharts)
		chart.Name = value[0].Name
		chart.Repo = RepoName
		chart.APIVersion = value[0].ApiVersion
		chart.Version = value[0].Version
		chart.AppVersion = value[0].AppVersion
		chart.Created = value[0].Created
		chart.Description = value[0].Description
		chart.Digest = value[0].Digest
		chart.Home = value[0].Home
		chart.Icon = value[0].Icon
		chart.Keywords = value[0].Keywords
		chart.Maintainers = value[0].Maintainers
		chart.Sources = value[0].Sources
		respMap[key] = *chart
	}

	g.JSON(http.StatusOK, gin.H{RepoName: respMap})

}

type Chart struct {
	Readme   string        `json:"readme"`
	Versions []ChartDetail `json:"versions"`
}

type ChartDetail struct {
	Annotations  map[string]string `json:"annotations,omitempty"`
	APIVersion   string            `json:"apiVersion"`
	AppVersion   string            `json:"appVersion"`
	Created      time.Time         `json:"created"`
	Dependencies []*chart.Dependency
	Description  string   `json:"description"`
	Digest       string   `json:"digest"`
	Home         string   `json:"home"`
	Icon         string   `json:"icon"`
	Keywords     []string `json:"keywords"`
	Maintainers  []*helmmeta.Maintainer
	Name         string   `json:"name"`
	Sources      []string `json:"sources"`
	Urls         []string `json:"urls"`
	Version      string   `json:"version"`
	Repo         string   `json:"repo"`
}

//GetCharts get Chartss
func (w *WorkloadsAPI) GetCharts(g *gin.Context) {
	//only compass harbor repo
	//repo :=g.Param("repo")

	chartName := g.Param("chart")
	if len(chartName) == 0 {
		return
	}
	_, contentResult, err := helm.GetRepo("", "", w.HarborAddress, "")
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	index, err := helm.ParseRepoIndex(contentResult)
	chartDetail := new(ChartDetail)
	var chartDetailResp []ChartDetail
	for _, value := range index.Entries {
		if value[0].Name == chartName {
			for _, entry := range value {
				chartDetail.Annotations = entry.Annotations
				chartDetail.AppVersion = entry.AppVersion
				chartDetail.APIVersion = entry.ApiVersion
				chartDetail.Created = entry.Created
				chartDetail.Dependencies = []*chart.Dependency{}
				chartDetail.Description = entry.Description
				chartDetail.Digest = entry.Digest
				chartDetail.Home = entry.Home
				chartDetail.Icon = entry.Icon
				chartDetail.Keywords = entry.Keywords
				chartDetail.Maintainers = entry.Maintainers
				chartDetail.Name = entry.Name
				chartDetail.Sources = entry.Sources
				chartDetail.Urls = entry.URLs
				chartDetail.Version = entry.Version
				chartDetail.Repo = "compass-harbor"
				chartDetailResp = append(chartDetailResp, *chartDetail)
			}
		}
	}

	url := fmt.Sprintf("%s/charts/%s", w.HarborAddress, chartName+"-"+chartDetailResp[0].Version+".tgz")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := helm.NetClient.Do(req)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	defer resp.Body.Close()
	ch := new(Chart)
	chart, _ := loader.LoadArchive(resp.Body)
	for _, raw := range chart.Raw {
		if raw.Name == "README.md" {
			ch.Readme = string(raw.Data)
		}
	}
	ch.Versions = chartDetailResp
	g.JSON(http.StatusOK, ch)

}

//GetChartValues get chart values
func (w *WorkloadsAPI) GetChartValues(g *gin.Context) {
	//repo :=g.Param("repo")

	chartName := g.Param("chart")
	if len(chartName) == 0 {
		return
	}
	_, contentResult, err := helm.GetRepo("", "", w.HarborAddress, "")
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	index, err := helm.ParseRepoIndex(contentResult)
	chartDetail := new(ChartDetail)
	var chartDetailResp []ChartDetail
	for _, value := range index.Entries {
		if value[0].Name == chartName {
			for _, entry := range value {
				chartDetail.Annotations = entry.Annotations
				chartDetail.AppVersion = entry.AppVersion
				chartDetail.APIVersion = entry.ApiVersion
				chartDetail.Created = entry.Created
				chartDetail.Dependencies = []*chart.Dependency{}
				chartDetail.Description = entry.Description
				chartDetail.Digest = entry.Digest
				chartDetail.Home = entry.Home
				chartDetail.Icon = entry.Icon
				chartDetail.Keywords = entry.Keywords
				chartDetail.Maintainers = entry.Maintainers
				chartDetail.Name = entry.Name
				chartDetail.Sources = entry.Sources
				chartDetail.Urls = entry.URLs
				chartDetail.Version = entry.Version
				chartDetail.Repo = RepoName
				chartDetailResp = append(chartDetailResp, *chartDetail)
			}
		}
	}

	url := fmt.Sprintf("%s/charts/%s", w.HarborAddress, chartName+"-"+chartDetailResp[0].Version+".tgz")
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := helm.NetClient.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()
	ch := new(Chart)
	chart, _ := loader.LoadArchive(resp.Body)
	for _, raw := range chart.Raw {
		if raw.Name == "values.yaml" {
			ch.Readme = string(raw.Data)
		}
	}
	ch.Versions = chartDetailResp
	g.JSON(http.StatusOK, ch.Readme)

}

type InstallChart struct {
	Chart     string      `json:"chart"`
	Values    helm.Values `json:"values"`
	Name      string      `json:"name"`
	Namespace string      `json:"namespace"`
	Version   string      `json:"version"`
}

//InstallChart install Chart
func (w *WorkloadsAPI) InstallChart(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	obj := &InstallChart{}
	err = json.Unmarshal(rawData, &obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	url := fmt.Sprintf("%s/charts/%s", w.HarborAddress, strings.Split(obj.Chart, "/")[1]+"-"+obj.Version+".tgz")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := helm.NetClient.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()
	chart, err := loader.LoadArchive(resp.Body)

	val, err := yaml.Marshal(obj.Values)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	rel, err := helm.CreateRelease(w.ActionInstance(obj.Namespace), obj.Name, obj.Namespace, string(val), chart, nil)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"log": rel,
		"release": gin.H{
			"name": rel.Name,
		},
	})

}

type Release struct {
	Name       string        `json:"name"`
	Namespace  string        `json:"namespace"`
	Revision   string        `json:"revision"`
	Updated    helmtime.Time `json:"updated"`
	Status     string        `json:"status"`
	Chart      string        `json:"chart"`
	AppVersion string        `json:"appVersion"`
}

//ListRelease list release from harbor
func (w *WorkloadsAPI) ListRelease(g *gin.Context) {
	actionConfig := w.ActionInstance("")
	cmd := action.NewList(actionConfig)
	cmd.AllNamespaces = true
	cmd.Limit = 10000
	cmd.StateMask = action.ListAll
	releases, err := cmd.Run()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	rel := new(Release)
	var rels []Release
	for _, re := range releases {
		rel.Chart = re.Chart.Metadata.Name + "-" + re.Chart.Metadata.Version
		rel.Name = re.Name
		rel.AppVersion = re.Chart.AppVersion()
		rel.Updated = re.Info.LastDeployed
		rel.Namespace = re.Namespace
		rel.Status = string(re.Info.Status)
		rel.Revision = strconv.Itoa(re.Version)
		rels = append(rels, *rel)
	}

	g.JSON(http.StatusOK, rels)
}

//FindReleasesByNamespce find a release
func (w *WorkloadsAPI) FindReleasesByNamespce(g *gin.Context) {
	ns := g.Param("namespace")
	if len(ns) == 0 {
		return
	}

	cmd := action.NewList(w.ActionInstance(ns))
	cmd.AllNamespaces = true
	cmd.Limit = 10000
	cmd.StateMask = action.ListAll
	releases, err := cmd.Run()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	rel := new(Release)
	var rels []Release
	for _, re := range releases {
		rel.Chart = re.Chart.Metadata.Name + "-" + re.Chart.Metadata.Version
		rel.Name = re.Name
		rel.AppVersion = re.Chart.AppVersion()
		rel.Updated = re.Info.LastDeployed
		rel.Namespace = re.Namespace
		rel.Status = string(re.Info.Status)
		rel.Revision = strconv.Itoa(re.Version)
		rels = append(rels, *rel)
	}

	g.JSON(http.StatusOK, rels)
}

type ReleaseDetail struct {
	Info      *release.Info          `json:"info,omitempty"`
	Manifest  string                 `json:"manifest,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Config    map[string]interface{} `json:"config,omitempty"`
	Version   int                    `json:"version,omitempty"`
	NameSpace string                 `json:"namespace,omitempty"`
	Resources *Resources             `json:"resources,omitempty"`
}

type Resources struct {
	items []runtime.Object `json:"items,omitempty"`
}

//FindReleaseByNamespce find a release by namespace
func (w *WorkloadsAPI) FindReleaseByNamespce(g *gin.Context) {
	ns := g.Param("namespace")
	name := g.Param("release")
	if len(ns) == 0 || len(name) == 0 {
		return
	}

	cmd := action.NewList(w.ActionInstance(ns))
	cmd.AllNamespaces = false
	cmd.Limit = 10000
	cmd.StateMask = action.ListAll
	releases, err := cmd.Run()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	rel := new(ReleaseDetail)
	for _, release := range releases {
		if release.Name == name {
			rel.Name = name
			rel.Manifest = release.Manifest
			rel.Info = release.Info
			rel.Config = release.Config
			rel.Version = release.Version
			rel.NameSpace = release.Namespace

			f := cmdutil.NewFactory(helm.NewConfigFlagsFromCluster(release.Namespace, w.RestConfig))
			result := *f.NewBuilder().
				WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
				Stream(bytes.NewBufferString(release.Manifest), "input").
				Flatten().ContinueOnError().Do()

			if err := result.Err(); err != nil {
				common.ToRequestParamsError(g, err)
				return
			}

			items, err := result.Infos()
			if err != nil {
				common.ToRequestParamsError(g, err)
				return
			}
			rel.Resources = new(Resources)
			for _, item := range items {
				obj, err := w.DynamicClient.Resource(item.Mapping.Resource).Namespace(release.Namespace).Get(context.Background(), item.Name, metav1.GetOptions{})
				if err != nil {
					common.ToRequestParamsError(g, err)
					return
				}
				rel.Resources.items = append(rel.Resources.items, obj)
			}
		}
	}

	e, err := k8sjson.Marshal(rel.Resources.items)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"info":      rel.Info,
		"manifest":  rel.Manifest,
		"name":      rel.Name,
		"config":    rel.Config,
		"version":   rel.Version,
		"namespace": rel.NameSpace,
		"resources": gin.H{
			"items": string(e),
		},
	})
}

func (w *WorkloadsAPI) FindReleaseValueByName(g *gin.Context) {
	ns := g.Param("namespace")
	name := g.Param("release")
	if len(ns) == 0 || len(name) == 0 {
		return
	}
	cmd := action.NewList(w.ActionInstance(""))
	cmd.AllNamespaces = false
	cmd.Limit = 10000
	cmd.StateMask = action.ListAll
	releases, err := cmd.Run()
	if err != nil {
		panic(err)
	}

	rel := make(map[string]interface{})
	for _, release := range releases {
		if release.Name == name {
			rel = release.Chart.Values
		}
	}
	b, err := yaml.Marshal(rel)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.String(http.StatusOK, string(b))
}

func (w *WorkloadsAPI) DeleteRelease(g *gin.Context) {
	ns := g.Param("namespace")
	name := g.Param("release")
	if len(ns) == 0 || len(name) == 0 {
		return
	}

	_, err := helm.DeleteRelease(w.ActionInstance(ns), name, false)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)

}

//UpgradeRelease upgrade
func (w *WorkloadsAPI) UpgradeRelease(g *gin.Context) {
	ns := g.Param("namespace")
	name := g.Param("release")
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	obj := &InstallChart{}
	err = json.Unmarshal(rawData, &obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	url := fmt.Sprintf("%s/charts/%s", w.HarborAddress, strings.Split(obj.Chart, "/")[1]+"-"+obj.Version+".tgz")
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := helm.NetClient.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()
	chart, err := loader.LoadArchive(resp.Body)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	val, err := yaml.Marshal(obj.Values)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	rel, err := helm.UpgradeRelease(w.ActionInstance(ns), name, string(val), chart, nil)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"log":     rel,
		"release": rel,
	})
}

type RollBackRelease struct {
	Revision int `json:"revision,omitempty"`
}

//RollbackRelease
func (w *WorkloadsAPI) RollbackRelease(g *gin.Context) {
	ns := g.Param("namespace")
	name := g.Param("release")
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	obj := &RollBackRelease{}
	err = json.Unmarshal(rawData, &obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	rel, err := helm.RollbackRelease(w.ActionInstance(ns), name, obj.Revision)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"message": rel,
	})
}

//HistoryRelease
func (w *WorkloadsAPI) HistoryRelease(g *gin.Context) {
	ns := g.Param("namespace")
	name := g.Param("release")
	if len(ns) == 0 || len(name) == 0 {
		return
	}

	history := action.NewHistory(w.ActionInstance(ns))
	history.Max = 256
	his, err := getHistory(history, name)

	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, his)

}

type releaseInfo struct {
	Revision    int           `json:"revision"`
	Updated     helmtime.Time `json:"updated"`
	Status      string        `json:"status"`
	Chart       string        `json:"chart"`
	AppVersion  string        `json:"app_version"`
	Description string        `json:"description"`
}

type releaseHistory []releaseInfo

func getHistory(client *action.History, name string) (releaseHistory, error) {
	hist, err := client.Run(name)
	if err != nil {
		return nil, err
	}

	releaseutil.Reverse(hist, releaseutil.SortByRevision)

	var rels []*release.Release
	for i := 0; i < min(len(hist), client.Max); i++ {
		rels = append(rels, hist[i])
	}

	if len(rels) == 0 {
		return releaseHistory{}, nil
	}

	releaseHistory := getReleaseHistory(rels)

	return releaseHistory, nil
}

func getReleaseHistory(rls []*release.Release) (history releaseHistory) {
	for i := len(rls) - 1; i >= 0; i-- {
		r := rls[i]
		c := formatChartname(r.Chart)
		s := r.Info.Status.String()
		v := r.Version
		d := r.Info.Description
		a := formatAppVersion(r.Chart)

		rInfo := releaseInfo{
			Revision:    v,
			Status:      s,
			Chart:       c,
			AppVersion:  a,
			Description: d,
		}
		if !r.Info.LastDeployed.IsZero() {
			rInfo.Updated = r.Info.LastDeployed

		}
		history = append(history, rInfo)
	}

	return history
}

func formatChartname(c *chart.Chart) string {
	if c == nil || c.Metadata == nil {
		// This is an edge case that has happened in prod, though we don't
		// know how: https://github.com/helm/helm/issues/1347
		return "MISSING"
	}
	return fmt.Sprintf("%s-%s", c.Name(), c.Metadata.Version)
}

func formatAppVersion(c *chart.Chart) string {
	if c == nil || c.Metadata == nil {
		// This is an edge case that has happened in prod, though we don't
		// know how: https://github.com/helm/helm/issues/1347
		return "MISSING"
	}
	return c.AppVersion()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
