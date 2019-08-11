package main

import (
	"github.com/stretchr/testify/require"
	"os/exec"
	"strings"
	"testing"
)

func normalizeCliOutput(text string) string  {
	symbolsToRemoveFromOutput := []string{"\x1b[93m", "\x1b[94m"}
	for _, s := range symbolsToRemoveFromOutput {
		text = strings.ReplaceAll(text, s, "")
	}
	return strings.TrimSpace(text)
}

func TestKlingonTool(t *testing.T) {

	tests := []struct {
		name    string
		klingon string
		species string
	}{
		{"Uhura", "0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0", "Human"},
		{"Goro", "This name cannot translate to klingon.", ""},
		{"Yari", "0xF8E8 0xF8D0 0xF8E1 0xF8D7", "Mintakan"},
	}

	cmd := exec.Command("go", "build")
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {
		cmd := exec.Command("go", "run", "main.go", test.name)
		b, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		output := normalizeCliOutput(string(b))

		lines := strings.Split(output, "\n")

		require.Equal(t, test.klingon, lines[0], test)
		if test.species != "" {
			require.Equal(t, test.species, lines[1], test)
		}
	}
}
