package logic

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"voter_api/data_access/repository"
	domain "voter_api/domain/user"
)

func FindVoter(idVoter string) (*domain.User, error) {
	user, err := repository.FindVoter(idVoter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, err
}

func RegisterUser(id string, username string, password string, role string) (*domain.User, error) {
	user, err := createUser(id, username, password, role)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, err
}

func createUser(id string, username string, password string, role string) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &domain.User{
		Id:             id,
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
	}
	return StoreUser(user)
}

func StoreUser(user *domain.User) (*domain.User, error) {
	err := repository.RegisterUser(user)
	if err != nil {
		return nil, fmt.Errorf("user cannot be created: %w", err)
	}
	return user, nil
}

func IsCorrectPassword(user *domain.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}
