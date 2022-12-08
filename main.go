package main

import (
	"fmt"
	"os"

	"github.com/chrisnharvey/advent-of-code/day01"
	"github.com/chrisnharvey/advent-of-code/day02"
	"github.com/chrisnharvey/advent-of-code/day03"
	"github.com/chrisnharvey/advent-of-code/day04"
	"github.com/chrisnharvey/advent-of-code/day05"
	"github.com/chrisnharvey/advent-of-code/day06"
	"github.com/chrisnharvey/advent-of-code/day07"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "advent-of-code",
		Short: "Advent of code challenges",
	}

	rootCmd.AddCommand(day01.NewCommand())
	rootCmd.AddCommand(day02.NewCommand())
	rootCmd.AddCommand(day03.NewCommand())
	rootCmd.AddCommand(day04.NewCommand())
	rootCmd.AddCommand(day05.NewCommand())
	rootCmd.AddCommand(day06.NewCommand())
	rootCmd.AddCommand(day07.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
