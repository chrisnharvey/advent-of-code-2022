package main

import (
	"fmt"
	"os"

	"github.com/chrisnharvey/advent-of-code/one"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "advent-of-code",
		Short: "Advent of code challenges",
	}

	rootCmd.AddCommand(one.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
