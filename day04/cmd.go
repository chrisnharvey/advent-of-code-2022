package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "day04",
		Short: "Day 4",
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: execute,
	}
}

func execute(cmd *cobra.Command, args []string) error {
	file, err := os.Open("day04/input")

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	totalContainng := 0
	totalOverlapping := 0

	for scanner.Scan() {
		line := scanner.Text()

		pairs := strings.Split(line, ",")

		assignments := map[int][]int{}

		for idx, assignment := range pairs {
			rng := strings.Split(assignment, "-")
			start, _ := strconv.Atoi(rng[0])
			end, _ := strconv.Atoi(rng[1])

			for i := start; i <= end; i++ {
				assignments[idx] = append(assignments[idx], i)
			}
		}

		matching1 := 0
		matching2 := 0

		for _, val := range assignments[0] {
			if slices.Contains(assignments[1], val) {
				matching1++
			}
		}

		for _, val := range assignments[1] {
			if slices.Contains(assignments[0], val) {
				matching2++
			}
		}

		if matching1 != 0 || matching2 != 0 {
			totalOverlapping++
		}

		if matching1 == len(assignments[0]) || matching2 == len(assignments[1]) {
			totalContainng++
		}
	}

	fmt.Printf("Total Overlapping: %d\nTotal Containing: %d\n", totalOverlapping, totalContainng)

	return nil
}
