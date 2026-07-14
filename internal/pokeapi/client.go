package pokeapi

import (
	"net/http"

	"github.com/aelam01/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient *http.Client
	cache      *pokecache.Cache
}

func NewClient(newClient *http.Client, newCache *pokecache.Cache) Client {
	c := Client{
		httpClient: newClient,
		cache:      newCache,
	}
	return c
}
