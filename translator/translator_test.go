package translator

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTranslateEnglishToKlingon(t *testing.T) {
	tests := []struct{
		input string
		output string
	}{
		{"Ms Uhura", "0xF8DA 0xF8E2 0x0020 0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0"},
		{"Data", "0xF8D3 0xF8D0 0xF8E3 0xF8D0"},
		{"Picard", ""}, // c cannot translate to klingon
		{"Spock", ""}, // c and k cannot translate to klingon
		{"Worf", ""}, // f cannot translate to klingon
	}

	for _, test := range tests {
		klingon, _ := TranslateEnglishToKlingon(test.input)
		require.Equal(t, klingon, test.output, fmt.Sprintf("input: '%s'", test.input))
	}
}
