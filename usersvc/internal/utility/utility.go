package utility

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// HashPassword creates a hashed password using Argon2id
func HashPassword(password string) (string, error) {
	// Parameters (you can tune these)
	const time = 1
	const memory = 64 * 1024
	const threads = 4
	const keyLen = 32
	const saltLen = 16

	// Generate random salt
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Derive key
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Encode as string (salt + hash)
	encoded := fmt.Sprintf("%s:%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return encoded, nil
}

// VerifyPassword checks password against hash
func VerifyPassword(password, encodedHash string) bool {
	// Split salt:hash
	var saltBase64, hashBase64 string
	_, err := fmt.Sscanf(encodedHash, "%[^:]:%s", &saltBase64, &hashBase64)
	if err != nil {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(saltBase64)
	if err != nil {
		return false
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(hashBase64)
	if err != nil {
		return false
	}

	// Parameters must match HashPassword
	const time = 1
	const memory = 64 * 1024
	const threads = 4
	const keyLen = 32

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Compare byte slices
	if len(hash) != len(expectedHash) {
		return false
	}
	for i := range hash {
		if hash[i] != expectedHash[i] {
			return false
		}
	}
	return true
}
