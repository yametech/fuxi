package auth

import (
	"net/http"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/micro/plugin"
	"github.com/yametech/fuxi/thirdparty/lib/token"
	"github.com/yametech/fuxi/thirdparty/lib/whitelist"
)

// JWTAuthWrapper
func JWTAuthWrapper(token *token.Token, whitelist *whitelist.Whitelist) plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO dynamic name list from configure server
			if whitelist.In(r.URL.Path) {
				log.Infof("url (%s) in white list", r.URL.Path)
				h.ServeHTTP(w, r)
				return
			}
			tokenHeader := r.Header.Get("Authorization")
			userFromToken, e := token.Decode(tokenHeader)

			if e != nil {
				log.Infof(`Jwt auth wrapper decode token error key "%s"`, tokenHeader)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			r.Header.Set("x-auth-username", userFromToken.UserName)
			h.ServeHTTP(w, r)
		})
	}
}
