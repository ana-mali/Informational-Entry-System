package commands

import(
	"fmt"
	"NotesApp/utilities.go"
	"NotesApp/models.go"
)

func Add(text string) error{
	notes, err := utilities.loadNotes()
	if err !=nil{
		return err
	}
	newID :=utilities.nextID(notes)
	note := models.Note{
		ID: newID,
		Text: text,
		CreatedAt: time.Now()
		UpdatedAt: nil
	}
	notes = append(notes, note)

	if err := utilities.SaveNotes(notes); err != nil {
		return err
	}
	fmt.Println("Note added with ID:", newID)
	return nil
}