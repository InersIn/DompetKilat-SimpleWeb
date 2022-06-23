package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_KEY = []byte("764b710a-6b62-4199-8f2f-0eccabd65091")

type Claims struct {
	Username string `json:"username"`
	Id       int    `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(username string, id int) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		Id:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		},
	)

	if err != nil {
		err = errors.New("Token not valid")
		return
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		err = errors.New("Token not valid")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return
	}

	return
}

func ExtractToken(signedToken string) (string, int) {
	token, _ := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		},
	)

	claims, _ := token.Claims.(*Claims)
	return claims.Username, claims.Id
}
