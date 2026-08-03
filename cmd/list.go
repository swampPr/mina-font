package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/swampPr/mina-font/internal/font"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available fonts",
	Long:  `List all available fonts`,
	Run: func(_ *cobra.Command, _ []string) {
		fontsDir := filepath.Join("fonts")
		fmt.Println(font.FindAvailable(fontsDir))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
