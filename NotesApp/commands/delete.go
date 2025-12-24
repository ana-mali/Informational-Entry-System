package commands

import(
	"fmt"
	"NotesApp/utilities"
	"NotesApp/models"
)

func Delete(id int) error{
	notes, err := utilities.LoadNotes()
	if err !=nil{
		return err
	}
	var newnotes []models.Note
	found :=false
	for _,note:=range notes{
		if note.ID==id{
			found = true
		}
		newnotes=append(newnotes, note)
	}
	if !found {
        return fmt.Errorf("no note found with ID %d", id)
    }
	if err := utilities.SaveNotes(newnotes); err != nil {
		return err
	}
	fmt.Println("Note deleted with ID:", id)
	return nil
}