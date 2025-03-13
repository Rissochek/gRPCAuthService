package auth

import (
	"AuthProject/internal/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type JWTCustomClaims struct {
	Username string
	UserId   uint
	jwt.StandardClaims
}

type JWTManager struct {
	SecretKey     string
	TokenDuration time.Duration
}

func (manager *JWTManager) GenerateJWT(user *model.User) (string, error) {
	claims := JWTCustomClaims{
		Username: user.Usermame,
		UserId:   user.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.TokenDuration).Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.SecretKey))
}

func (manager *JWTManager) VerifyJWT(user_token string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		user_token,
		&JWTCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("wrong jwt encrypting method")
			}

			return []byte(manager.SecretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (manager *JWTManager) GetSecretKeyFromEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("no .env file found") //if u have this error, but .env file is existing then try to execute with this command go run .\cmd\server\main.go
	}

	secret, exists := os.LookupEnv("SECRET_KEY")
	if !exists {
		log.Fatalf("secret key value is not set in .env file.")
	}
	manager.SecretKey = secret
}
