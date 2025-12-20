package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Note struct {
	ID        int
	Text      string
	CreatedAt time.Time
}

func newNote(Text) *Note {
	n := Note()
}

func loadNotes() ([]Note, error) {
	file, err := os.Open("data/notes.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var notes []Note
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&notes); err != nil {
		return nil, err
	}
	return notes, nil
}

func nextID(notes []Note) int {
	if len(notes)<1:
		return 1
	else:
		var maxID int = 0
		for _, note := range notes {
			if note.ID > maxID {
				maxID = note.ID
			}
		}
	return maxID+1
}
func SaveNotes(notes []models.Note) error {
	file, err := os.Create("data/notes.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(notes)
}