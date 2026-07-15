package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID
	Username  string
	CreatedAt time.Time
	Active    bool
	Settings  *UserSettings
}

type UserSettings struct {
	PreferredVideoCodec string
	PreferredAudioCodec string
	Cookies             *Cookies
}

type Cookies struct {
	Youtube string
}
