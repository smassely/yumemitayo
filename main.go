package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	pen = iota
	eraser
	max
)

type pos struct {
	x int32
	y int32
}

type grid struct {
	height     int32
	width      int32
	block_size int32
	grid_start pos
}

type card struct {
	color    rl.Color
	grid     grid
	grid_pos pos
	active   bool
	hover    bool
}

func newPos(x int32, y int32) *pos {
	return &pos{
		x: x,
		y: y,
	}
}
func newGrid(height int32, width int32, block_size int32, grid_start pos) *grid {
	return &grid{
		height:     height,
		width:      width,
		block_size: block_size,
		grid_start: grid_start,
	}
}

func newCard(grid grid, grid_pos pos) *card {
	max_x := grid.width
	max_y := grid.height

	x := grid_pos.x
	y := grid_pos.y
	if grid_pos.x > max_x-1 || grid_pos.x < 0 {
		x = max_x
	}
	if grid_pos.y > max_y-1 || grid_pos.y < 0 {
		y = max_y
	}

	return &card{
		color:    rl.Black,
		grid:     grid,
		grid_pos: *newPos(x, y),
		active:   true,
		hover:    false,
	}
}

func getCardAbsPos(c *card) *pos {
	return newPos((c.grid.grid_start.x + c.grid.block_size*c.grid_pos.x), (c.grid.grid_start.y + c.grid.block_size*c.grid_pos.y))
}

func checkCardInteraction(c *card) bool {
	rec := rl.Rectangle{float32(getCardAbsPos(c).x), float32(getCardAbsPos(c).y), float32(c.grid.block_size), float32(c.grid.block_size)}
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rec) {
		return true
	} else {
		return false
	}
}

func main() {
	// == Declarations ==
	gameGrid := newGrid(8, 8, 100, *newPos(10, 10))
	cards := make([]card, 0)

	windowWidth := (gameGrid.width)*gameGrid.block_size + gameGrid.grid_start.x*2
	windowHeight := (gameGrid.height)*gameGrid.block_size + gameGrid.grid_start.y*2
	tool := pen
	// ==Game Event Loop==

	rl.InitWindow(windowWidth, windowHeight, "gridstuff")
	defer rl.CloseWindow()
	rl.SetTargetFPS(100)
	for cardx := 0; cardx < int(gameGrid.width); cardx++ {
		for cardy := 0; cardy < int(gameGrid.height); cardy++ {
			cards = append(cards, *newCard(*gameGrid, *newPos(int32(cardx), int32(cardy))))
		}
	}
	for !rl.WindowShouldClose() {
		for card := 0; card < len(cards); card++ {
			c := &cards[card]
			if checkCardInteraction(c) {
				c.hover = true
				if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
					switch tool {
					case pen:
						c.color = rl.DarkGreen
					case eraser:
						c.color = rl.Black
					}

				}
			} else {
				c.hover = false
			}
		}

		if rl.IsKeyPressed(rl.KeyA) {
			tool++
			tool = tool % max
		}

		// ==Game Draw Loop ==
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for card := 0; card < len(cards); card++ {
			c := cards[card]
			rec := rl.NewRectangle(float32(getCardAbsPos(&c).x+(c.grid.block_size/10)*0), float32(getCardAbsPos(&c).y+int32(c.grid.block_size/10)*0), float32(c.grid.block_size), float32(c.grid.block_size))

			rl.DrawRectangleRec(rec, c.color)
			if c.hover {
				rl.DrawRectangleLinesEx(rec, 2, rl.White)
			}
		}
		rl.EndDrawing()
	}

}
