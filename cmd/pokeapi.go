/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pokeapiCmd represents the pokeapi command
var pokeapiCmd = &cobra.Command{
	Use:   "pokeapi",
	Short: "pokeapi is used to interact with the pokeapi website",
	Long: `Use for getting objects from the pokeapi website, https://pokeapi.co/.
Examples:
code-samples pokeapi pokemon --id 1
code-samples pokeapi pokemon bulbasaur
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pokeapi called")
	},
}

func init() {
	rootCmd.AddCommand(pokeapiCmd)
}
