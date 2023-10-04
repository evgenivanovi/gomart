package auth

import "golang.org/x/crypto/bcrypt"

const HASH = "gomart"

// GenerateHashPassword
// This function takes a plain text password as input and returns a hash value generated from it using a one-way hashing algorithm.
// The purpose of this function is to store a user's password securely in the database.
// This way, even if the database is compromised, the attacker cannot retrieve the original password as it is encrypted.
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CompareHashPassword
// This function takes the user input password and the hashed password stored in the database and compares them.
// If the hashes match, it returns true.
// This function is used to verify if the user has entered the correct password during login.
func CompareHashPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
