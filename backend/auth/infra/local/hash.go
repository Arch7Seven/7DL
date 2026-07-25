package local

import (
	"7DL/config"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

type Credentials struct {
	Salt         string
	PasswordHash string
}

func generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	return salt, err
}

func generatePasswordHash(argonConfig config.Argon2Config, password string, salt []byte) string {
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		argonConfig.Iteration,
		argonConfig.Memory,
		argonConfig.Threads,
		argonConfig.KeyLength,
	)

	return base64.RawStdEncoding.EncodeToString(hash)
}

func NewPasswordHash(hashConfig config.Hash, password string) (*Credentials, error) {
	salt, err := generateSalt(hashConfig.SaltSize)
	if err != nil {
		return nil, err
	}
	hash := generatePasswordHash(hashConfig.Argon2Config, password, salt)

	return &Credentials{
		Salt:         base64.RawStdEncoding.EncodeToString(salt),
		PasswordHash: hash,
	}, nil
}

func verifyPassword(argonConfig config.Argon2Config, password string, credentials *Credentials) bool {

	salt, _ := base64.RawStdEncoding.DecodeString(credentials.Salt)
	hash, _ := base64.RawStdEncoding.DecodeString(credentials.PasswordHash)

	newHash := argon2.IDKey(
		[]byte(password),
		salt,
		argonConfig.Iteration,
		argonConfig.Memory,
		argonConfig.Threads,
		uint32(len(hash)),
	)

	return subtle.ConstantTimeCompare(hash, newHash) == 1
}
