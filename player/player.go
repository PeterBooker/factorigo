package player

import (
	"log"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/peterbooker/factorigo/assets"
)

const (
	stages = 22
	animH  = 70
	animW  = 45
)

type sprites struct {
	basicIdle, basicIdleColor, basicRun, basicRunColor *pixel.Picture
}

type Player struct {
	Anim     AnimState
	Dir      Direction
	lastPos  pixel.Vec
	Moving   *moving
	progress int
	sprites  *sprites
	dt       float64
}

type AnimState uint8

const (
	Idle AnimState = iota
	Run
	Mine
)

type Direction uint8

const (
	N Direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

type moving struct {
	N, E, S, W bool
}

// New ...
func New() *Player {

	p := &Player{
		Anim:     Idle,
		Dir:      S,
		progress: 0,
		Moving:   &moving{},
	}

	err := p.loadSprites()
	if err != nil {
		log.Fatalf("Error loading player sprites: %s\n", err)
	}

	return p

}

func (p *Player) updateDirection() Direction {

	if p.Moving.N {
		if p.Moving.W {
			return NW
		} else if p.Moving.E {
			return NE
		} else {
			return N
		}
	}

	if p.Moving.S {
		if p.Moving.W {
			return SW
		} else if p.Moving.E {
			return SE
		} else {
			return S
		}
	}

	if p.Moving.E {
		return E
	}

	if p.Moving.W {
		return W
	}

	return p.Dir

}

func (p *Player) isMoving() bool {

	if p.Moving.N || p.Moving.E || p.Moving.S || p.Moving.W {
		return true
	}

	return false
}

func (p *Player) idleRect() pixel.Rect {

	animRate := 1.0 / 10

	th := 73.00
	tw := 53.00

	sh := 584.00
	//sw := 1056.00

	minX := 0.00 + (tw * float64(p.progress))
	minY := (sh - th) - (th * float64(p.Dir))

	maxX := minX + tw
	maxY := minY + th

	rect := pixel.R(minX, minY, maxX, maxY)

	if p.progress == 21 {
		p.progress = 0
	} else {
		i := int(math.Floor(p.dt / animRate))
		if i > 0 {
			p.progress++
			p.dt = 0.00
		}
	}

	return rect

}

func (p *Player) runRect() pixel.Rect {

	animRate := 1.0 / 30

	th := 71.00
	tw := 48.00

	sh := 568.00
	//sw := 1056.00

	minX := 0.00 + (tw * float64(p.progress))
	minY := (sh - th) - (th * float64(p.Dir))

	maxX := minX + tw
	maxY := minY + th

	rect := pixel.R(minX, minY, maxX, maxY)

	if p.progress == 21 {
		p.progress = 0
	} else {
		i := int(math.Floor(p.dt / animRate))
		if i > 0 {
			p.progress++
			p.dt = 0.00
		}
	}

	return rect

}

func (p *Player) Draw(win *pixelgl.Window, pos pixel.Vec, dt float64) {

	if p.isMoving() {
		p.Anim = Run
	} else {
		p.Anim = Idle
	}

	p.Dir = p.updateDirection()
	p.Moving.N = false
	p.Moving.E = false
	p.Moving.S = false
	p.Moving.W = false

	var base *pixel.Sprite
	//var color *pixel.Sprite

	p.dt += dt

	switch p.Anim {
	case Run:
		// Run
		base = pixel.NewSprite(*p.sprites.basicRun, p.runRect())
		//color = pixel.NewSprite(*p.sprites.basicRunColor, pixel.R(0, 0, 45, 75))
	case Mine:
		// Mine
	default:
		// Idle
		base = pixel.NewSprite(*p.sprites.basicIdle, p.idleRect())
		//color = pixel.NewSprite(*p.sprites.basicIdleColor, pixel.R(0, 0, 45, 75))
	}

	base.Draw(win, pixel.IM.Moved(pos))

}

func (p *Player) loadSprites() error {

	p.sprites = &sprites{}

	IdleSS, err := assets.LoadPicture(assets.PlayerBasicIdle)
	if err != nil {
		return err
	}
	p.sprites.basicIdle = &IdleSS

	basicIdleColor, err := assets.LoadPicture(assets.PlayerBasicIdleColor)
	if err != nil {
		return err
	}
	p.sprites.basicIdleColor = &basicIdleColor

	basicRun, err := assets.LoadPicture(assets.PlayerBasicRun)
	if err != nil {
		return err
	}
	p.sprites.basicRun = &basicRun

	basicRunColor, err := assets.LoadPicture(assets.PlayerBasicRunColor)
	if err != nil {
		return err
	}
	p.sprites.basicRunColor = &basicRunColor

	return nil

}
