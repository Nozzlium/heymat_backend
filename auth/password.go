package auth

import "golang.org/x/crypto/bcrypt"

func hashPassword(
	password string,
) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		10,
	)
	return string(bytes), err
}

func compareHashWithPassword(
	hash string,
	password string,
) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
