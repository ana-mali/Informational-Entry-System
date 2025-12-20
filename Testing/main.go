package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "help" || os.Args[1] == "--help" {
		fmt.Println(`This program is for Testing simple go Features
The first step is a greeter/argument inspector, where if there is no arguments will ask for your name to greet. Otherwise it show all your arguments
The next is a simple addition, using client input where you will be asked to type a number and press enter to continue
		`)
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Please type your name")
		var name string
		fmt.Scan(&name)
		fmt.Println("Hello ", name)
	} else {
		for i := 0; i < len(os.Args); i++ {
			fmt.Println(i)
			fmt.Println(os.Args[i])
		}
	}
	var x int
	var y int
	fmt.Println("Type one number to add: ")
	fmt.Scan(&x)
	fmt.Println("Type another number to add together: ")
	fmt.Scan(&y)
	var z int = x + y
	fmt.Printf("%d + %d = %d", x, y, z)
}
