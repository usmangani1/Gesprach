package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	redis "github.com/redis/go-redis/v9"
	"github.com/usmangani1/gesprach/service"
)

type signUpUserRequest struct {
	userName string
	password string
}

type LoginUserRequest struct {
	userName string
	password string
}

type LogoutUserRequest struct {
	userName string
}

func handleUser(w http.ResponseWriter, r *http.Request) {

	engine := &service.Engine{
		PostgresDB: InitSql(),
		Redis:      InitRedis(),
	}

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		if strings.Contains(r.URL.Path, "user") {
			req := new(signUpUserRequest)
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				fmt.Println("Failed to decode request", err)
				return
			}
			return engine.signUpUser(req.userName, req.password)
		} else if strings.Contains(r.URL.Path, "login") {
			// handle login here
			req := new(LoginUserRequest)
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				fmt.Println("Failed to decode request", err)
				return
			}
			return engine.LoginUser(req.userName, req.password)
		} else if strings.Contains(r.URL.Path, "logout") {
			// handle logout here
			req := new(LogoutUserRequest)
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				fmt.Println("Failed to decode request", err)
				return
			}
			return engine.LogoutUser(req.userName)
		}
	case "GET":
		// get all the users here
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func InitSql() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/GoJudge", "user", "password", "connectionName"))
	return
}

func InitRedis() (client *redis.NewClient) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return
}
