package middleware

import (
	"context"
	"event-planning-app/internal/util"
	"net/http"
)

type AuthMiddleware struct {
	response util.Response
	jwt      util.JWTManager
}

func NewAuthMiddleware(jwtManager util.JWTManager) *AuthMiddleware {
	return &AuthMiddleware{
		jwt: jwtManager,
	}
}

func (j *AuthMiddleware) JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil || c == nil {
			j.response.Error(w, http.StatusUnauthorized, "unauthorized access", err.Error())
			return
		}

		tokenString := c.Value
		claims, errToken := j.jwt.ValidateToken(tokenString)
		if errToken != nil {
			j.response.Error(w, http.StatusUnauthorized, "invalid Token", errToken.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})

}
