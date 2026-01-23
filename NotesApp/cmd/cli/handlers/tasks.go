package handlers

import (
	"fmt"
	"NotesApp/services"
	"strconv"
	"time"
	"flag"
	"strings"
)

type stringSliceFlag []string

func (s *stringSliceFlag) String() string {
	return strings.Join(*s, ", ")
}

func (s *stringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}


func HandleTasks(cmd string, args []string) {
	switch cmd {

	case "add":
		if len(args) < 1 {
			fmt.Println("Please provide task name.")
			return
		}

		name := args[0]

		var priority *string
		var due *time.Time

		for i := 1; i < len(args); i++ {
			switch args[i] {

			case "--priority":
				if i+1 >= len(args) {
					fmt.Println("Missing value for --priority")
					return
				}
				p := args[i+1]
				priority = &p
				i++

			case "--due":
				if i+1 >= len(args) {
					fmt.Println("Missing value for --due")
					return
				}
				parsed, err := time.Parse("2006-01-02", args[i+1])
				if err != nil {
					fmt.Println("Due date must be YYYY-MM-DD")
					return
				}
				due = &parsed
				i++
			}
		}

		task, err := services.AddTask(name, priority, due)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task added with ID:", task.ID)

	case "list":
		tasks, err := services.ListTasks()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, t := range tasks {
			optionals := ""

			if t.Priority != nil {
				optionals += fmt.Sprintf(" | priority: %s", *t.Priority)
			}

			if t.DueDate != nil {
				optionals += fmt.Sprintf(
					" | due: %s",
					t.DueDate.Format("2006-01-02"),
				)
			}

			fmt.Printf("[%d] %s%s\n", t.ID, t.Name, optionals)
		}


	case "delete":
		if len(args) < 1 {
			fmt.Println("Please provide task ID.")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID must be a number.")
			return
		}
		if err := services.DeleteTask(id); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task deleted:", id)

	case "edit":
		if len(args) < 2 {
			fmt.Println("Usage: app tasks edit <taskID> <flag> <new data>..")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Task ID must be a number.")
			return
		}

		fs := flag.NewFlagSet("task edit", flag.ExitOnError)

		name := fs.String("name", "", "Edit task name")
		priority := fs.String("priority", "", "Edit task priority")
		dueStr := fs.String("due", "", "Edit task due date (YYYY-MM-DD)")
		var clearFields stringSliceFlag
		fs.Var(&clearFields, "clear", "Clear a field (can be repeated: priority | due)")

		if err := fs.Parse(args[1:]); err != nil {
			fmt.Println("Error parsing flags:", err)
			return
		}
		clearPriority := false
		clearDue := false

		for _, field := range clearFields {
			switch field {
			case "priority":
				clearPriority = true
			case "due":
				clearDue = true
			default:
				fmt.Printf("Unknown clear field: %s\n", field)
				return
			}
		}
		var namePtr *string
		if *name != "" {
			namePtr = name
		}

		var priorityPtr *string
		if !clearPriority && *priority != "" {
			priorityPtr = priority
		}

		var duePtr *time.Time
		if !clearDue && *dueStr != "" {
			parsed, err := time.Parse("2006-01-02", *dueStr)
			if err != nil {
				fmt.Println("Invalid date format. Use YYYY-MM-DD.")
				return
			}
			duePtr = &parsed
		}

		if namePtr == nil && priorityPtr == nil && duePtr == nil && !clearPriority && !clearDue {
			fmt.Println("No changes provided. Use flags to edit the task.")
			return
		}

		task, err := services.EditTask(
			id,
			namePtr,
			priorityPtr,
			duePtr,
			clearPriority,
			clearDue,
		)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task updated with id: ",task.ID)

	default:
		fmt.Println("Unknown task command:", cmd)
	}
}
