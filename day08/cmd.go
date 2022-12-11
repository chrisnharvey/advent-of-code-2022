package day08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "day08",
		Short: "Day 8",
		Args: func(cmd *cobra.Command, args []string) error {

			return nil
		},
		RunE: execute,
	}
}

func execute(cmd *cobra.Command, args []string) error {
	file, err := os.Open("day08/input")

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rows := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		row := []int{}

		for _, colS := range strings.Split(line, "") {
			col, _ := strconv.Atoi(colS)
			row = append(row, col)
		}

		rows = append(rows, row)
	}

	// Calculate edge
	visible := (len(rows) * 2) + ((len(rows[0]) * 2) - 4)
	largestInRow := map[int]int{}
	largestInCol := map[int]int{}

	for r := 1; r < len(rows)-1; r++ {
		for c := 1; c < len(rows[0])-1; c++ {
			// From left
			for lc := 0; lc < c; lc++ {
				if rows[r][lc] >= rows[r][c] {

				}
			}
		}
	}

	fmt.Println(visible)

	return nil
}
