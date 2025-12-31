package jwtauth

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JWTSession struct {
	SecretKey []byte
}

func NewJWTSession(secretKey []byte) *JWTSession {
	return &JWTSession{SecretKey: secretKey}
}

func CreateJWTToken(email string, secretKey []byte) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "hhtg",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}

	return token, nil
}

