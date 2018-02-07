package assets

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/faiface/pixel"
)

var (
	Background     string
	BackgroundLogo string

	Grass            string
	GrassInnerCorner string
	GrassOuterCorner string
	GrassSide        string

	Dirt            string
	DirtInnerCorner string
	DirtOuterCorner string
	DirtSide        string

	DirtDark            string
	DirtDarkInnerCorner string
	DirtDarkOuterCorner string
	DirtDarkSide        string

	Sand            string
	SandInnerCorner string
	SandOuterCorner string
	SandSide        string

	Water            string
	WaterInnerCorner string
	WaterOuterCorner string
	WaterSide        string

	DeepWater            string
	DeepWaterInnerCorner string
	DeepWaterOuterCorner string
	DeepWaterSide        string

	IronOre   string
	CopperOre string
	Coal      string
	Stone     string

	BrownFluff1 string
	BrownFluff2 string
	BrownFluff3 string
	BrownFluff4 string

	PlayerBasicIdle      string
	PlayerBasicIdleColor string
	PlayerBasicRun       string
	PlayerBasicRunColor  string
)

// Setup generates all the absolute resource URLs
func Setup(gameDir string) {

	// Menu Resources
	Background = filepath.Join(gameDir, "data", "core", "graphics", "background-image.jpg")
	BackgroundLogo = filepath.Join(gameDir, "data", "core", "graphics", "background-image-logo.png")

	// Terrain Resources
	Grass = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "grass", "grass1.png")
	GrassInnerCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "grass", "grass-inner-corner.png")
	GrassOuterCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "grass", "grass-outer-corner.png")
	GrassSide = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "grass", "grass-side.png")

	Dirt = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "dirt", "dirt1.png")
	DirtInnerCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "dirt", "dirt-inner-corner.png")
	DirtOuterCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "dirt", "dirt-outer-corner.png")
	DirtSide = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "dirt", "dirt-side.png")

	DirtDark = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "dirt-dark", "dirt-dark1.png")
	DirtDarkInnerCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "dirt-dark", "dirt-dark-inner-corner.png")
	DirtDarkOuterCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "dirt-dark", "dirt-dark-outer-corner.png")
	DirtDarkSide = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "dirt-dark", "dirt-dark-side.png")

	Sand = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "sand", "sand1.png")
	SandInnerCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "sand", "sand-inner-corner.png")
	SandOuterCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "sand", "sand-outer-corner.png")
	SandSide = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "sand", "sand-side.png")

	Water = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "water", "water1.png")
	WaterInnerCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "water", "water-inner-corner.png")
	WaterOuterCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "water", "water-outer-corner.png")
	WaterSide = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "water", "water-side.png")

	DeepWater = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "deepwater", "deepwater1.png")
	DeepWaterInnerCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "deepwater", "deepwater-inner-corner.png")
	DeepWaterOuterCorner = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "deepwater", "deepwater-outer-corner.png")
	DeepWaterSide = filepath.Join(gameDir, "data", "base", "graphics", "terrain", "deepwater", "deepwater-side.png")

	// Decorative Terrain Resources
	BrownFluff1 = filepath.Join(gameDir, "data", "base", "graphics", "decorative", "brown-fluff", "brown-fluff-01.png")
	BrownFluff2 = filepath.Join(gameDir, "data", "base", "graphics", "decorative", "brown-fluff", "brown-fluff-02.png")
	BrownFluff3 = filepath.Join(gameDir, "data", "base", "graphics", "decorative", "brown-fluff", "brown-fluff-03.png")
	BrownFluff4 = filepath.Join(gameDir, "data", "base", "graphics", "decorative", "brown-fluff", "brown-fluff-04.png")

	// Ingame Resources
	IronOre = filepath.Join(gameDir, "data", "base", "graphics", "entity", "iron-ore", "iron-ore.png")
	CopperOre = filepath.Join(gameDir, "data", "base", "graphics", "entity", "copper-ore", "copper-ore.png")
	Coal = filepath.Join(gameDir, "data", "base", "graphics", "entity", "coal", "coal.png")
	Stone = filepath.Join(gameDir, "data", "base", "graphics", "entity", "stone", "stone.png")

	// Player Animations
	PlayerBasicIdle = filepath.Join(gameDir, "data", "base", "graphics", "entity", "player", "player-basic-idle.png")
	PlayerBasicIdleColor = filepath.Join(gameDir, "data", "base", "graphics", "entity", "player", "player-basic-idle-color.png")
	PlayerBasicRun = filepath.Join(gameDir, "data", "base", "graphics", "entity", "player", "player-basic-run.png")
	PlayerBasicRunColor = filepath.Join(gameDir, "data", "base", "graphics", "entity", "player", "player-basic-run-color.png")

}

// LoadPicture ...
func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// LoadSound ...
func LoadSound(path string) (*os.File, error) {

	var file *os.File
	var err error

	file, err = os.Open(path)
	if err != nil {
		return file, err
	}

	return file, nil

}
