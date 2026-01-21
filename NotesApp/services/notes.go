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
	newid := utilities.NextID(utilities.AsIdentifiable(notes))

	note := models.Note{
		ID: newid,
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
func EditNote(id int, text string) (models.Note, error) {
	notes, err := utilities.LoadNotes()
	if err != nil {
		return models.Note{}, err
	}
	var notetoedit *models.Note
	for i := range notes {
		if notes[i].ID == id {
			notetoedit = &notes[i]
			break
		}
	}
	if notetoedit == nil {
		return models.Note{}, fmt.Errorf("Note not found.")
	}
	notetoedit.Text = text
	now := time.Now()
	notetoedit.UpdatedAt = &now
	if err := utilities.SaveNotes(notes); err != nil {
		return models.Note{}, err
	}
	return *notetoedit, err
}
