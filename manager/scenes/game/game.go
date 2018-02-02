package game

import (
	"image/color"

	"github.com/faiface/pixel/pixelgl"
)

// Game ...
type Game struct {
	Title string
	color color.RGBA
}

// New ...
func New() *Game {
	return &Game{
		Title: "Game",
	}
}

func (s *Game) Setup(win *pixelgl.Window) {

}

func (s *Game) Render(win *pixelgl.Window) string {

	return "game"

}
