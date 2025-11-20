package logger

import (
	"context"
	"net/http"

	"github.com/pc-configurator/components/pkg/router"
	"github.com/rs/zerolog/log"
)

type contextErrKey struct{}

func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := context.WithValue(r.Context(), contextErrKey{}, &err)

		ww := router.NewWrapWriter(w)

		next.ServeHTTP(ww, r.WithContext(ctx))

		event := log.Info()

		if err != nil {
			event = log.Error().Err(err)
		}

		event.Int("code", ww.Code()).Str("method", r.Method).Str("path", r.URL.Path).Send()
	}

	return http.HandlerFunc(fn)
}

func SetCtxError(ctx context.Context, err error) {
	ctxErr, ok := ctx.Value(contextErrKey{}).(*error)
	if ok {
		*ctxErr = err
	}
}
