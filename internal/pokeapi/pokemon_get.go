package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonGet(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name + "/"

	val, inCache := c.cache.Get(url)
	if inCache {
		return cachedPokemon(val)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}

func cachedPokemon(val []byte) (Pokemon, error) {
	pokemon := Pokemon{}
	err := json.Unmarshal(val, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, err
}
