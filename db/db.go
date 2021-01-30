package db

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

const dbURLEnvVarName = "DB_URL"

var connection *pgx.Conn

// ConnectDB ...
func ConnectDB() error {
	dbURL, dbURLExists := os.LookupEnv(dbURLEnvVarName)

	if dbURLExists == false {
		return errors.New(dbURLEnvVarName + " env variable has not been defined")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		return err
	}

	log.Println("DATABASE CONNECTED")
	connection = conn
	return nil
}

// DisconnectDB ...
func DisconnectDB() error {
	if connection == nil {
		return nil
	}

	err := connection.Close(context.Background())

	if err != nil {
		return err
	}

	return nil

}

// CreatePost ...
func CreatePost(postBody string) error {
	if connection == nil {
		return errors.New("Database connection not established")
	}

	var sqlCommand = `
		INSERT INTO post (body)
		VALUES ($1)
	`

	_, err := connection.Exec(context.Background(), sqlCommand, postBody)

	if err != nil {
		return err
	}

	return nil
}

// GetPosts ...
func GetPosts(limit, offset int) ([]Post, error) {
	if connection == nil {
		return nil, errors.New("Database connection not established")
	}

	var sqlQuery = `
		SELECT body, created_on
		FROM post
		ORDER BY created_on DESC
		LIMIT $1
		OFFSET $2
	`

	rows, err := connection.Query(context.Background(), sqlQuery, limit, offset)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var posts []Post

	for rows.Next() {
		var body string
		var createdOn time.Time

		err = rows.Scan(&body, &createdOn)
		if err != nil {
			return nil, err
		}

		posts = append(posts, Post{body, createdOn})
	}

	return posts, nil

}
