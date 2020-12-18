package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yametech/fuxi/common"
)

type LoginHandle struct {
	*Authorization
}

func parseUri(uri string) (isNamespaced bool, namespace string, bypass bool) {
	if strings.Contains(uri, "/api/metrics") || strings.Contains(uri, "/watch") {
		bypass = true
		return
	}

	if strings.Contains(uri, "/namespaces") && !strings.HasPrefix(uri, "/workload/api/v1") {
		actions := strings.Split(uri, "/")
		actions = trimSpace(actions)
		idx := index(actions, "namespaces")
		if idx == -1 {
			isNamespaced = false
			return
		}
		isNamespaced = true
		namespace = actions[idx+1]
	}

	return
}

type userAuth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (h *LoginHandle) Check(username string, w http.ResponseWriter, r *http.Request) bool {
	isNamespaced, namespace, bypass := parseUri(r.URL.Path)
	if username == "admin" {
		return true
	}
	if bypass {
		return true
	}
	if !isNamespaced {
		writeResponse(w, http.StatusForbidden, "Unauthorized Access")
	}
	allow, err := h.allowNamespaceAccess(username, namespace)
	if !allow || err != nil {
		writeResponse(w, http.StatusForbidden, "Not Allowed Namespaces")
	}
	return allow
}

func (h *LoginHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get(common.HttpRequestUserHeaderKey)
	userAuth := &userAuth{}
	if username == "" {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := json.Unmarshal(bs, userAuth); err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if userAuth.UserName == "" || userAuth.Password == "" {
			writeResponse(w, http.StatusBadRequest, "you are bug user")
			return
		}
	} else {
		userAuth.UserName = username
	}

	if r.Method == http.MethodPost && r.URL.Path == "/user-login" {
		// user Login
		cfgData, err := h.Auth(userAuth.UserName, userAuth.Password)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		writeResponse(w, http.StatusOK, cfgData)
		return
	}

	if r.Method == http.MethodGet && r.URL.Path == "/config" {
		// h.Config()
		// Reserved for future use
	}

	w.WriteHeader(http.StatusUnauthorized)
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
