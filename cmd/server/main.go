package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/tu_usuario/repoGet/internal/db"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := 3306 // Default MySQL port

	conn, err := db.NewMySQLConnection(user, password, host, port)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to MySQL database 'eventtct'.")
}
