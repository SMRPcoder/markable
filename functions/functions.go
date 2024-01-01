package functions

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type JWTuser struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	jwt.RegisteredClaims
}

func EncodeJwt(user JWTuser) (string, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error", err.Error())
	}
	var (
		key []byte
		t   *jwt.Token
		s   string
	)
	mySecretKey := os.Getenv("JWT_SECRET")
	if mySecretKey == "" {
		return "", errors.New("JWT_SECRET not set in environment variables")
	}
	key = []byte(mySecretKey)
	exptime := time.Now().Add(time.Hour * 24)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, JWTuser{
		Username: user.Username,
		Name:     user.Name,
		Id:       user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exptime),
			Issuer:    "markable",
		},
	})
	str, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	s = str
	return s, nil

}

func DecodeJwt(tokenString string) (JWTuser, error) {
	if err := godotenv.Load(); err != nil {
		return JWTuser{}, errors.New(err.Error())
	}

	mySecretKey := os.Getenv("JWT_SECRET")
	if mySecretKey == "" {
		return JWTuser{}, errors.New("JWT_SECRET not set in environment variables")
	}
	key := []byte(mySecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &JWTuser{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return JWTuser{}, err
	} else if claims, ok := token.Claims.(*JWTuser); ok {
		return *claims, nil
	} else {
		return JWTuser{}, errors.New("unknown claims type, cannot proceed")
	}
}
