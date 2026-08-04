// Package cmd provides cmd  INFO:  Root CMD
package cmd

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/swampPr/mina-font/internal/font"
)

var rootCmd = &cobra.Command{
	Use:   "mina-font",
	Short: "A simple tool that turns basic text into an ASCII text art of the text.",
	Long:  `A simple tool that turns basic text into an ASCII text art of the text.`,
	Run: func(cmd *cobra.Command, _ []string) {
		fontName, _ := cmd.Flags().GetString("font")
		text, _ := cmd.Flags().GetString("text")
		if strings.TrimSpace(fontName) == "" && strings.TrimSpace(text) == "" {
			color.Red("You must provide both a font name and a text")
			os.Exit(1)
		}

		font.BuildASCII(fontName, text)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("font", "f", "ascii", "The FIGlet font to use")
	rootCmd.Flags().StringP("text", "t", "HELLO WORLD", "The text to use")
}
