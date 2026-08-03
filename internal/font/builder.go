package font

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

// BuildASCII function  INFO:  Builds the user's text into ASCII art
func BuildASCII(fontName string, str string) {
	font, err := LoadFontFile(fontName)
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}

	glyphHeight := len(font["a"])
	var finalString strings.Builder

	for row := range glyphHeight {
		for _, v := range str {
			glyph, ok := font[string(v)]
			if !ok {
				color.Red("Couldn't find glyph for %s: Skipping character", string(v))
				continue
			}

			finalString.WriteString(glyph[row])
		}

		finalString.WriteString("\n")
	}

	fmt.Println(finalString.String())
}
