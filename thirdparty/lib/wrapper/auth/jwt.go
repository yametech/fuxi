package auth

import (
	"github.com/micro/micro/plugin"
	"github.com/yametech/fuxi/thirdparty/lib/token"
	"github.com/yametech/fuxi/thirdparty/lib/whitelist"
	"net/http"
)

// JWTAuthWrapper
func JWTAuthWrapper(token *token.Token, whitelist *whitelist.Whitelist, loginHandler http.Handler) plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// dynamic name list from configure server
			//if whitelist.In(r.URL.Path) {
			//	h.ServeHTTP(w, r)
			//	return
			//}
			_ = whitelist

			//var tokenHeader string
			if r.URL.Path == "/login" {
				loginHandler.ServeHTTP(w, r)
				return
			}
			if r.URL.Path == "/config" {
				loginHandler.ServeHTTP(w, r)
				return
			}

			//tokenHeader = r.Header.Get("Authorization")
			//userFromToken, e := token.Decode(tokenHeader)
			//if e != nil {
			//	w.WriteHeader(http.StatusUnauthorized)
			//	return
			//}
			//
			//r.Header.Set("x-auth-username", userFromToken.UserName)

			// Config
			if r.URL.Path == "/config" {
				loginHandler.ServeHTTP(w, r)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
