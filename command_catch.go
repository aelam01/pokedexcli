package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon")
	}

	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	pokemonInfo, err := cfg.client.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	catch := false
	if rand.Intn(700) > pokemonInfo.BaseExperience {
		catch = true
	}

	if catch {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.pokedex[pokemon] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
