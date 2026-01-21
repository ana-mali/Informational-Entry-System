package utilities

import (
	"NotesApp/models"
	"encoding/json"
	"os"
)

func LoadLists() ([]models.List, error) {
	file, err := os.Open("data/lists.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lists []models.List
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&lists); err != nil {
		return nil, err
	}
	return lists, nil
}
func LoadNotes() ([]models.Note, error) {
	file, err := os.Open("data/notes.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var notes []models.Note
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&notes); err != nil {
		return nil, err
	}
	return notes, nil
}
func LoadTasks() ([]models.Task, error) {
	file, err := os.Open("data/tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var tasks []models.Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
func NextID(items []models.Identifiable) int {
	maxID := 0
	for _, item := range items {
		if item.GetID() > maxID {
			maxID = item.GetID()
		}
	}
	return maxID + 1
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
func SaveTasks(tasks []models.Task) error {
	file, err := os.Create("data/tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(tasks)
}
func SaveLists(lists []models.List) error {
	file, err := os.Create("data/lists.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(lists)
}
func AsIdentifiable[T models.Identifiable](items []T) []models.Identifiable {
	result := make([]models.Identifiable, len(items))
	for i := range items {
		result[i] = items[i]
	}
	return result
}
func NextItemID(items []models.Item) int {
	maxID := 0
	for _, item := range items {
		if item.ID > maxID {
			maxID = item.ID
		}
	}
	return maxID + 1
}