package app

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"klingon-tool/app/stapi"
	stapiTypes "klingon-tool/app/stapi/types"
	"klingon-tool/app/translator"
)

func Run() {
	var input string

	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please write a star track character name: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.TrimSpace(text)
	} else {
		args := os.Args[1:]
		input = strings.Join(args, " ")
	}

	inputInKlingon, err := translator.TranslateEnglishToKlingon(input)
	if err != nil {
		fmt.Println("This name cannot translate to klingon.")
		return
	}

	species, err := getCharacterSpecies(input)
	if err != nil {
		fmt.Println("Unable to get this character species.")
		return
	}

	fmt.Println(inputInKlingon)
	fmt.Println(species)
}

func speciesArrayToSpeciesName(species []stapiTypes.Species) string {
	if len(species) == 0 {
		return "Unknown"
	}

	speciesNames := make([]string, len(species))

	for i, s := range species {
		speciesNames[i] = s.Name
	}

	count := len(speciesNames)
	if count == 2 {
		return fmt.Sprintf(
			"half '%s' and half '%s'",
			speciesNames[0], speciesNames[1],
		)
	} else if count == 3 {
		return fmt.Sprintf(
			"third '%s', third '%s' and third '%s'",
			speciesNames[0], speciesNames[1], speciesNames[2],
		)
	}

	return strings.Join(speciesNames, " ")
}

func getCharacterSpecies(name string) (string, error) {
	client := stapi.New(createHttpClient())

	searchQuery := url.Values{}
	searchQuery.Add("name", name)
	searchResponse, _, err := client.Character.Search(searchQuery)
	if err != nil {
		return "", err
	}
	if len(searchResponse.Characters) == 0 {
		return "", fmt.Errorf("stapi was unable to find the character named '%s'", name)
	}

	firstCharacter := searchResponse.Characters[0]

	query := url.Values{}
	query.Add("uid", firstCharacter.Uid)
	response, _, err := client.Character.Fetch(query)
	if err != nil {
		return "", err
	}

	return speciesArrayToSpeciesName(response.Character.CharacterSpecies), nil
}
