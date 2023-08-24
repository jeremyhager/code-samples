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

	json.Unmarshal(pokeByte, &pokemon)
	return pokemon, nil
}
