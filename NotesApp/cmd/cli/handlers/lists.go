package handlers

import (
	"fmt"
	"strconv"
	"NotesApp/services"
)

func HandleLists(cmd string, args []string) {
	switch cmd {
	case "add":
		if len(args) < 1 {
			fmt.Println("Please provide list name.")
			return
		}
		list, err := services.CreateList(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("List added with ID:", list.ID)

	case "list":
		lists, err := services.GetLists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(lists) == 0 {
			fmt.Println("No lists found.")
			return
		}
		for _, l := range lists {
			fmt.Printf("[%d] %s\n", l.ID, l.Name)

			if len(l.Items) == 0 {
				fmt.Println("  (no items)")
				continue
			}

			for _, item := range l.Items {
				check := " "
				if item.Check {
					check = "x"
				}

				fmt.Printf("  - [%s] (%d) %s\n", check, item.ID, item.Text)
			}

			fmt.Println()
		}


	case "delete":
		if len(args) < 1 {
			fmt.Println("Please provide list ID.")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID must be a number.")
			return
		}
		if err := services.DeleteList(id); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("List deleted:", id)
	case "item":
		if len(args) < 1 {
			fmt.Println("Please provide an item command (add, remove).")
			return
		}

		itemCmd := args[0]

		switch itemCmd {

		case "add":
			if len(args) < 3 {
				fmt.Println("Usage: list item add <listID> \"item text\"")
				return
			}

			listID, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("List ID must be a number.")
				return
			}

			text := args[2]

			item, err := services.AddItem(listID, text)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Item added with ID:", item.ID)

		case "remove":
			if len(args) < 3 {
				fmt.Println("Usage: list item remove <listID> <itemID>")
				return
			}

			listID, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("List ID must be a number.")
				return
			}

			itemID, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Item ID must be a number.")
				return
			}

			if err := services.RemoveItem(listID, itemID); err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Item removed.")
		}
	default:
		fmt.Println("Unknown list command:", cmd)
	}
}
