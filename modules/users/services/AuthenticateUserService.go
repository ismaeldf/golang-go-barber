package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	providers "ismaeldf/golang-gobarber/modules/users/providers/HashProvider/models"
	"ismaeldf/golang-gobarber/modules/users/repositories"
	"strings"
	"time"
)

var mySigningKey = []byte("my_secret_key")

type ResponseAuthenticateUser struct {
	User  entities.User
	Token string
}

type MyCustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type authenticateUserService struct {
	usersRepository repositories.IUserRepository
	hashProvider providers.IHashProvider
}

const errorMsg = "Incorrect Email/Password combination"

func NewAuthenticateUserService(repository repositories.IUserRepository, hashProvider providers.IHashProvider) *authenticateUserService {
	return &authenticateUserService{usersRepository: repository, hashProvider: hashProvider}
}

func (s *authenticateUserService) Execute(email string, password string) (*ResponseAuthenticateUser, error) {
	user := s.usersRepository.FindByEmail(email)
	if user.Id == "" {
		return nil, errors.New(errorMsg)
	}

	passwordMatched := s.isCorrectPassword(user, password)
	if !passwordMatched {
		return nil, errors.New(errorMsg)
	}

	token := createToken(user.Id)

	response := ResponseAuthenticateUser{
		User: user,
		Token: token,
	}

	return &response, nil
}

func createToken(userId string) string{
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

func DecodeToken(tokenString string) (*string, error){
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

func (a *authenticateUserService) isCorrectPassword(user entities.User, password string) bool {
	return a.hashProvider.CompareHash(password, user.Password)
}








