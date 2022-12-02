package day02

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var winningMoves, matchingMoves, losingMoves, moveOutcomes map[string]string
var moveScores, outcomeScores map[string]int

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "day02",
		Short: "Day 2",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			if args[0] != "standard" && args[0] != "outcome" {
				return errors.New("only 'standard' and 'outcome' modes are supported")
			}
			return nil
		},
		RunE: execute,
	}
}

func execute(cmd *cobra.Command, args []string) error {
	initValues()
	gameMode := "standard"

	if len(args) > 0 {
		gameMode = args[0]
	}

	file, err := os.Open("day02/input")

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		game := strings.Split(line, " ")

		played := game[1]

		if gameMode == "outcome" {
			played = shouldPlay(game[0], game[1])
		}

		score := moveScores[played]

		if played == winningMoves[game[0]] {
			// Win
			score += outcomeScores["win"]
		} else if played == matchingMoves[game[0]] {
			// Draw
			score += outcomeScores["draw"]
		} else {
			// Lose
			score += outcomeScores["lose"]
		}

		totalScore += score
	}

	fmt.Println(totalScore)

	file.Close()

	return nil
}

func shouldPlay(opponentMove string, desiredOutcome string) string {
	switch moveOutcomes[desiredOutcome] {
	case "lose":
		return losingMoves[opponentMove]
	case "draw":
		return matchingMoves[opponentMove]
	case "win":
		return winningMoves[opponentMove]
	}

	panic("invalid outcome")
}

func initValues() {
	winningMoves = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	matchingMoves = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	losingMoves = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	moveOutcomes = map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}

	moveScores = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	outcomeScores = map[string]int{
		"lose": 0,
		"draw": 3,
		"win":  6,
	}
}
