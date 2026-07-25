package local

import (
	"7DL/auth/domain"
	repository "7DL/auth/infra/local/sqlite"
	"7DL/config"
	identitydomain "7DL/identity/domain"
	"7DL/sqlite"
	"fmt"
)

type Container struct {
}

type Auth struct {
	Config   *config.AuthConfig
	Identity identitydomain.IdentityPort
	Repo     *repository.SQLite
}

func New(
	config *config.AuthConfig,
	identity identitydomain.IdentityPort,
	sqliteDb *sqlite.SQLiteDB,
) *Auth {
	repo := repository.New(sqliteDb)

	return &Auth{
		Config: config,
		Repo:   repo,
	}
}

func (auth *Auth) Login(payload *domain.LoginPayload) error {
	fmt.Printf("mode: %s username: %s password: %s", "login", payload.Username, payload.Password)
	return nil
}
