package middlewares

import (
	"context"
	"net/http"

	"github.com/HsiaoCz/app-stone/confuse/csc/storage"
	"github.com/HsiaoCz/app-stone/confuse/csc/types"
)

type AuthMiddlewares struct {
	sen storage.SessionStorer
}

func AuthMiddlewaresInit(sen storage.SessionStorer) *AuthMiddlewares {
	return &AuthMiddlewares{
		sen: sen,
	}
}

func (a *AuthMiddlewares) AuthSessionsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "please login", http.StatusNonAuthoritativeInfo)
			return
		}
		// 这一步有点儿多余了
		// token := cookie.Value
		// if token == "" {
		// 	http.Error(w, "please login", http.StatusNonAuthoritativeInfo)
		// 	return
		// }
		session, err := a.sen.GetSessionByToken(r.Context(), cookie.Value)
		if err != nil {
			http.Error(w, "please login", http.StatusNonAuthoritativeInfo)
			return
		}
		ctx := context.WithValue(r.Context(), types.UserSessionKey, session)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
