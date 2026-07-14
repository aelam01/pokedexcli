package main

import (
	"net/http"
	"time"

	"github.com/aelam01/pokedexcli/internal/pokeapi"
	"github.com/aelam01/pokedexcli/internal/pokecache"
)

func main() {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	cache := pokecache.NewCache(10 * time.Second)
	c := pokeapi.NewClient(httpClient, cache)

	cfg := config{
		client: c,
	}

	startRepl(&cfg)
}
