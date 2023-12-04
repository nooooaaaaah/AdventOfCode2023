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

	// testGame := testGame()
	// fmt.Println(testGame)
	// blue, red, green := totalMovesByColor(testGame.Rounds)
	// fmt.Println(checkTotals(blue, red, green))
	// fmt.Println(blue, red, green)
	games.printGames()
	var sumOfValidGameIDs int
	for _, game := range games {
		blue, red, green := totalMovesByColor(game.Rounds)
		fmt.Printf("Game %d: blue: %d, red: %d, green: %d\n", game.ID, blue, red, green)
		if validateGame(blue, red, green) {
			sumOfValidGameIDs += game.ID
		}
	}
	fmt.Println(sumOfValidGameIDs)

	// blue, red, green := totalMovesByColor(games[0].Rounds)
	// fmt.Println(checkTotals(blue, red, green))
	// fmt.Println(blue, red, green)

	// games.printGames()
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
			color := roundPart[2:]
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

func (g *Games) printGames() {
	for _, game := range *g {
		fmt.Println(game)
	}
}

func openFile(filename string) (*os.File, error) {
	return os.Open(filename)
}

// 12 red max
// 13 green max
// 14 blue max
func validateGame(blueTotal int, redTotal int, greenTotal int) bool {
	if blueTotal > 14 || redTotal > 12 || greenTotal > 13 {
		return false
	}
	return true
}

func totalMovesByColor(rounds []Round) (int, int, int) {
	var blueTotal int
	var redTotal int
	var greenTotal int
	for _, round := range rounds {
		for _, move := range round.Moves {
			switch move.Color {
			case "blue":
				blueTotal += move.Number
			case "red":
				redTotal += move.Number
			case "green":
				greenTotal += move.Number
			}
		}
	}
	return blueTotal, redTotal, greenTotal
}

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
