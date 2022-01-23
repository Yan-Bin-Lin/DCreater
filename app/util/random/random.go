package random

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an rror if the system's secure random
// number generator fails to function correctly
func GetRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GetRandomString(s int) (string, error) {
	b, err := GetRandomBytes(s)
	return base64.RawURLEncoding.EncodeToString(b), err
}