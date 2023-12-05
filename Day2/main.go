package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Game struct {
	ID     int
	Rounds []Round
}

type Move struct {
	Number int
	Color  string
}

type Rounds []Round

type Round struct {
	ID    int
	Moves []Move
}

type Games []Game

func main() {
	var games Games
	importGamesFromFile(&games)

	var validGames int
	for _, game := range games {
		blue, red, green := totalMovesByColor(game.Rounds)
		if blue && red && green {
			validGames += game.ID
		}
		// games.printGames()
	}
	fmt.Printf("Valid Games: %d \n", validGames)
}

func importGamesFromFile(games *Games) {
	f, err := openFile("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		processLine(scanner.Text(), games)
	}
}

func processLine(line string, games *Games) {
	gameString := strings.TrimPrefix(line, "Game ")
	gameStringParts := strings.Split(gameString, ":") // [100  10 blue, 2 red; 7 green, 20 blue, 9 red; 8 red, 6 green, 2 blue]
	id, err := strconv.Atoi(gameStringParts[0])
	if err != nil {
		panic(err)
	}
	rounds := strings.Split(gameStringParts[1], ";") // [ 10 blue, 2 red  7 green, 20 blue, 9 red  8 red, 6 green, 2 blue]

	var rds []Round
	for i, round := range rounds {
		var moves []Move
		roundParts := strings.Split(round, ",") //[ 12 red 13 green 14 blue ]
		for _, roundPart := range roundParts {
			roundPart = strings.TrimSpace(roundPart)
			num, err := strconv.Atoi(strings.TrimFunc(roundPart, func(r rune) bool {
				return !unicode.IsNumber(r)
			}))
			if err != nil {
				panic(err)
			}
			color := strings.TrimSpace(roundPart[2:])
			move := Move{Number: num, Color: color}
			moves = append(moves, move)
		}
		rd := Round{ID: i, Moves: moves}
		rds = append(rds, rd)

	}
	game := Game{ID: id, Rounds: rds}
	games.addGame(game)

}
func (g *Games) addGame(game Game) {
	*g = append(*g, game)
}

func openFile(filename string) (*os.File, error) {
	return os.Open(filename)
}

// (blue, red, green)
// returns true if all moves are under the max number of moves for that color
func totalMovesByColor(rounds []Round) (bool, bool, bool) {
	blueValid, redValid, greenValid := true, true, true
	for _, round := range rounds {
		// fmt.Printf("Round: %d \n", round.ID)
		for _, move := range round.Moves {
			switch move.Color {
			case "blue":
				blueValid = blueValid && move.Number <= 14
				fmt.Printf("blueValid: %t, Number: %v \n", blueValid, move.Number)
			case "red":
				redValid = redValid && move.Number <= 12
				fmt.Printf("redValid: %t, Number: %v \n", redValid, move.Number)
			case "green":
				greenValid = greenValid && move.Number <= 13
				fmt.Printf("greenValid: %t, Number: %v \n", greenValid, move.Number)
			}
		}
	}
	return blueValid, redValid, greenValid
}

// func (g *Games) printGames() {
// 	for _, game := range *g {
// 		fmt.Println(game)
// 	}
// }

// func testGame() Game {
// 	var testGame Game
// 	testGame.ID = 98
// 	testGame.Rounds = []Round{
// 		{
// 			ID: 0,
// 			Moves: []Move{
// 				{
// 					Number: 18,
// 					Color:  "green",
// 				},
// 				{
// 					Number: 16,
// 					Color:  "red",
// 				},
// 				{
// 					Number: 1,
// 					Color:  "blue",
// 				},
// 			},
// 		},
// 		{
// 			ID: 1,
// 			Moves: []Move{
// 				{
// 					Number: 3,
// 					Color:  "red",
// 				},
// 				{
// 					Number: 2,
// 					Color:  "blue",
// 				},
// 				{
// 					Number: 20,
// 					Color:  "green",
// 				},
// 			},
// 		},
// 		{
// 			ID: 2,
// 			Moves: []Move{
// 				{
// 					Number: 1,
// 					Color:  "blue",
// 				},
// 				{
// 					Number: 20,
// 					Color:  "green",
// 				},
// 				{
// 					Number: 14,
// 					Color:  "red",
// 				},
// 			},
// 		},
// 		{
// 			ID: 3,
// 			Moves: []Move{
// 				{
// 					Number: 14,
// 					Color:  "red",
// 				},
// 				{
// 					Number: 2,
// 					Color:  "green",
// 				},
// 			},
// 		},
// 	}
// 	return testGame
// }
