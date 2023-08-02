package service

type Database interface {
	DoesUserExist(userName string) (userExists bool)
	CreateUser(userName, password string) (err error)
	ValidateLogin(userName, password string) (isPasswordValid bool)
	GetAllUsers() (users []string, err error)
}
