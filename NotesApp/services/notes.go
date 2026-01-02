package services
import(
	"fmt"
	"time"
	"NotesApp/utilities"
	"NotesApp/models"
)

func AddNote(text string) (models.Note, error){
	notes, err := utilities.LoadNotes()
	if err !=nil{
		return models.Note{},err
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
		return models.Note{},err
	}
	return note,nil
}
func DeleteNote(id int) error{
	notes, err := utilities.LoadNotes()
	if err !=nil{
		return err
	}
	var newnotes []models.Note
	found :=false
	for _,note:=range notes{
		if note.ID==id{
			found = true
		}else{
			newnotes=append(newnotes, note)
		}
	}
	if !found {
        return fmt.Errorf("no note found with ID %d", id)
    }
	if err := utilities.SaveNotes(newnotes); err != nil {
		return err
	}
	return nil
}
func ListNotes() ([]models.Note, error){
    notes, err := utilities.LoadNotes()
    if err != nil {
        return nil,err
    }

    return notes, nil
}