package functions

import (
	"errors"
	"os"
	"time"

	"github.com/SMRPcoder/markable/errorlog"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type JWTuser struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	jwt.StandardClaims
}

func EncodeJwt(user JWTuser) (string, error) {
	if err := godotenv.Load(); err != nil {
		errorlog.LogError(err.Error())
	}
	var (
		key []byte
		t   *jwt.Token
		s   string
	)
	mySecretKey := os.Getenv("MY_SECRET_JWT")
	key = []byte(mySecretKey)
	expTime := time.Now().Add(24 * time.Hour)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTuser{
		Username: user.Username,
		Name:     user.Name,
		Id:       user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	})
	str, err := t.SignedString(key)
	if err != nil {
		return "", errorlog.LogError(err.Error())
	}
	s = str
	return s, nil

}

func DecodeJwt(tokenString string) (JWTuser, error) {
	if err := godotenv.Load(); err != nil {
		return JWTuser{}, errors.New(err.Error())
	}

	mySecretKey := os.Getenv("MY_SECRET_JWT")
	key := []byte(mySecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &JWTuser{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return JWTuser{}, errors.New(err.Error())
	}

	if claims, ok := token.Claims.(*JWTuser); ok && token.Valid {
		return *claims, nil
	} else {
		return JWTuser{}, errors.New("invalid jwt")
	}
}
