package implementations

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var mySigningKey = []byte("my_secret_key")

type MyCustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type JwtTokenProvider struct{}

func (j *JwtTokenProvider) CreateToken(userId string) string {
	var expirationTime = time.Now().Add(24 * time.Hour).Unix()

	claims := MyCustomClaims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString
}

func (j *JwtTokenProvider) DecodeToken(tokenString string) (*string, error){
	removedBearer := strings.Trim(strings.ReplaceAll(tokenString, "Bearer", ""), " ")

	token, err := jwt.ParseWithClaims(removedBearer, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return &claims.Id, nil
	} else {
		return nil, err
	}
}
