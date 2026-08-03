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
	// Uncomment the following line if your bare application
	// has an action associated with it:
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
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ascii-font-gen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("font", "f", " ", "The FIGlet font to use")
	rootCmd.Flags().StringP("text", "t", " ", "The text to use")
}
