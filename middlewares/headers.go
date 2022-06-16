package middlewares

import "net/http"

var Headers = map[string]string{
	"Strict-Transport-Security": `max-age=63072000; includeSubDomains; preload`,
	"Access-Control-Allow-Origin": `"*"`,
	"X-Frame-Options":             `deny`,
	"Cache-Control":               `no-store`,
	"Clear-Site-Data":             `"zlatanov.me"`,
	"Feature-Policy":              `microphone 'none'; camera 'none'`,
	"Expires":                     `0`,
	"X-XSS-Protection":            `1; mode=block`,
	"Authorization": "eyJ1c2VyX2lkIjogMSwgInVzZXJfdHlwZSI6Imd1ZXN0IiwgImF1dGhvcml6YXRpb25fbGV2ZWwiOiAwfQ==",
}

func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for headerName, headerValue := range Headers {
			w.Header().Add(headerName, headerValue)
		}
		next.ServeHTTP(w, r)
	})
}
