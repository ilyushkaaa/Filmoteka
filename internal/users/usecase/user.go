package usecase

import (
	"github.com/ilyushkaaa/Filmoteka/internal/users/entity"
	"github.com/ilyushkaaa/Filmoteka/internal/users/repo"
	passwordHash "github.com/ilyushkaaa/Filmoteka/pkg/password_hash"
)

type UserUseCase interface {
	Login(username, password string) (*entity.User, error)
	Register(username, password string) (*entity.User, error)
	GetUserRole(userID uint64) (string, error)
}

type UserUseCaseApp struct {
	userRepo repo.UserRepo
}

func (uc *UserUseCaseApp) Login(username, password string) (*entity.User, error) {
	hashPassword, err := passwordHash.GetHashPassword(password)
	if err != nil {
		return nil, err
	}
	loggedInUser, err := uc.userRepo.Login(username, hashPassword)
	if err != nil {
		return nil, err
	}
	if loggedInUser == nil {
		return nil, ErrBadCredentials
	}
	return loggedInUser, nil
}

func (uc *UserUseCaseApp) Register(username, password string) (*entity.User, error) {
	loggedInUser, err := uc.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if loggedInUser != nil {
		return nil, ErrUserAlreadyExists
	}
	hashPassword, err := passwordHash.GetHashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser, err := uc.userRepo.Register(username, hashPassword)
	if err != nil || newUser == nil {
		return nil, err
	}
	return newUser, nil
}

func (uc *UserUseCaseApp) GetUserRole(userID uint64) (string, error) {
	role, err := uc.userRepo.GetUserRole(userID)
	if err != nil {
		return "", err
	}
	if role == "" {
		return "", ErrNoUser
	}
	return role, nil
}
