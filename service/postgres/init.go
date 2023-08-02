package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

var client *sql.DB

func InitSql() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/GoJudge", user, password, connectionName))
	return
}

func GetClient(ctx context.Context) *sql.DB {
	return client
}
