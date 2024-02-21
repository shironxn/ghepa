package util

import (
	"errors"
	"event-planning-app/config"
	"event-planning-app/internal/core/domain"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
}

func (j *JWTManager) GenerateToken(w http.ResponseWriter, user domain.User) (*domain.UserDetails, error) {
	conf := config.GetConfig()

	expirationTime := jwt.NewNumericDate(time.Now().Add(10 * time.Hour))
	claims := &domain.Claims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	if user.Email == conf.Admin.Email {
		claims.Role = "admin"
	} else {
		claims.Role = "user"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.JWTSecret))

	expireTime := time.Now().Add(10 * time.Hour)
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Expires:  expireTime,
	}

	http.SetCookie(w, cookie)

	details := domain.UserDetails{
		Token:   tokenString,
		Expired: expireTime.String(),
	}

	return &details, err
}

func (j *JWTManager) ValidateToken(token string) (jwt.Claims, error) {
	conf := config.GetConfig()

	tokens, err := jwt.ParseWithClaims(token, &domain.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !tokens.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := tokens.Claims.(*domain.Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
