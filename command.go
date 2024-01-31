package main

import (
	"fmt"
	"os"
)

func callbackHelp() {
	helpInfo := `Welcome to the Pokedex help menu!
Here are your available commands:`
	fmt.Println(helpInfo)

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
}

func callbackExit() {
	os.Exit(0)
}
