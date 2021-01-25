package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"ismaeldf.melo/golang/go-barber/models"
	"ismaeldf.melo/golang/go-barber/repositories"
	"log"
	"time"
)

var mySigningKey = []byte("my_secret_key")

type ResponseAuthenticateUser struct {
	User models.User
	Token string
}

type MyCustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type authenticateUserService struct {
	usersRepository *repositories.UsersRepository
}

const errorMsg = "Incorrect Email/Password combination"

func NewAuthenticateUserService(repository *repositories.UsersRepository) *authenticateUserService {
	return &authenticateUserService{repository}
}

func (s *authenticateUserService) Execute(email string, password string) (*ResponseAuthenticateUser, error) {
	user := s.usersRepository.FindByEmail(email)
	if user.ID == "" {
		return nil, errors.New(errorMsg)
	}

	passwordMatched := isCorrectPassword(user, password)
	if !passwordMatched {
		return nil, errors.New(errorMsg)
	}

	token := createToken(user.ID)

	response := ResponseAuthenticateUser{
		User: user,
		Token: token,
	}

	return &response, nil
}

func createToken(userId string) string{
	var expirationTime = time.Now().Add(1 * time.Hour).Unix()

	claims := MyCustomClaims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err !=nil {
		log.Fatal(err)
	}

	return tokenString
}

//func (s *authenticateUserService) DecodeToken(tokenString string){
//	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return mySigningKey, nil
//	})
//
//	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
//		fmt.Printf("%v %v", claims.Id, claims.StandardClaims.ExpiresAt)
//	} else {
//		fmt.Println(err)
//	}
//}

func isCorrectPassword(user models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}







