package day03

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "day03",
		Short: "Day 3",
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: execute,
	}
}

func execute(cmd *cobra.Command, args []string) error {
	file, err := os.Open("day03/input")

	if err != nil {
		return err
	}

	groupsOf := 3

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	total := 0

	groups := [][]string{}
	group := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		half := len(line) / 2

		var usedChars []string

		compartment1 := []byte(line[:half])
		compartment2 := []byte(line[half:])

		if len(group) == groupsOf {
			// New group
			groups = append(groups, group)
			group = []string{}
		}

		group = append(group, line)

		for _, val := range compartment1 {
			char := string(val)

			if slices.Contains(usedChars, char) {
				continue
			}

			if !bytes.ContainsAny(compartment2, char) {
				continue
			}

			total += characterScore(char)

			usedChars = append(usedChars, char)

		}
	}

	groups = append(groups, group)

	badgeTotal := 0

	for _, group := range groups {
		occurenceCount := map[string]int{}

		for _, elf := range group {
			seen := []string{}

			for _, char := range strings.Split(elf, "") {
				if slices.Contains(seen, char) {
					continue
				}

				seen = append(seen, char)

				if _, ok := occurenceCount[char]; !ok {
					occurenceCount[char] = 0
				}

				occurenceCount[char]++
			}
		}

		for char, count := range occurenceCount {
			if count == groupsOf {
				badgeTotal += characterScore(char)
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Group Badge Total: %d\n", badgeTotal)

	file.Close()

	return nil
}

func characterScore(char string) int {
	alphabet := []byte("abcdefghijklmnopqrstuvwxyz")

	score := bytes.IndexAny(alphabet, strings.ToLower(char)) + 1

	if char == strings.ToUpper(char) {
		// Its caps, add 26
		score += 26
	}

	return score
}
