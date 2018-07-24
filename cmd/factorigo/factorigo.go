package main

import (
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/peterbooker/factorigo/assets"
	"github.com/peterbooker/factorigo/config"
	"github.com/peterbooker/factorigo/manager"
)

func main() {
	pixelgl.Run(run)
}

func run() {

	if len(os.Args) < 1 {
		log.Fatalln("Factorigo requires the install path of Factorio.")
	}

	gameDir := os.Args[1]

	// Setup Global Config
	config.Setup(gameDir)
	assets.Setup(gameDir)

	// Get Primary Monitor
	monitor := pixelgl.PrimaryMonitor()
	w, h := monitor.Size()

	cfg := pixelgl.WindowConfig{
		Title:       "Factorigo",
		Bounds:      pixel.R(0, 0, w, h),
		Resizable:   true,
		Undecorated: false,
		VSync:       false,
	}

	fullscreen := false

	if fullscreen {
		cfg.Monitor = monitor
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)
	win.SetCursorVisible(true)

	manager.Start(win)

	manager.Render(win)

}
