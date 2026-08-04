package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/swampPr/mina-font/internal/font"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available fonts",
	Long:  `List all available fonts`,
	Run: func(_ *cobra.Command, _ []string) {
		available, err := font.FindAvailable()
		if err != nil {
			fmt.Printf("Something went wrong: %v", err)
			os.Exit(1)
		}
		fmt.Println(available)
		fmt.Printf("\nTotal available fonts: ")
		color.Green("%d\n", len(available))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
