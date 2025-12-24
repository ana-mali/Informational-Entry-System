package commands

import(
	"fmt"
	"time"
	"NotesApp/utilities"
	"NotesApp/models"
)

func Add(text string) error{
	notes, err := utilities.LoadNotes()
	if err !=nil{
		return err
	}
	newID :=utilities.NextID(notes)
	note := models.Note{
		ID: newID,
		Text: text,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	notes = append(notes, note)

	if err := utilities.SaveNotes(notes); err != nil {
		return err
	}
	fmt.Println("Note added with ID:", newID)
	return nil
}