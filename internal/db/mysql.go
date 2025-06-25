package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// NewMySQLConnection creates and returns a new MySQL database connection.
func NewMySQLConnection(user, password, host string, port int) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/eventtct?parseTime=true", user, password, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
