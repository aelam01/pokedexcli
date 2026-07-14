package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aelam01/pokedexcli/internal/pokeapi"
)

type config struct {
	client   pokeapi.Client
	next     *string
	previous *string
	pokedex  map[string]pokeapi.PokemonInfo
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Get the list of pokemon in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
