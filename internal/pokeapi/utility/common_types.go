package utility

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"Language"`
}
