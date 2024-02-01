package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func callbackHelp(cfg *config, args ...string) error {
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

func callbackExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokeman in %s:\n", locationArea.Name)
	for _, pokeman := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokeman.Pokemon.Name)
	}
	return nil
}

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

	return nil
}

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex:")
	for _, pokeman := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokeman.Name)
	}

	return nil
}
