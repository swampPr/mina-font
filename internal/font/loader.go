// Package font provides   INFO:  Handles loading font file and building the ascii string
package font

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/fatih/color"
)

// LoadFontFile function  INFO:  Validates the inputted font name, Loads the font file into memory and returns it
func LoadFontFile(name string) (map[string][]string, error) {
	var font map[string][]string

	fontsDir := filepath.Join("fonts")
	availableFonts, err := FindAvailable(fontsDir)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(availableFonts, name) {
		color.Red("Invalid font name: %s", name)
		os.Exit(1)
	}

	fileName := fmt.Sprintf("%s.json", name)
	fontPath := filepath.Join(fontsDir, fileName)
	data, err := os.ReadFile(fontPath) //#nosec
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
func FindAvailable(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
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
