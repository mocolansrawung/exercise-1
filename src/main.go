package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	boardSize = 5
	numShips  = 3
)

type Coordinate struct {
	X, Y int
}

type Ship struct {
	Coordinates []Coordinate
}

type Game struct {
	Board [boardSize][boardSize]bool
	Ships []Ship
}

func (g *Game) placeShips() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numShips; i++ {
		var ship Ship
		for j := 0; j < 3; j++ {
			x := rand.Intn(boardSize)
			y := rand.Intn(boardSize)
			ship.Coordinates = append(ship.Coordinates, Coordinate{x, y})
		}
		g.Ships = append(g.Ships, ship)
	}
}

func (g *Game) fireShot(x, y int) bool {
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		fmt.Println("Invalid coordinates. Try again.")
		return false
	}

	if g.Board[x][y] {
		fmt.Println("You've already fired these coordinates. Try again.")
		return false
	}

	g.Board[x][y] = true
	for _, ship := range g.Ships {
		for _, coord := range ship.Coordinates {
			if coord.X == x && coord.Y == y {
				fmt.Println("Hit!")
				return true
			}
		}
	}
	fmt.Println("Miss!")
	return false
}

func (g *Game) printBoard() {
	fmt.Println("  0 1 2 3 4")
	for i := 0; i < boardSize; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < boardSize; j++ {
			if g.Board[i][j] {
				fmt.Print("X ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var g Game
	g.placeShips()

	fmt.Println("Welcome to the game: Battelship!")
	fmt.Println("Try to sink all the ships!")

	for {
		fmt.Println()
		g.printBoard()

		var x, y int
		fmt.Print("Enter X coordinate: ")
		fmt.Scan(&x)
		fmt.Print("Enter Y coordinate: ")
		fmt.Scan(&y)

		if g.fireShot(x, y) {
			allSunk := true
			for _, ship := range g.Ships {
				for _, coord := range ship.Coordinates {
					if !g.Board[coord.X][coord.Y] {
						allSunk = false
						break
					}
				}
				if !allSunk {
					break
				}
			}
			if allSunk {
				fmt.Println("Congratulations! You sunk all the ships!")
				break
			}
		}
	}
}
