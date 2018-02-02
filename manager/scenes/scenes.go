package scenes

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/peterbooker/factorigo/manager/scenes/game"
	"github.com/peterbooker/factorigo/manager/scenes/menu"
)

// Scene ...
type Scene interface {
	Setup(win *pixelgl.Window)
	Render(win *pixelgl.Window) string
}

func GetMenuScene() Scene {
	return menu.New()
}

func GetGameScene() Scene {
	return game.New()
}
