package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
	"net/url"
	"strings"
)

type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

type Owner struct {
	Id        uint32 `json:"id"`
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	Language  string `json:"language"`
	Username  string `json:"username"`
}

type Repository struct {
	Id              uint32      `json:"id"`
	Owner           Owner       `json:"owner"`
	Name            string      `json:"name"`
	FullName        string      `json:"full_name"`
	Description     string      `json:"description"`
	Empty           bool        `json:"empty"`
	Private         bool        `json:"private"`
	Fork            bool        `json:"fork"`
	Parent          string      `json:"parent"`
	Mirror          bool        `json:"mirror"`
	Size            uint32      `json:"size"`
	HtmlUrl         string      `json:"html_url"`
	SshUrl          string      `json:"ssh_url"`
	CloneUrl        string      `json:"clone_url"`
	Website         string      `json:"website"`
	StarsCount      uint32      `json:"stars_count"`
	ForksCount      uint32      `json:"forks_count"`
	WatchersCount   uint32      `json:"watchers_count"`
	OpenIssuesCount uint32      `json:"open_issues_count"`
	DefaultBranch   string      `json:"default_branch"`
	Archived        bool        `json:"archived"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	Permissions     Permissions `json:"permissions"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Commit struct {
	Id           string `json:"id"`
	Message      string `json:"message"`
	Url          string `json:"url"`
	Author       User   `json:"author"`
	Committer    User   `json:"committer"`
	Verification string `json:"verification"`
	Timestamp    string `json:"timestamp"`
}

type giteaWebHook struct {
	Secret     string     `json:"secret"`
	Ref        string     `json:"ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Commits    []Commit   `json:"commits"`
	Repository Repository `json:"repository"`
	Pusher     Owner      `json:"pusher"`
	Sender     Owner      `json:"sender"`
}

func checkUrl(RequestHtmlUrl string, tektonWebHookGit string) error {
	htmlUrl, err := url.Parse(RequestHtmlUrl)
	if err != nil {
		return err
	}
	gitUrl, err := url.Parse(tektonWebHookGit)
	if err != nil {
		return err
	}
	gitUrl.Path = strings.TrimRight(gitUrl.Path, ".git")
	if (htmlUrl.Host != gitUrl.Host) || (htmlUrl.Path != gitUrl.Path) {
		return errors.New("git url checksum error")
	}

	return nil
}

func checkBranch(jobBranch string, giteaWebHookRef string) bool {
	if jobBranch == giteaWebHookRef {
		return true
	}
	if strings.Contains(giteaWebHookRef, "/") {
		refList := strings.Split(giteaWebHookRef, "/")
		if len(refList) > 0 && refList[len(refList)-1] == jobBranch {
			return true
		}
	}
	return false
}

func (w *WorkloadsAPI) TriggerGiteaWebHook(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.tektonWebHook.Get(namespace, name)
	if err != nil {
		g.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "message": "webhook instance does not exist"})
		return
	}

	tektonWebHook := &v1.TektonWebHook{}
	if err := common.RuntimeObjectToInstanceObj(item, tektonWebHook); err != nil {
		g.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "message": "webhook instance serialization error"})
		return
	}

	giteaWebHook := giteaWebHook{}
	err = json.Unmarshal(rawData, &giteaWebHook)
	if err != nil {
		g.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "message": "webhook instance serialization error"})
		return
	}

	if tektonWebHook.Spec.Secret != giteaWebHook.Secret {
		g.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "message": "webhook secret error."})
		return
	}

	if err := checkUrl(giteaWebHook.Repository.HtmlUrl, tektonWebHook.Spec.Git); err != nil {
		g.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "message": "webhook git url check error."})
		return
	}

	for _, job := range tektonWebHook.Spec.Jobs {
		if checkBranch(job.Branch, giteaWebHook.Ref) {

			item, err := w.pipelineRun.Get(namespace, job.PipelineRun)
			if err != nil {
				common.ToInternalServerError(g, rawData, err)
				return
			}

			unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&item)
			if err != nil {
				common.ToRequestParamsError(g, err)
				return
			}
			fmt.Print(len(job.Params))
			if len(job.Params) > 0 {
				unstructuredObj["spec"].(map[string]interface{})["params"] = job.Params
			}
			metadata := map[string]interface{}{
				"annotations": unstructuredObj["metadata"].(map[string]interface{})["annotations"],
				"labels":      unstructuredObj["metadata"].(map[string]interface{})["labels"],
				"name":        unstructuredObj["metadata"].(map[string]interface{})["name"],
				"namespace":   unstructuredObj["metadata"].(map[string]interface{})["namespace"],
			}
			pipelineRunUnstructuredObj := map[string]interface{}{
				"apiVersion": unstructuredObj["apiVersion"],
				"kind":       unstructuredObj["kind"],
				"metadata":   metadata,
				"spec":       unstructuredObj["spec"],
			}

			pipelineRunUnstructured := &unstructured.Unstructured{
				Object: pipelineRunUnstructuredObj,
			}
			_ = w.pipelineRun.Delete(namespace, job.PipelineRun)
			_, _, err = w.pipelineRun.Apply(namespace, job.PipelineRun, pipelineRunUnstructured)
			if err != nil {
				g.JSON(http.StatusInternalServerError,
					gin.H{"code": http.StatusInternalServerError, "message": "webhook trigger error, please check pipeline run"})
				return
			}
			g.JSON(http.StatusOK,
				gin.H{"code": http.StatusOK, "message": "webhook triggered successfully"})
			return
		}
	}

	g.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "message": "webhook did not trigger any instances"})
	return
}
