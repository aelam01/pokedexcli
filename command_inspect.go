package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon")
	}

	pokemon := args[0]
	if _, ok := cfg.pokedex[pokemon]; !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	pokemonInfo := cfg.pokedex[pokemon]

	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height:%d\n", pokemonInfo.Height)
	fmt.Printf("Weight:%d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemonInfo.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}

	return nil
}
