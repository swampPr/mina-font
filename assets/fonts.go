// Package assets provides assets  INFO:  Provide fonts embed
package assets

import "embed"

// Fonts  INFO:  The fonts directory embedded into the bin of the tool
//
//go:embed fonts/*.json
var Fonts embed.FS
