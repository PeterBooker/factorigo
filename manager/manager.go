package manager

import (
	"image"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/peterbooker/factorigo/manager/scenes"
)

var (
	sceneList    map[string]scenes.Scene
	currentScene string
)

func init() {

	sceneList = make(map[string]scenes.Scene)

	sceneList["menu"] = scenes.GetMenuScene()
	sceneList["game"] = scenes.GetGameScene()

	currentScene = "menu"

}

// Start ...
func Start(win *pixelgl.Window) {

	sceneList["menu"].Setup(win)
	sceneList["game"].Setup(win)

}

func Scene(win *pixelgl.Window) {

	currentScene = sceneList[currentScene].Render(win)

}

// Render ...
func Render(win *pixelgl.Window) {

	// Set Framerate
	fps := time.Tick(time.Second / 120)

	for !win.Closed() {

		win.Clear(image.Black)

		// Close with Escape Key
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))

		Scene(win)

		win.Update()

		<-fps

	}

}
