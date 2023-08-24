package pokemonforms

import (
	"github.com/jeremyhager/code-samples/internal/pokeapi/pokemon"
	pokeutility "github.com/jeremyhager/code-samples/internal/pokeapi/utility"
)

type PokemonForm struct {
	ID           string                       `json:"id"`
	Name         string                       `json:"name"`
	Order        int                          `json:"order"`
	FormOrder    int                          `json:"form_order"`
	IsDefault    bool                         `json:"is_default"`
	IsBattleOnly bool                         `json:"is_battle_only"`
	IsMega       bool                         `json:"is_mega"`
	FormName     string                       `json:"form_name"`
	Pokemon      pokeutility.NamedAPIResource `json:"pokemon"`
	Types        []pokemon.PokemonFormType    `json:"types"`
	Sprites      PokemonFormSprites           `json:"sprites"`
	VersionGroup pokeutility.NamedAPIResource `json:"version_group"`
	Names        []pokeutility.Name           `json:"names"`
	FormNames    []pokeutility.Name           `json:"form_names"`
}

type PokemonFormSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
}
