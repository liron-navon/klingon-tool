package types

// FullCharacter - a partial struct from stapi's FullCharacter entity
type FullCharacter struct {
	Uid              string    `json:"uid"`
	Name             string    `json:"name"`
	CharacterSpecies []Species `json:"characterSpecies"`
}

// character - a partial struct from stapi's character entity
type Character struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}
