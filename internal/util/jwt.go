package util

import (
	"errors"
	"event-planning-app/config"
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/response"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	Name  string
	Email string
	jwt.RegisteredClaims
}

type JWTManager struct {
}

func (j *JWTManager) GenerateToken(w http.ResponseWriter, user domain.User) (*response.UserDetails, error) {
	conf := config.GetConfig()

	expirationTime := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &jwtClaims{
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.JWTSecret))

	expireTime := time.Now().Add(1 * time.Hour)
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Expires:  expireTime,
	}

	http.SetCookie(w, cookie)

	details := response.UserDetails{
		Token:   tokenString,
		Expired: expireTime.String(),
	}

	return &details, err
}

func (j *JWTManager) ValidateToken(token string) error {
	conf := config.GetConfig()

	tokens, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.JWTSecret), nil
	})

	if err != nil {
		return err
	}

	if !tokens.Valid {
		return errors.New("token is not valid")
	}

	claims, ok := tokens.Claims.(*jwtClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return errors.New("token has expired")
	}

	return nil
}
