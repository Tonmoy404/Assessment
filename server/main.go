package main

import (
	"fmt"

	"github.com/Tonmoy404/Assessment/cmd"
)

func main() {
	var str string
	fmt.Println("Enter string BOTH to seed and start: ")
	fmt.Println("Enter string START to start the project without seeding: ")
	fmt.Scan(&str)

	if str == "BOTH" {
		cmd.SeedDatabase()
		cmd.Execute()
	} else {
		cmd.Execute()
	}

}
