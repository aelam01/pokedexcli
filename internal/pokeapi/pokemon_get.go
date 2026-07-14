package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemon string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemon

	pokemonInfo := PokemonInfo{}

	val, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return PokemonInfo{}, err
		}
		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, err
	}

	return pokemonInfo, nil
}
