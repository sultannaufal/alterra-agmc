package models

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var Isbn = rand.Int()

type Book struct {
	ID        int
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	Writer    string `json:"writer"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
