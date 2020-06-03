package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yametech/fuxi/thirdparty/lib/token"
)

type LoginHandle struct {
	*token.Token
	AuthorizationStorage
}

func (h *LoginHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var username string
	var password string

	if r.Method == http.MethodPost && r.URL.Path == "/login" {
		ok, err := h.Auth(username, password)
		if !ok || err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		expireTime := time.Now().Add(time.Hour * 24).Unix()
		tokenStr, err := h.Encode("go.micro.gateway.auth", username, expireTime)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		config := newUserConfig(username, tokenStr, []string{})
		bytesData, err := json.Marshal(config)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		writeResponse(w, http.StatusOK, bytesData)
		return
	}

	if r.Method == http.MethodGet && r.URL.Path == "/config" {
		writeResponse(w, http.StatusOK, []byte{})
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
