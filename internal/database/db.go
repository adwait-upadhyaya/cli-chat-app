package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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

func LoginUser(username, password string) error {
	type user struct {
		Username string
		Email    string
		Password string
	}

	loggedInUser := user{}
	err := DbConnection.QueryRow(context.Background(), "SELECT username,email,password FROM USERS WHERE username = $1", username).Scan(&loggedInUser.Username, &loggedInUser.Email, &loggedInUser.Password)

	if err != nil {
		log.Fatal("Invalid Credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(loggedInUser.Password), []byte(password))

	if err != nil {
		log.Fatal("Invalid Credentials")
	}

	fmt.Printf("Logged in succesfully")

	return nil
}
