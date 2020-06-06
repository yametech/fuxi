package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/util/common"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yametech/fuxi/thirdparty/lib/token"
)

type LoginHandle struct {
	*token.Token
	AuthorizationStorage
}

type userAuth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var dataBase = map[string]string{
	"admin": "admin",
	"dev":   "dev",
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
		// TODO login
		//ok, err := h.Auth(userAuth.UserName, userAuth.Password)
		//if !ok || err != nil {
		//	writeResponse(w, http.StatusBadRequest, err.Error())
		//	return
		//}

		if pwd, exist := dataBase[userAuth.UserName]; !exist {
			writeResponse(w, http.StatusUnauthorized, "{message: user not exist}")
			return
		} else if pwd != userAuth.Password {
			writeResponse(w, http.StatusUnauthorized, "{message: password incorrect}")
			return
		}
		expireTime := time.Now().Add(time.Hour * 24).Unix()
		tokenStr, err := h.Encode("go.micro.gateway.login", userAuth.UserName, expireTime)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		config := newUserConfig(userAuth.UserName, tokenStr, []string{})
		bytesData, err := json.Marshal(config)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		writeResponse(w, http.StatusOK, bytesData)
		return
	}

	if r.Method == http.MethodGet && r.URL.Path == "/config" {
		if userAuth.UserName == "admin" {
			expireTime := time.Now().Add(time.Hour * 24).Unix()
			tokenStr, err := h.Encode("go.micro.gateway.login", userAuth.UserName, expireTime)
			if err != nil {
				writeResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			config := newUserConfig(userAuth.UserName, tokenStr, []string{})
			bytesData, err := json.Marshal(config)
			if err != nil {
				writeResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			writeResponse(w, http.StatusOK, bytesData)
		}
		return
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
