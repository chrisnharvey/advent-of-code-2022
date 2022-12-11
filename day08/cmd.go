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

	for r := 1; r < len(rows)-1; r++ {
		for c := 1; c < len(rows[0])-1; c++ {
			// From top
			if isVisibleFromTop(rows, r, c) {
				visible++
				continue
			}

			// From left
			if isVisibleFromLeft(rows, r, c) {
				visible++
				continue
			}

			// From bottom
			if isVisibleFromBottom(rows, r, c) {
				visible++
				continue
			}

			// From right
			if isVisibleFromRight(rows, r, c) {
				visible++
				continue
			}
		}
	}

	fmt.Println(visible)

	return nil
}

func isVisibleFromTop(rows [][]int, r int, c int) bool {
	for tr := 0; tr < r; tr++ {
		if rows[tr][c] >= rows[r][c] {
			return false
		}
	}

	return true
}

func isVisibleFromLeft(rows [][]int, r int, c int) bool {
	for lc := 0; lc < c; lc++ {
		if rows[r][lc] >= rows[r][c] {
			return false
		}
	}

	return true
}

func isVisibleFromBottom(rows [][]int, r int, c int) bool {
	for br := len(rows) - 1; br > r; br-- {
		if rows[br][c] >= rows[r][c] {
			return false
		}
	}

	return true
}

func isVisibleFromRight(rows [][]int, r int, c int) bool {
	for rc := len(rows[0]) - 1; rc > c; rc-- {
		if rows[r][rc] >= rows[r][c] {
			return false
		}
	}

	return true
}
