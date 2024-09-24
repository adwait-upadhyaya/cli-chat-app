package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var DbConnection *pgx.Conn

func Connect() (*pgx.Conn, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", dbUsername, dbPassword, dbName)

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("COnnected to Pg DB")
	return conn, nil
}

func init() {
	var err error
	DbConnection, err = Connect()

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
}

func RegisterUser(username, email, password string) error {
	_, err := DbConnection.Exec(context.Background(), "INSERT INTO USERS (username, email, password) VALUES ($1, $2, $3)", username, email, password)

	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}

	return nil
}
