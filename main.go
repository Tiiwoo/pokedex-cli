package main

import (
	"fmt"
	"log"

	"github.com/Tiiwoo/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	// startRepl()
}
