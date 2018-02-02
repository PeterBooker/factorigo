package text

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"github.com/peterbooker/factorigo/config"
	"golang.org/x/image/font"
)

// Font File Names
const (
	TitilliumWebRegular = "TitilliumWeb-Regular.ttf"
	TitilliumWebBold    = "TitilliumWeb-Bold.ttf"
	DejaVuSansRegular   = "DejaVuSans.ttf"
	DejaVuSansBold      = "DejaVuSans-Bold.ttf"
)

var atlas *text.Atlas

// Setup ...
func Setup() {

	cfg := config.Get()

	fontPath := filepath.Join(cfg.FontDir, "TitilliumWeb-Regular.ttf")

	face, err := loadTTF(fontPath, 52)
	if err != nil {
		panic(err)
	}

	atlas = text.NewAtlas(face, text.ASCII)

}

// NewFont ...
func New(pos pixel.Vec, name string, size float64) *text.Text {

	cfg := config.Get()

	fontPath := filepath.Join(cfg.FontDir, name)

	face, err := loadTTF(fontPath, size)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)

	text := text.New(pos, atlas)

	text.LineHeight = atlas.LineHeight() * 1.5

	return text

}

// GetLineHeight ...
func GetLineHeight() float64 {
	return atlas.LineHeight()
}

func loadTTF(path string, size float64) (font.Face, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil

}
