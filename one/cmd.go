package one

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "one",
		Short: "Challenge 1",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("pass number to sum")
			}

			_, err := strconv.Atoi(args[0])

			return err
		},
		RunE: execute,
	}
}

func execute(cmd *cobra.Command, args []string) error {
	topNo, _ := strconv.Atoi(args[0])

	file, err := os.Open("one/input")

	if err != nil {
		return err
	}

	activeElf := 1
	elves := map[int]int{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			activeElf++
		}

		calories, _ := strconv.Atoi(line)

		elves[activeElf] += calories
	}

	file.Close()

	elfVals := make([]int, 0, len(elves))

	for _, value := range elves {
		elfVals = append(elfVals, value)
	}

	sort.Slice(elfVals, func(i, n int) bool {
		return elfVals[i] > elfVals[n]
	})

	sum := 0

	for _, val := range elfVals[:topNo] {
		sum += val
	}

	fmt.Println(sum)

	return nil
}
