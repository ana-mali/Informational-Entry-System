package services
import(
	"fmt"
	"time"
	"NotesApp/utilities"
	"NotesApp/models"
)
func AddTask(name string, priority *string, due *time.Time) (models.Task, error) {
	tasks, err := utilities.LoadTasks()
	if err != nil {
		return models.Task{}, err
	}

	newID := utilities.NextID(utilities.AsIdentifiable(tasks))

	task := models.Task{
		ID:        newID,
		Name:      name,
		CreatedAt: time.Now(),
		Priority:  priority, // *string or nil
		DueDate:   due, // *time.Time or nil 
	}

	tasks = append(tasks, task)

	if err := utilities.SaveTasks(tasks); err != nil {
		return models.Task{}, err
	}

	return task, nil
}
func DeleteTask(id int) error{
	tasks, err := utilities.LoadTasks()
	if err !=nil{
		return err
	}
	var newtasks []models.Task
	found :=false
	for _,task:=range tasks{
		if task.ID==id{
			found = true
		}else{
			newtasks=append(newtasks, task)
		}
	}
	if !found {
        return fmt.Errorf("no task found with ID %d", id)
    }
	if err := utilities.SaveTasks(newtasks); err != nil {
		return err
	}
	return nil
}
func ListTasks() ([]models.Task, error){
    tasks, err := utilities.LoadTasks()
    if err != nil {
        return nil,err
    }
    return tasks, nil
}

func EditTask(
	id int,
	name *string,
	priority *string,
	due *time.Time,
	clearPriority bool,
	clearDue bool,
) (models.Task, error) {
	tasks, err := utilities.LoadTasks()
	if err != nil {
		return models.Task{}, err
	}
	var tasktoedit *models.Task
	for i := range tasks {
		if tasks[i].ID == id {
			tasktoedit = &tasks[i]
			break
		}
	}
	if tasktoedit == nil {
		return models.Task{}, fmt.Errorf("Task not found.")
	}
	if name != nil {
		tasktoedit.Name = *name
	}
	if clearPriority {
		tasktoedit.Priority = nil
	} else if priority != nil {
		tasktoedit.Priority = priority
	}

	if clearDue {
		tasktoedit.DueDate = nil
	} else if due != nil {
		tasktoedit.DueDate = due
	}
	now := time.Now()
	tasktoedit.UpdatedAt = &now
	if err := utilities.SaveTasks(tasks); err != nil {
		return models.Task{}, err
	}
	return *tasktoedit, err
}