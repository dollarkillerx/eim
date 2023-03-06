package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dollarkillerx/eim/internal/pkg/enum"
	"github.com/dollarkillerx/eim/internal/utils"
	"github.com/rs/xid"
)

// Context  get user from jwt and put user into ctx
func Context() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := xid.New().String()
			ctx := context.WithValue(r.Context(), enum.RequestId, reqID)

			ctx = context.WithValue(ctx, enum.RequestReceivedAtCtxKey, time.Now())

			// user agent
			userAgent := r.Header.Get("User-Agent")
			ctx = context.WithValue(ctx, enum.UserAgentCtxKey, userAgent)

			// token
			tokenString, _ := utils.GetTokenFromHeader(r.Header)
			if tokenString != "" {
				ctx = context.WithValue(ctx, enum.TokenCtxKey, tokenString)
			}

			r = r.WithContext(ctx)

			log.Println(r.URL.String())
			next.ServeHTTP(w, r)
		})
	}
}

func Safety() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			salt := utils.Md5Encode(fmt.Sprintf("%d-9776e538-59ba-473f-8ccf-1d72031e360f", time.Now().UnixMilli()/10000))
			reqSalt := r.Header.Get("salt")
			if salt == reqSalt {
				next.ServeHTTP(w, r)
				return
			}

			respData := map[string]interface{}{
				"errors": []map[string]interface{}{
					{
						"message": "500 server exception",
						"path":    []string{"main"},
					},
				},
			}

			marshal, err := json.Marshal(respData)
			if err == nil {
				w.Write(marshal)
			}
			//w.WriteHeader(500)
		})
	}
}
