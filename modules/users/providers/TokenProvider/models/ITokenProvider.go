package models


type ITokenProvider interface {
	CreateToken(userId string) string
	DecodeToken(tokenString string) (*string, error)
}
