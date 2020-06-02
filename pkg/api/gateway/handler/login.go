package handler

import (
	"encoding/json"
	"net/http"
)

type config struct {
	LensVersion       string   `json:"lensVersion"`
	LensTheme         string   `json:"lensTheme"`
	UserName          string   `json:"userName"`
	Token             string   `json:"Token"`
	AllowedNamespaces []string `json:"allowedNamespaces"`
	IsClusterAdmin    bool     `json:"isClusterAdmin"`
	ChartEnable       bool     `json:"chartEnable"`
	KubectlAccess     bool     `json:"kubectlAccess"`
}

func newConfig(user string, token string, allowedNamespaces []string) *config {
	isClusterAdmin := false
	if user == "admin" {
		isClusterAdmin = true
	}
	return &config{
		LensVersion:       "1.0",
		LensTheme:         "",
		UserName:          user,
		Token:             token,
		AllowedNamespaces: allowedNamespaces,
		IsClusterAdmin:    isClusterAdmin,
		ChartEnable:       true,
		KubectlAccess:     true,
	}
}

type LoginHandle struct {
	*AuthorizationStorage
}

func (h *LoginHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bs, err := json.Marshal(newConfig("admin", "", []string{}))
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	switch r.Method {
	case http.MethodPost:
		writeResponse(w, http.StatusOK, bs)
		return
	case http.MethodGet:
		/*
			get config need authorization
		*/

		//baseUser := r.Header.Get("x-auth-username")
		//if !h.Exist(baseUser) {
		//	writeResponse(w, http.StatusBadRequest, "{message: baseUser not exists}")
		//	return
		//}
		writeResponse(w, http.StatusOK, bs)
		return
	default:
	}

	w.WriteHeader(http.StatusBadRequest)
}

func writeResponse(w http.ResponseWriter, status int, data interface{}) {
	var _data []byte
	switch data.(type) {
	case string:
		_data = []byte(data.(string))
	case []byte:
		_data = data.([]byte)
	}
	w.WriteHeader(status)
	w.Write(_data)
	return
}
