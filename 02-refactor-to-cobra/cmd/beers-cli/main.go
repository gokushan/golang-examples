package main

import (
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.Execute()
	rootCmd = &cobra.Command{Use: "store-cli"}
	rootCmd.AddCommand(cli.InitShopCmd())
	rootCmd.Execute()
}
