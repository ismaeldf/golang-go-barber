package implementations

import "golang.org/x/crypto/bcrypt"

type BCryptHashProvider struct {}

func (b *BCryptHashProvider) GenerateHash(payload string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	return string(hashedPassword)
}


func (b *BCryptHashProvider) CompareHash(payload string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(payload))
	return err == nil
}
