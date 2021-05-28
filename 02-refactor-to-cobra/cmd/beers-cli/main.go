package main

import (
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "cobra-test-cli",
	}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
