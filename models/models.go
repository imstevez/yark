package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewSqlDB(configs map[string]interface{}) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		configs["username"].(string),
		configs["password"].(string),
		configs["protocol"].(string),
		configs["host"].(string),
		configs["port"].(string),
		configs["database"].(string),
	)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(time.Duration(configs["max_life_time"].(int)) * time.Second)
	sqlDB.SetMaxOpenConns(configs["max_open_conns"].(int))
	sqlDB.SetMaxIdleConns(configs["max_idle_conns"].(int))
	return sqlDB, nil
}
