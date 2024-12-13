package middlewares

import "net/http"

type AuthMiddlewares struct {
}

func AuthMiddlewaresInit() *AuthMiddlewares {
	return &AuthMiddlewares{}
}

func (a *AuthMiddlewares) AuthSessionsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
