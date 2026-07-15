package local

import (
	"7DL/auth/domain"
	"7DL/config"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

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

func genSaltAndPasswordHash(hashConfig config.Hash, password string) *domain.Credentials {
	salt, _ := generateSalt(hashConfig.SaltSize) // manage error
	hash := generatePasswordHash(hashConfig.Argon2Config, password, salt)

	return &domain.Credentials{
		Salt:         base64.RawStdEncoding.EncodeToString(salt),
		PasswordHash: hash,
	}
}

func verifyPassword(argonConfig config.Argon2Config, password string, credentials *domain.Credentials) bool {

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
