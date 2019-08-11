package types

// characterSearchResponse - a partial search response from stapi
type SearchResponse struct {
	Characters []Character `json:"characters"`
	Species    []Species   `json:"species"`
}

// FetchResponse - a partial fetch response from stapi
type FetchResponse struct {
	Character FullCharacter `json:"character"`
	Species   Species       `json:"species"`
}
