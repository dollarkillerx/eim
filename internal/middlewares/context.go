package middlewares

import (
	"context"
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

			next.ServeHTTP(w, r)
		})
	}
}
