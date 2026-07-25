package local

import (
	"7DL/config"
	"testing"
)

func TestHashSamePassword(t *testing.T) {
	password := "kodfglkdfk_*-edorfkgèçofdpo4g5ert45@"

	hashConfig := config.GetAuthConfig().Hash

	cred, err := NewPasswordHash(hashConfig, password)
	if err != nil {
		t.Fatal(err.Error())
	}

	check := verifyPassword(hashConfig.Argon2Config, password, cred)

	if !check {
		t.Fatalf("Hash is not the same")
	}
}
