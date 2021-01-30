package db

import "time"

// Post ...
type Post struct {
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"createdOn"`
}
