package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	redis "github.com/redis/go-redis/v9"
	"github.com/usmangani1/gesprach/service"
)

var _ service.Database = &database{}

type database struct {
	db    *sql.DB
	cache *redis.Client
}

func (db *database) CreateUser(userName, password string) (err error) {
	_, err = db.db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", userName, password)
	if err != nil {
		fmt.Println("Unable to create user", err)
	}
	return
}

func (db *database) DoesUserExist(userName string) (userExists bool, err error) {
	err = db.db.QueryRow("SELECT username FROM users WHERE username=?", userName).Scan(&userName)
	if err != nil && err == sql.ErrNoRows {
		userExists = true
	}
	return
}

func (db *database) ValidateLogin(userName, password string) (isValidPassword bool, err error) {
	var expectedPassword string
	err = db.db.QueryRow("SELECT username, password FROM users WHERE username=?", userName).Scan(&userName, &expectedPassword)
	if expectedPassword == password {
		isValidPassword = true
	}
	return
}

func (db *database) generateUserToken(ctx context.Context, userName string) (token string, err error) {
	token = uuid.New().String()
	err = db.cache.Set(ctx, userName, token, 300).Err()
	if err != nil {
		return "", err
	}
	return
}

func (db *database) InvalidateUserToken(ctx context.Context, userName string) (token string, err error) {
	err = db.cache.Del(ctx, userName).Err()
	if err != nil {
		return "", err
	}
	return
}

func (db *database) GetAllUsers() (users []string, err error) {
	userInfo, err := db.db.QueryContext(context.Background(), "SELECT username, password FROM users")
	if err != nil {
		return
	}
	defer userInfo.Close()
	for userInfo.Next() {
		var user string
		err = userInfo.Scan(&user)
		users = append(users, user)
	}
	return
}
