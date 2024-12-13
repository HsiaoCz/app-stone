package middlewares

import "net/http"

func SessionMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "no cookie so please login", http.StatusNonAuthoritativeInfo)
			return
		}
		_ = cookie
	}
}
