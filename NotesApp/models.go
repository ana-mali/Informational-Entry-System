package main

import "time"

type Note struct {
	ID        int
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}