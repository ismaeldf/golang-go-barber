package models

type IHashProvider interface {
	GenerateHash(payload string) string
	CompareHash(payload string, hashed string) bool
}
