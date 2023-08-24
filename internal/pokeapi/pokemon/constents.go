package pokemon

import pokeutility "github.com/jeremyhager/code-samples/internal/pokeapi/utility"

type Pokemon struct {
	ID                     int                            `json:"id"`
	Name                   string                         `json:"name"`
	BaseExperiance         int                            `json:"base_experience"`
	Height                 int                            `json:"height"`
	IsDefault              bool                           `json:"is_default"`
	Order                  int                            `json:"order"`
	Weight                 int                            `json:"weight"`
	Abilities              []PokemonAbility               `json:"abilities"`
	Forms                  pokeutility.NamedAPIResource   `json:"forms"`
	GameIndices            []pokeutility.VersionGameIndex `json:"game_indices"`
	HeldItems              []PokemonHeldItem              `json:"held_items"`
	LocationAreaEncounters []string                       `json:"location_area_encounters"`
	Moves                  []PokemonMove                  `json:"moves"`
	PastTypes              []PokemonTypePast              `json:"past_types"`
	Sprites                PokemonSprites                 `json:"sprites"`
	Species                pokeutility.NamedAPIResource   `json:"species"`
	Stats                  []PokemonStat                  `json:"stats"`
	Types                  []PokemonType                  `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool                         `json:"is_hidden"`
	Slot     int                          `json:"slot"`
	Ability  pokeutility.NamedAPIResource `json:"ability"`
}

type PokemonType struct {
	Slot int                          `json:"slot"`
	Type pokeutility.NamedAPIResource `json:"type"`
}

type PokemonFormType struct {
	Slot int                          `json:"slot"`
	Type pokeutility.NamedAPIResource `json:"type"`
}

type PokemonTypePast struct {
	Generation pokeutility.NamedAPIResource `json:"generation"`
	Types      []PokemonType                `json:"types"`
}

type PokemonHeldItem struct {
	Name           pokeutility.NamedAPIResource `json:"item"`
	VersionDetails []PokemonHeldItemVersion     `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version pokeutility.NamedAPIResource `json:"version"`
	Rarity  int                          `json:"rarity"`
}

type PokemonMove struct {
	Move                pokeutility.NamedAPIResource `json:"move"`
	VersionGroupDetails PokemonMoveVersion           `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod pokeutility.NamedAPIResource `json:"move_learn_method"`
	VersionGroup    pokeutility.NamedAPIResource `json:"version_group"`
	LevelLearnedAt  pokeutility.NamedAPIResource `json:"level_learned_at"`
}

type PokemonStat struct {
	Stat     pokeutility.NamedAPIResource `json:"stat"`
	Effort   int                          `json:"effort"`
	BaseStat int                          `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}
