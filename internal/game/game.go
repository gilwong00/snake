package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 660
	screenHeight = 500
	gridSize     = 20 // start off with a 20x20 grid
	gameSpeed    = time.Second / 5
)

// Use Point to maintain the position of the where
// the snake is on our grid
type Point struct {
	x int
	y int
}

type Game struct {
	snake       []Point
	direction   Point
	lastUpdated time.Time
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	if time.Since(g.lastUpdated) < gameSpeed {
		return nil
	}
	g.lastUpdated = time.Now()
	g.moveSnake(g.direction)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.snake {
		vector.DrawFilledRect(
			screen,
			float32(p.x*gridSize),
			float32(p.y*gridSize),
			gridSize,
			gridSize,
			color.White,
			true,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) InitializeSnake() {
	g.snake = []Point{
		{
			// use screenWidth and divide it by the gridSize to respect
			// the width and divide by 2 so it starts in the middle.
			x: screenWidth / gridSize / 2,
			y: screenWidth / gridSize / 2,
		},
	}
	g.direction = Point{x: 1, y: 0}
}

func (g *Game) RunGame() error {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake")
	g.InitializeSnake()
	return ebiten.RunGame(g)
}

func (g *Game) moveSnake(direction Point) {
	head := g.snake[0]
	newHead := Point{
		x: head.x + direction.x,
		y: head.y + direction.y,
	}

	g.snake = append(
		// append new head
		[]Point{newHead},
		// append the rest of the snake
		g.snake[:len(g.snake)-1]...,
	)
}
