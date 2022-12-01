package main

import (
	"fmt"
	"os"

	"github.com/chrisnharvey/advent-of-code/day01"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "advent-of-code",
		Short: "Advent of code challenges",
	}

	rootCmd.AddCommand(day01.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
