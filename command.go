package main

import (
	"fmt"
	"os"
)

func callbackHelp() error {
	helpInfo := `Welcome to the Pokedex help menu!
Here are your available commands:`
	fmt.Println(helpInfo)

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func callbackExit() error {
	os.Exit(0)
	return nil
}
