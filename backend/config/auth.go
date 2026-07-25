package config

type AuthConfig struct {
	Hash              Hash
	UsernameMaxLength int
	UsernameMinLength int
	PasswordMaxLength int
	PasswordMinLength int
}
type Hash struct {
	Argon2Config Argon2Config
	SaltSize     int
}

type Argon2Config struct {
	Iteration uint32
	Memory    uint32 //kib
	Threads   uint8
	KeyLength uint32
}

func GetAuthConfig() AuthConfig {

	return AuthConfig{
		PasswordMaxLength: 100,
		PasswordMinLength: 8,
		UsernameMaxLength: 32,
		UsernameMinLength: 1,
		Hash: Hash{
			Argon2Config: Argon2Config{
				Iteration: 3,
				Memory:    128 * 1024,
				Threads:   2,
				KeyLength: 32,
			},
			SaltSize: 32,
		},
	}
}
