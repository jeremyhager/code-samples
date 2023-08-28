package pokemon

import (
	"encoding/json"

	"github.com/jeremyhager/code-samples/internal/pokeapi/pokeclient"
)

func Get(id string) (Pokemon, error) {
	client := pokeclient.Init("pokemon")

	pokeByte, err := client.Get(id)
	if err != nil {
		return Pokemon{}, err
	}
	var pokemon Pokemon

	err = json.Unmarshal(pokeByte, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
