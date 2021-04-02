package middleware

import (
	"net/http"
	"time"
)

var (
	// Unix epoch time
	epoch = time.Unix(0, 0).Format(time.RFC1123)

	// Taken from https://github.com/mytrile/nocache
	noCacheHeaders = map[string]string{
		"Expires":         epoch,
		"Cache-Control":   "no-cache, private, no-store, must-revalidate, max-age=0",
		"Pragma":          "no-cache",
		"X-Accel-Expires": "0",
	}
)

// NoCacheMiddleware adds no cache headers to the response header, that tell browsers
// not to cache the content
func NoCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range noCacheHeaders {
			w.Header().Set(k, v)
		}

		next.ServeHTTP(w, r)
	})
}
