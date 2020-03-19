package hystrix

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	statuscode "github.com/yametech/fuxi/thirdparty/lib/http"
)

// BreakerWrapper hystrix breaker
func BreakerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Method + "-" + r.RequestURI
		err := hystrix.Do(name,
			func() error {
				sct := &statuscode.StatusCodeTracker{ResponseWriter: w, Status: http.StatusOK}
				h.ServeHTTP(sct.WrappedResponseWriter(), r)

				if sct.Status >= http.StatusBadRequest {
					str := fmt.Sprintf("status code %d", sct.Status)
					log.Println(str)
					h.ServeHTTP(sct.WrappedResponseWriter(), r)
					return errors.New(str)
				}
				return nil
			},
			nil,
		)
		if err != nil {
			log.Printf("hystrix breaker %s err: %s", name, err)
			return
		}
	})
}
