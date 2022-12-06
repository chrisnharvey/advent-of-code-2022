package day06

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "day06",
		Short: "Day 6",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("please specify marker length")
			}

			if _, err := strconv.Atoi(args[0]); err != nil {
				return errors.New("marker length must be an integer")
			}

			return nil
		},
		RunE: execute,
	}
}

func execute(cmd *cobra.Command, args []string) error {
	file, err := os.Open("day06/input")

	if err != nil {
		return err
	}

	markerLength, _ := strconv.Atoi(args[0])

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanRunes)

	chars := []string{}
	pos := 0

	for scanner.Scan() {
		char := scanner.Text()

		chars = append(chars, char)
		pos++

		if len(chars) >= markerLength {
			seen := map[string]bool{}
			current := chars[len(chars)-markerLength:]

			for _, seenChar := range current {
				seen[seenChar] = true
			}

			if len(seen) == len(current) {
				fmt.Println(pos)
				os.Exit(0)
			}
		}
	}

	return nil
}
