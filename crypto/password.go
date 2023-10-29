package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type passwordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

// Generate Encrypted Password
func GeneratePassword(password string) (string, error) {
	c := &passwordConfig{
		time:    2,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	// Generate Salt
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	// Base64 encode password
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	full := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
	return full, err
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}

// Compare Password with the Hashed value
func ComparePassword(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	c := &passwordConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	decodedSalt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	c.keyLen = uint32(len(decodedHash))
	comparisonHash := argon2.IDKey([]byte(password), decodedSalt, c.time, c.memory, c.threads, c.keyLen)
	return (subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1), nil
}
