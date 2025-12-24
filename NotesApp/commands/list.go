package commands

import 	(
	"NotesApp/utilities"
	"fmt"
)
func List() error {
    notes, err := utilities.LoadNotes()
    if err != nil {
        return err
    }

    if len(notes) == 0 {
        fmt.Println("No notes found.")
        return nil
    }

    for _, note := range notes {
        fmt.Printf("[%d] %s\n", note.ID, note.Text)
    }

    return nil
}