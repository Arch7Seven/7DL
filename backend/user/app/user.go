package userusecase

import (
	user "7DL/user/domain"
	"time"

	"github.com/google/uuid"
)

// TODO :
// max username char : 20

type UserUseCase struct {
	UserRepo user.UserRepository
}

func NewUserUseCase(userRepo user.UserRepository) *UserUseCase {
	return &UserUseCase{}
}

func (uc *UserUseCase) NewAccount(username string)

func (uc *UserUseCase) CreateUser(username string) {

	var userUUID uuid.UUID
	for {
		userUUID = uuid.New()
		exist, err := uc.UserRepo.UserExistFromUUID(userUUID)

		if err != nil {
			panic(err) // to manage
		}

		if exist {
			continue
		}
		break
	}

	user := &user.User{
		UUID:      userUUID,
		Username:  username,
		CreatedAt: time.Now(),
	}

	uc.UserRepo.Save(user)
}
