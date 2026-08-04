// Package font provides   INFO:  Handles loading font file and finding all available fonts
package font

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/fatih/color"

	"github.com/swampPr/mina-font/assets"
)

// LoadFontFile function  INFO:  Validates the inputted font name, Loads the font file into memory and returns it
func LoadFontFile(name string) (map[string][]string, error) {
	var font map[string][]string

	availableFonts, err := FindAvailable()
	if err != nil {
		return nil, err
	}

	if !slices.Contains(availableFonts, name) {
		color.Red("Invalid font name: %s", name)
		os.Exit(1)
	}

	fontPath := fmt.Sprintf("fonts/%s.json", name)
	data, err := assets.Fonts.ReadFile(fontPath) //#nosec
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &font)
	if err != nil {
		return nil, err
	}

	return font, err
}

// FindAvailable function  INFO:  Gets all available fonts in the fonts dir
func FindAvailable() ([]string, error) {
	entries, err := assets.Fonts.ReadDir("fonts")
	if err != nil {
		return nil, err
	}

	availableFonts := make([]string, 0, len(entries))

	for _, entry := range entries {
		sanitized, _, _ := strings.Cut(entry.Name(), ".json")
		availableFonts = append(availableFonts, sanitized)
	}

	return availableFonts, err
}
