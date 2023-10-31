package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

// Make sure that (SCREEN_HEIGHT - MENU_HEIGHT) / BOARD_HEIGHT is a whole number.
const BOARD_WIDTH = 50
const BOARD_HEIGHT = 50

type conway struct {
	board [BOARD_WIDTH][BOARD_HEIGHT]int
}

const widthNb = SCREEN_WIDTH / BOARD_WIDTH
const heightNb = (SCREEN_HEIGHT - MENU_HEIGHT) / BOARD_HEIGHT

func (c conway) draw(renderer *sdl.Renderer) {
	for i := int32(0); i < BOARD_WIDTH; i++ {
		for j := int32(0); j < BOARD_HEIGHT; j++ {
			rect := sdl.Rect{X: i * widthNb, Y: (j * heightNb) + MENU_HEIGHT, W: widthNb, H: heightNb}

			if c.board[i][j] == 0 {
				renderer.SetDrawColor(255, 255, 255, 255)
				renderer.FillRect(&rect)
			} else {
				renderer.SetDrawColor(0, 0, 0, 255)
				renderer.FillRect(&rect)
			}

			// Draw edges
			renderer.SetDrawColor(32, 0, 128, 255)
			// Vertical lines
			vertRect := sdl.Rect{X: i*(SCREEN_WIDTH/BOARD_WIDTH) - 1, Y: MENU_HEIGHT, W: 2, H: SCREEN_HEIGHT - MENU_HEIGHT}
			renderer.FillRect(&vertRect)
			// Horizontal Lines
			horRect := sdl.Rect{X: i, Y: j*((SCREEN_HEIGHT-MENU_HEIGHT)/BOARD_HEIGHT) - 1 + MENU_HEIGHT, W: SCREEN_WIDTH, H: 2}
			renderer.FillRect(&horRect)
		}
	}
}

func (c *conway) nextFrame() {
	tmpTable := c.board

	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			// do we die ?
			var nbOfNeighbour = 0

			// Top left
			if i-1 >= 0 && j-1 >= 0 && c.board[i-1][j-1] == 1 {
				nbOfNeighbour++
			}
			// Top middle
			if j-1 >= 0 && c.board[i][j-1] == 1 {
				nbOfNeighbour++
			}
			// Top Right
			if i+1 <= BOARD_WIDTH-1 && j-1 >= 0 && c.board[i+1][j-1] == 1 {
				nbOfNeighbour++
			}
			// Middle Right
			if i+1 <= BOARD_WIDTH-1 && c.board[i+1][j] == 1 {
				nbOfNeighbour++
			}
			// Bottom right
			if i+1 <= BOARD_WIDTH-1 && j+1 <= BOARD_HEIGHT-1 && c.board[i+1][j+1] == 1 {
				nbOfNeighbour++
			}
			// Botom Middle
			if j+1 <= BOARD_WIDTH-1 && c.board[i][j+1] == 1 {
				nbOfNeighbour++
			}
			// Bottom Left
			if i-1 >= 0 && j+1 <= BOARD_HEIGHT-1 && c.board[i-1][j+1] == 1 {
				nbOfNeighbour++
			}
			// Middle Left
			if i-1 >= 0 && c.board[i-1][j] == 1 {
				nbOfNeighbour++
			}

			// cell is alive
			if c.board[i][j] == 1 {
				if nbOfNeighbour < 2 {
					tmpTable[i][j] = 0
				} else if nbOfNeighbour == 2 || nbOfNeighbour == 3 {
					// We live, nothing to do
				} else if nbOfNeighbour > 3 {
					tmpTable[i][j] = 0
				}
			} else {
				// cell is dead
				if nbOfNeighbour == 3 {
					tmpTable[i][j] = 1
				}
			}
		}
	}

	c.board = tmpTable
}

func (c *conway)clear() {
  for i := 0; i < BOARD_WIDTH; i++ {
    for j := 0; j < BOARD_HEIGHT; j++ {
      c.board[i][j] = 0
    }
  }
}

func (c *conway)random() {
  for i := 0; i < BOARD_WIDTH; i++ {
    for j := 0; j < BOARD_HEIGHT; j++ {
      c.board[i][j] = rand.Intn(2)
    }
  }
}

func (c *conway) click(x int32, y int32) {
	if y < MENU_HEIGHT {
		return
	}

	xPos := x / (SCREEN_WIDTH / BOARD_WIDTH)
	yPos := (y - MENU_HEIGHT) / ((SCREEN_HEIGHT - MENU_HEIGHT) / BOARD_HEIGHT)
	if c.board[xPos][yPos] == 1 {
		c.board[xPos][yPos] = 0
	} else {
		c.board[xPos][yPos] = 1
	}
}
