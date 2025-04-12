package main

import (
	_ "fmt"
	_ "os"
	"time"
	"math/rand"
	"github.com/fatih/color"
)

const (
	GAME_WIDTH  = 10
	GAME_HEIGHT = 20
)

var MOVE_DELAY int = 500

var inputColor = color.New(color.FgRed)
var borderColor = color.New(color.FgCyan)
var fixedColor = color.New(color.FgGreen)
var gameBoard = make([][]int, GAME_HEIGHT)

func newObject() {

	randShape := rand.Intn(5)
	xOffset := rand.Intn(GAME_WIDTH - 2)

	if randShape == 0 { // z shape
		gameBoard[0][1 + xOffset] = 1
		gameBoard[0][2 + xOffset] = 1
		gameBoard[1][0 + xOffset] = 1
		gameBoard[1][1 + xOffset] = 1
	}

	if randShape == 1 { // L shape
		gameBoard[0][1 + xOffset] = 1
		gameBoard[1][1 + xOffset] = 1
		gameBoard[2][1 + xOffset] = 1
		gameBoard[2][2 + xOffset] = 1
	}

	if randShape == 2 { // Square shape
		gameBoard[0][1 + xOffset] = 1
		gameBoard[0][2 + xOffset] = 1
		gameBoard[1][1 + xOffset] = 1
		gameBoard[1][2 + xOffset] = 1
	}

	if randShape == 3 { // T shape
		gameBoard[1][0 + xOffset] = 1
		gameBoard[1][1 + xOffset] = 1
		gameBoard[1][2 + xOffset] = 1
		gameBoard[0][1 + xOffset] = 1
	}

	if randShape == 4 { // LIIIINNNEEEE shape
		gameBoard[0][1 + xOffset] = 1
		gameBoard[1][1 + xOffset] = 1
		gameBoard[2][1 + xOffset] = 1
		gameBoard[3][1 + xOffset] = 1
	}

}

func checkIfNewObjectNeeded() {
	allLocked := true

	for y := range gameBoard {
		for x := range gameBoard[y] {
			// inputColor.Printf("%d ", gameBoard[y][x])
			if gameBoard[y][x] == 1 {
				allLocked = false
			}
			// time.Sleep(1 * time.Millisecond)
			
		}
		// inputColor.Print("\n")
	}

	// time.Sleep(500 * time.Millisecond)

	if allLocked{
		inputColor.Println("LOCKED. New Object needed.")
		time.Sleep(1000 * time.Millisecond)
		newObject()
	}
}

func setObjects() {

	for y := range gameBoard {
		for x := range gameBoard[y] {
			if gameBoard[y][x] == 1 {
				gameBoard[y][x] = 2
			}
		}
	}
}

func calculateBoard(){

	setPiece := false

	piecesToMove := []int{}

	for y := len(gameBoard) - 1; y >= 0; y-- {
		for x := len(gameBoard[y]) - 1; x >= 0; x-- {
			
			if y < len(gameBoard) - 1{

				if gameBoard[y][x] == 1{

					if gameBoard[y + 1][x] == 2 {
						setPiece = true
					} else if !setPiece {
						piecesToMove = append(piecesToMove, y, x)
					}

				}

			} else if gameBoard[y][x] == 1 {
				setObjects()
			}

		}

	}

	if setPiece == true{
		setObjects()
	} else {
		for i := 0; i < len(piecesToMove); i += 2{

			gameBoard[piecesToMove[i]][piecesToMove[i + 1]] = 0
			gameBoard[piecesToMove[i] + 1][piecesToMove[i + 1]] = 1

		}
	}

}

func drawBoard() {
	inputColor.Print("\033[3J")
	inputColor.Print("\033[2J") // Clear the screen
	inputColor.Print("\033[H")  // Move the cursor to the top left corner

	// borderColor.Println("□ □ □ □ □ □ □ □ □ □ □ □ ")
	for top := 0; top <= GAME_WIDTH + 1; top += 1{
		borderColor.Print("□ ")
	}

	borderColor.Print("\n")

	for y := range gameBoard {
		borderColor.Print("□ ")
		for x := range gameBoard[y] {

			if gameBoard[y][x] == 0 {
				inputColor.Print("  ")
			} else if gameBoard[y][x] == 1 {
				inputColor.Print("■ ")
			} else if gameBoard[y][x] == 2 {
				fixedColor.Print("■ ")
			}
		}

		borderColor.Print("□")
		inputColor.Print("\n")

	}

	// borderColor.Println("□ □ □ □ □ □ □ □ □ □ □ □ ")
	
	for bottom := 0; bottom <= GAME_WIDTH + 1; bottom += 1{
		borderColor.Print("□ ")
	}

	borderColor.Print("\n")

}

func main() {
	
	rand.Seed(time.Now().UnixNano())

	for i := range gameBoard {
		gameBoard[i] = make([]int, GAME_WIDTH)
	}

	// Initialize the game board with 0s
	for y := range gameBoard {
		for x := range gameBoard[y] {
			gameBoard[y][x] = 0
		}
	}

	drawBoard()

	for {
		checkIfNewObjectNeeded()
		calculateBoard()
		drawBoard()
		time.Sleep(time.Duration(MOVE_DELAY) * time.Millisecond)
	}
}
