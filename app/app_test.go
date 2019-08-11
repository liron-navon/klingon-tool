package app

import (
	"github.com/stretchr/testify/require"
	"klingon-tool/app/stapi/types"
	"testing"
)

func TestGetCharacterSpecies(t *testing.T) {
	tests := []struct {
		name    string
		species string
	}{
		{"Seven of Nine", "Human"},
		{"Data", "Unknown"},
		{"Phlox", "Denobulan"},
	}

	for _, test := range tests {
		species, err := getCharacterSpecies(test.name)
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, test.species, species, test)
	}
}

func TestSpeciesArrayToSpeciesName(t *testing.T) {
	tests := []struct {
		arr  []types.Species
		name string
	}{
		{[]types.Species{}, "Unknown"},
		{[]types.Species{{Name: "A"}}, "A"},
		{[]types.Species{{Name: "A"}, {Name: "B"}}, "half 'A' and half 'B'"},
		{[]types.Species{{Name: "A"}, {Name: "B"}, {Name: "C"}}, "third 'A', third 'B' and third 'C'"},
		{[]types.Species{{Name: "A"}, {Name: "B"}, {Name: "C"}, {Name: "D"}}, "A B C D"},
	}

	for _, test := range tests {
		speciesName := speciesArrayToSpeciesName(test.arr)
		require.Equal(t, test.name, speciesName, test)
	}
}
