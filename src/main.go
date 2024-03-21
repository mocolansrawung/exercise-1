package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	boardSize = 10
)

type Coordinate struct {
	X, Y int
}

type Ship struct {
	Name     string
	Size     int
	Position []Coordinate
}

type Player struct {
	Name     string
	Board    [][]bool
	Ships    []Ship
	Opponent *Player
}

type Game struct {
	Players [2]Player
	Turn    int // 0 or 1 to indicate which player's turn
}

func (p *Player) placeShips() {
	rand.Seed(time.Now().UnixNano())
	shipTypes := []struct {
		name string
		size int
	}{
		{"Carrier", 5},
		{"Battelship", 4},
		{"Cruiser", 3},
		{"Submarine", 3},
		{"Destroyer", 2},
	}

	for _, st := range shipTypes {
		var ship Ship
		ship.Name = st.name
		ship.Size = st.size
		for {
			orientation := rand.Intn(2)
			startX := rand.Intn(boardSize)
			startY := rand.Intn(boardSize)
			if p.canPlaceShip(startX, startY, orientation, st.size) {
				for i := 0; i < st.size; i++ {
					if orientation == 0 {
						ship.Position = append(ship.Position, Coordinate{startX + i, startY})
					} else {
						ship.Position = append(ship.Position, Coordinate{startX, startY + i})
					}
				}
				break
			}
		}
		p.Ships = append(p.Ships, ship)
	}
}

func (p *Player) canPlaceShip(startX, startY, orientation, size int) bool {
	if orientation == 0 && startX+size > boardSize {
		return false
	}
	if orientation == 1 && startY+size > boardSize {
		return false
	}
	for _, ship := range p.Ships {
		for i := 0; i < size; i++ {
			var x, y int
			if orientation == 0 {
				x = startX + i
				y = startY
			} else {
				x = startX
				y = startY + i
			}
			for _, pos := range ship.Position {
				if pos.X == x && pos.Y == y {
					return false
				}
			}
		}
	}
	return true
}

func (p *Player) fireShot(x, y int, opponent *Player) bool {
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		fmt.Println("Invalid coordinates. Try again.")
		return false
	}

	if p.Board[x][y] {
		fmt.Println("You've already fired these coordinates. Try again.")
		return false
	}

	p.Board[x][y] = true
	for _, ship := range opponent.Ships {
		for _, coord := range ship.Position {
			if coord.X == x && coord.Y == y {
				fmt.Println("Hit!")
				return true
			}
		}
	}
	fmt.Println("Miss!")
	return false
}

func (p *Player) allShipsSunk() bool {
	for _, ship := range p.Ships {
		for _, pos := range ship.Position {
			if !p.Board[pos.X][pos.Y] {
				return false
			}
		}
	}
	return true
}

func main() {
	var g Game
	g.Players[0].Name = "Player 1"
	g.Players[1].Name = "Player 2"

	for i := range g.Players {
		g.Players[i].Board = make([][]bool, boardSize)
		for j := range g.Players[i].Board {
			g.Players[i].Board[j] = make([]bool, boardSize)
		}
		g.Players[i].placeShips()
	}

	fmt.Println("Welcome to the game: Battleship!")

	currentPlayer := 0
	opponent := 1 - currentPlayer

	for {
		fmt.Printf("\n%s's turn:\n", g.Players[currentPlayer].Name)
		var x, y int
		fmt.Print("Enter X coordinate: ")
		fmt.Scan(&x)
		fmt.Print("Enter Y coordinate: ")
		fmt.Scan(&y)

		if g.Players[currentPlayer].fireShot(x, y, &g.Players[opponent]) {
			if g.Players[opponent].allShipsSunk() {
				fmt.Printf("%s wins!\n", g.Players[currentPlayer].Name)
				break
			}
		}
		currentPlayer, opponent = opponent, currentPlayer // switch players for next round
	}
}
