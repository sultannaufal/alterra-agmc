package model

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var Isbn = rand.Int()

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Isbn      string    `json:"isbn"`
	Writer    string    `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
