package auth

import (
	"strconv"
	"time"

	"github.com/faiz-gh/go-postgresql-starter/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userID int) (string, error) {
	expiration := time.Second * time.Duration(config.ENV.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.Itoa(userID),
		"expired_at": time.Now().Add(expiration).Unix(),
	})

	secret := []byte(config.ENV.JWTSecret)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}