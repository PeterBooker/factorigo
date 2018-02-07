package game

import (
	"fmt"
	"image/color"
	"math"
	"math/big"
	"time"

	"github.com/faiface/pixel"
	"github.com/peterbooker/factorigo/player"
	"github.com/peterbooker/factorigo/world"

	"github.com/faiface/pixel/pixelgl"
)

// Game ...
type Game struct {
	Title        string
	color        color.RGBA
	world        *world.World
	player       *player.Player
	dt           float64
	time         time.Time
	last         time.Time
	mousePos     pixel.Vec
	camPos       pixel.Vec
	camSpeed     float64
	camZoom      float64
	camZoomSpeed float64
	frames       int
	second       <-chan time.Time
	mapSize      int
	screen       pixel.Rect
	fpsLimit     time.Duration
	fps          <-chan time.Time
	surface      *pixelgl.Canvas
	ui           *pixelgl.Canvas
}

// New ...
func New() *Game {
	return &Game{
		Title: "Game",
	}
}

// Setup ...
func (s *Game) Setup(win *pixelgl.Window) {

	err := world.SetupSprites()
	if err != nil {
		fmt.Printf("Failed to setup Sprites: %s\n", err)
	}

	s.camPos = pixel.V(500, 500)
	s.mousePos = s.camPos
	s.camSpeed = 400.0
	s.camZoom = 1.0
	s.camZoomSpeed = 1.2
	s.frames = 0
	s.second = time.Tick(time.Second)
	s.mapSize = 300
	s.screen = win.Bounds()
	s.fpsLimit = time.Duration(120)
	s.fps = time.Tick(time.Second / s.fpsLimit)

	s.world = world.New(s.mapSize, s.mapSize)
	s.world.Chunk()

	s.player = player.New()

	s.surface = pixelgl.NewCanvas(pixel.R(0, 0, 3000, 3000))

	for _, chk := range s.world.LiveChunks {
		chk.Draw(s.surface)
	}

	s.ui = pixelgl.NewCanvas(pixel.R(0, 0, 1920, 1080))

}

// Render ...
func (s *Game) Render(win *pixelgl.Window) string {

	s.dt = time.Since(s.last).Seconds()
	s.last = time.Now()

	cam := pixel.IM.Scaled(s.camPos, s.camZoom).Moved(win.Bounds().Center().Sub(s.camPos))
	win.SetMatrix(cam)

	s.surface.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	s.mousePos = cam.Unproject(win.MousePosition())

	if win.Pressed(pixelgl.KeyLeft) || win.Pressed(pixelgl.KeyA) {
		s.camPos.X -= s.camSpeed * s.dt
		s.player.Moving.W = true
	}
	if win.Pressed(pixelgl.KeyRight) || win.Pressed(pixelgl.KeyD) {
		s.camPos.X += s.camSpeed * s.dt
		s.player.Moving.E = true
	}
	if win.Pressed(pixelgl.KeyDown) || win.Pressed(pixelgl.KeyS) {
		s.camPos.Y -= s.camSpeed * s.dt
		s.player.Moving.S = true
	}
	if win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyW) {
		s.camPos.Y += s.camSpeed * s.dt
		s.player.Moving.N = true
	}

	s.camZoom = updateCamZoom(s.camZoom, s.camZoomSpeed, win.MouseScroll().Y)

	s.player.Draw(win, s.camPos, s.dt)

	s.frames++
	select {
	case <-s.second:
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", "Factorigo", s.frames))
		s.frames = 0
	default:
	}

	<-s.fps

	return "game"

}

func updateCamZoom(zoom, camZoomSpeed, mScrollY float64) float64 {

	max := 3.00
	min := 0.50

	zoom *= math.Pow(camZoomSpeed, mScrollY)

	bZoom := big.NewFloat(zoom)
	bMin := big.NewFloat(min)
	bMax := big.NewFloat(max)

	minResult := bZoom.Cmp(bMin)

	if minResult < 0 {
		return min
	}

	maxResult := bZoom.Cmp(bMax)

	if maxResult > 0 {
		return max
	}

	return zoom

}
