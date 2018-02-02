package menu

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/peterbooker/factorigo/config"
	"github.com/peterbooker/factorigo/text"
)

const (
	bHeight = 120
	bWidth  = 300
)

// Menu ...
type Menu struct {
	Title   string
	color   color.RGBA
	buttons []*Button
}

// Button ...
type Button struct {
	title  string
	bounds pixel.Rect
	color  color.RGBA
}

// New ...
func New() *Menu {
	return &Menu{
		Title: "Menu",
	}
}

// Setup ...
func (s *Menu) Setup(win *pixelgl.Window) {

	center := win.Bounds().Center()
	pos := center.Sub(pixel.V(bWidth/2, 0))

	// Create Buttons
	buttons := []string{"Start Game", "Quit"}

	for _, title := range buttons {

		button := &Button{
			title: title,
			color: color.RGBA{50, 50, 50, 255},
			bounds: pixel.Rect{
				Min: pos.Sub(pixel.V(0, bHeight)),
				Max: pos.Add(pixel.V(bWidth, 0)),
			},
		}
		s.buttons = append(s.buttons, button)

		pos = pos.Sub(pixel.V(0, bHeight+60))

	}

}

// Render ...
func (s *Menu) Render(win *pixelgl.Window) string {

	// TODO: Position is all wrong
	// Currently works for 1920x1080, but falls apart as the resolution changes.

	cfg := config.Get()

	// Render Menu Title
	titleFont := text.New(pixel.V(50, 500), text.TitilliumWebRegular, 52.00)
	titleFont.Color = color.RGBA{255, 255, 255, 255}
	fmt.Fprintln(titleFont, cfg.Title)

	titleFont.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(titleFont.Bounds().Center()).Add(pixel.V(0, 200))))

	subTitleFont := text.New(pixel.V(50, 500), text.TitilliumWebRegular, 28.00)
	subTitleFont.Color = color.RGBA{255, 255, 255, 255}
	fmt.Fprintln(subTitleFont, cfg.Version)

	subTitleFont.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(subTitleFont.Bounds().Center()).Add(pixel.V(0, 160))))

	pos := win.MousePosition()

	for _, b := range s.buttons {
		if b.bounds.Contains(pos) {
			b.color = color.RGBA{90, 90, 90, 255}
		} else {
			b.color = color.RGBA{50, 50, 50, 255}
		}
	}

	if win.Pressed(pixelgl.MouseButtonLeft) {

		pos := win.MousePosition()

		for _, b := range s.buttons {
			if b.bounds.Contains(pos) {
				switch b.title {
				case "Start Game":
					// Start Game
					return "game"
				case "Quit":
					// Quit Game
					win.SetClosed(true)
				}
			}
		}

	}

	for _, b := range s.buttons {
		// Draw Button Box
		btn := imdraw.New(nil)
		btn.Color = b.color
		btn.Push(b.bounds.Min, b.bounds.Max)
		btn.Rectangle(0)
		btn.Draw(win)

		// Draw Button Text
		buttonFont := text.New(pixel.V(0, 0), text.TitilliumWebRegular, 36.00)
		buttonFont.Color = color.RGBA{255, 255, 255, 255}

		buttonFont.Dot.X -= buttonFont.BoundsOf(b.title).W() / 2
		fmt.Fprintln(buttonFont, b.title)

		buttonFont.Draw(win, pixel.IM.Moved(b.bounds.Max.Sub(pixel.V(bWidth/2, (bHeight+30)/2))))
	}

	return "menu"

}
