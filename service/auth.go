package service

import (
	"errors"
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

type Engine struct {
	PostgresDB Database
	Redis      *redis.Client
}

// signUpUser lets the user sign up with the user name and the password,
// Will return an error if the user is already present in the system.
func (e *Engine) signUpUser(userName string, password string) (err error) {
	// check if the user already exists in the system.
	userExists := e.PostgresDB.DoesUserExist(userName)
	if userExists {
		errors.New("User Already Exists")
		return
	}

	err = e.PostgresDB.CreateUser(userName, password)
	if err != nil {
		fmt.Println("Unable to create user", err)
		return
	}
	return
}

// LoginUser lets the user login with the user name and password
// Will return an error if the user does not exist or if the password is invalid.
func (e *Engine) LoginUser(userName string, password string) (token string, err error) {
	// check if the user is valid.
	userExists := e.PostgresDB.DoesUserExist(userName)
	if !userExists {
		err = errors.New("Invalid User")
		return
	}
	// Validate the user password here
	isPasswordValid := e.PostgresDB.ValidateLogin(userName, password)
	if !isPasswordValid {
		err = errors.New("Invalid Password")
		return
	}

	// generate token and initiate session.
	token, tokenErr := generateUserToken(userName)
	if tokenErr != nil {
		fmt.Println("Failed to generate token")
		return
	}

	return
}

// LogoutUser closes the user session with the user name.
func (e *Engine) LogoutUser(userName string, password string) (token string, err error) {
	// check if the user is valid.
	userExists := e.PostgresDB.DoesUserExist(userName)
	if !userExists {
		err = errors.New("Invalid User")
		return
	}

	// generate token and initiate session.
	err = InvalidateUserToken(userName)
	if err != nil {
		fmt.Println("Failed to generate token")
		return
	}

	return
}

// LogoutUser closes the user session with the user name.
func (e *Engine) GetAllUsers() (users []string, err error) {
	// check if the user is valid.
	users, err = e.PostgresDB.GetAllUsers()
	if err != nil {
		fmt.Println("Failed to fetch users")

	}
	return
}
