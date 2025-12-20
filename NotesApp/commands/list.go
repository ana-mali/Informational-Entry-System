packkage commands

import "NotesApp/utilities.go"
func List(){
	notes, err := utilities.loadNotes()
	if err !=nil{
		return err
	}
	return notes
}