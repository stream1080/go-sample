package encrypt

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(src string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ValidPassword use golang.org/x/crypto/bcrypt compare passwords for equality
func ValidPassword(plaintext, ciphertext string) bool {
	if len(ciphertext) <= 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(ciphertext), []byte(plaintext))

	return err == nil
}
