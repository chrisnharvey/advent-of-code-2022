package day05

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "day05",
		Short: "Day 5",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("please pass a CrateMover model")
			}

			if args[0] != "9000" && args[0] != "9001" {
				return errors.New("invalid CrateMover model")
			}

			return nil
		},
		RunE: execute,
	}
}

func execute(cmd *cobra.Command, args []string) error {
	crateMoverModel := args[0]

	file, err := os.Open("day05/input")

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	maxCols := 0
	rows := []map[int]string{}
	moves := [][3]int{}

	scanMode := "grid"

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Blank line, we're switching scan mode
			scanMode = "moves"

			continue
		}

		if scanMode == "grid" {
			row := parseGridLine(line)
			rows = append(rows, row)

			if len(rows) > maxCols {
				maxCols = len(row)
			}

			continue
		}

		if scanMode == "moves" {
			moves = append(moves, parseMoveLine(line))
		}
	}

	cols := map[int][]string{}

	// Build cols
	for i := len(rows); i > 0; i-- {
		for col, val := range rows[i-1] {
			cols[col] = append(cols[col], val)
		}
	}

	if crateMoverModel == "9000" {
		// part 1
		for _, move := range moves {
			for i := 0; i < move[0]; i++ {
				cols[move[2]] = append(cols[move[2]], cols[move[1]][len(cols[move[1]])-1])
				cols[move[1]] = cols[move[1]][:len(cols[move[1]])-1]
			}
		}

		for i := 1; i <= len(cols); i++ {
			col := cols[i]
			fmt.Print(col[len(col)-1])
		}

		return nil
	}

	if crateMoverModel == "9001" {

		// part 2
		for _, move := range moves {
			cols[move[2]] = append(cols[move[2]], cols[move[1]][len(cols[move[1]])-move[0]:]...)
			cols[move[1]] = cols[move[1]][:len(cols[move[1]])-move[0]]
		}

		for i := 1; i <= len(cols); i++ {
			col := cols[i]
			fmt.Print(col[len(col)-1])
		}
	}

	return nil
}

func parseMoveLine(line string) [3]int {
	exp := regexp.MustCompile("[0-9]+")
	val := exp.FindAllString(line, -1)

	move, _ := strconv.Atoi(string(val[0]))
	from, _ := strconv.Atoi(string(val[1]))
	to, _ := strconv.Atoi(string(val[2]))

	return [3]int{move, from, to}
}

func parseGridLine(line string) map[int]string {

	row := map[int]string{}

	n := 1

	for i := 0; i <= len(line); i = i + 4 {
		start := i
		end := start + 4
		if end > len(line) {
			end = len(line)
		}

		col := strings.Trim(line[start:end], " ")

		if strings.Contains(col, "[") {
			col = string(col[1])

			row[n] = col
		}

		n++
	}

	return row

}
