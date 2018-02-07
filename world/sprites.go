package world

import (
	"github.com/faiface/pixel"
	"github.com/peterbooker/factorigo/assets"
)

var (
	sprites *spriteManager
)

type spriteManager struct {
	waterSprite, waterSideSprite, waterInnerSprite, waterOuterSprite *pixel.Picture
	grassSprite, grassSideSprite, grassInnerSprite, grassOuterSprite *pixel.Picture
	dirtSprite, dirtSideSprite, dirtInnerSprite, dirtOuterSprite     *pixel.Picture
	sandSprite, sandSideSprite, sandInnerSprite, sandOuterSprite     *pixel.Picture
}

func SetupSprites() error {

	sprites = &spriteManager{}

	// Water Sprites
	waterSS, err := assets.LoadPicture(assets.Water)
	if err != nil {
		return err
	}
	sprites.waterSprite = &waterSS

	waterSideSS, err := assets.LoadPicture(assets.WaterSide)
	if err != nil {
		return err
	}
	sprites.waterSideSprite = &waterSideSS

	waterInnerSS, err := assets.LoadPicture(assets.WaterInnerCorner)
	if err != nil {
		return err
	}
	sprites.waterInnerSprite = &waterInnerSS

	waterOuterSS, err := assets.LoadPicture(assets.WaterOuterCorner)
	if err != nil {
		return err
	}
	sprites.waterOuterSprite = &waterOuterSS

	// Grass Sprites
	grassSS, err := assets.LoadPicture(assets.Grass)
	if err != nil {
		return err
	}
	sprites.grassSprite = &grassSS

	grassSideSS, err := assets.LoadPicture(assets.GrassSide)
	if err != nil {
		return err
	}
	sprites.grassSideSprite = &grassSideSS

	grassInnerSS, err := assets.LoadPicture(assets.GrassInnerCorner)
	if err != nil {
		return err
	}
	sprites.grassInnerSprite = &grassInnerSS

	grassOuterSS, err := assets.LoadPicture(assets.GrassOuterCorner)
	if err != nil {
		return err
	}
	sprites.grassOuterSprite = &grassOuterSS

	// Dirt Sprites
	dirtSS, err := assets.LoadPicture(assets.Dirt)
	if err != nil {
		return err
	}
	sprites.dirtSprite = &dirtSS

	dirtSideSS, err := assets.LoadPicture(assets.DirtSide)
	if err != nil {
		return err
	}
	sprites.dirtSideSprite = &dirtSideSS

	dirtInnerSS, err := assets.LoadPicture(assets.DirtInnerCorner)
	if err != nil {
		return err
	}
	sprites.dirtInnerSprite = &dirtInnerSS

	dirtOuterSS, err := assets.LoadPicture(assets.DirtOuterCorner)
	if err != nil {
		return err
	}
	sprites.dirtOuterSprite = &dirtOuterSS

	// Sand Sprites
	sandSS, err := assets.LoadPicture(assets.Sand)
	if err != nil {
		return err
	}
	sprites.sandSprite = &sandSS

	sandSideSS, err := assets.LoadPicture(assets.SandSide)
	if err != nil {
		return err
	}
	sprites.sandSideSprite = &sandSideSS

	sandInnerSS, err := assets.LoadPicture(assets.SandInnerCorner)
	if err != nil {
		return err
	}
	sprites.sandInnerSprite = &sandInnerSS

	sandOuterSS, err := assets.LoadPicture(assets.SandOuterCorner)
	if err != nil {
		return err
	}
	sprites.sandOuterSprite = &sandOuterSS

	return nil

}
