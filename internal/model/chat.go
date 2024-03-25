package model

import "time"

type Chat struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Username  string
	Password  string
	Email     string
}
