package middleware

import (
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
			j.response.Error(w, http.StatusUnauthorized, "Unauthorized access", err.Error())
			return
		}

		tokenString := c.Value
		errToken := j.jwt.ValidateToken(tokenString)
		if errToken != nil {
			j.response.Error(w, http.StatusUnauthorized, "Invalid Token", errToken.Error())
			return
		}

		next.ServeHTTP(w, r)
	})

}
