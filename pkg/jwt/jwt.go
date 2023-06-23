package jwt

import (
	"errors"
	"github.com/HuckOps/notify/src/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

type ServerClaims struct {
	User User
	jwt.StandardClaims
}

func GetToken(user User) (string, error) {
	claims := ServerClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Config.JWT.Exp) * time.Second).Unix(),
			Issuer:    config.Config.JWT.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.JWT.Secret))
}

func ParseToken(tokenString string) (*ServerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ServerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*ServerClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
