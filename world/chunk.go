package world

import (
	"image/color"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type chunk struct {
	area         pixel.Rect
	SurfaceType  [][]SurfaceTileType
	SurfaceShape [][]SurfaceTileShape

	sprites *spriteManager
	batches *batchManager
}

type tile struct {
	tile *pixel.Sprite
	pos  pixel.Vec
}

type tileBatch struct {
	tiles  []tile
	sprite pixel.Picture
	frames []pixel.Rect
	batch  *pixel.Batch
}

type batchManager struct {
	waterBatch      *tileBatch
	waterSideBatch  *tileBatch
	waterInnerBatch *tileBatch
	waterOuterBatch *tileBatch
	grassBatch      *tileBatch
	grassSideBatch  *tileBatch
	grassInnerBatch *tileBatch
	grassOuterBatch *tileBatch
	dirtBatch       *tileBatch
	dirtSideBatch   *tileBatch
	dirtInnerBatch  *tileBatch
	dirtOuterBatch  *tileBatch
}

func (t *tileBatch) loadFrames() {

	for x := t.sprite.Bounds().Min.X; x < t.sprite.Bounds().Max.X; x += 32 {
		for y := t.sprite.Bounds().Min.Y; y < t.sprite.Bounds().Max.Y; y += 32 {
			t.frames = append(t.frames, pixel.R(x, y, x+32, y+32))
		}
	}

}

func (t *tileBatch) loadFullTileFrames() {

	max := 8.00 * 64

	for x := 0.00; x < max; x += 64 {
		y := 512.00
		t.frames = append(t.frames, pixel.R(x, y, x+64, y+64))
	}

}

func (t *tileBatch) loadOuterTileFrames() {

	tileSize := 64.00
	width := 8.00 * 64
	height := 4.00 * 64

	xStart := 2432.00
	xMax := xStart + width
	yStart := 320.00
	yMax := yStart + height

	for x := xStart; x < xMax; x += tileSize {
		for y := yStart; y < yMax; y += tileSize {
			t.frames = append(t.frames, pixel.R(x, y, x+tileSize, y+tileSize))
		}
	}

}

func (t *tileBatch) loadInnerTileFrames() {

	tileSize := 64.00
	width := 8.00 * 64
	height := 4.00 * 64

	xStart := 3008.00
	xMax := xStart + width
	yStart := 320.00
	yMax := yStart + height

	for x := xStart; x < xMax; x += tileSize {
		for y := yStart; y < yMax; y += tileSize {
			t.frames = append(t.frames, pixel.R(x, y, x+tileSize, y+tileSize))
		}
	}

}

func (t *tileBatch) loadSideTileFrames() {

	tileSize := 64.00
	width := 8.00 * 64
	height := 4.00 * 64

	xStart := 3584.00
	xMax := xStart + width
	yStart := 320.00
	yMax := yStart + height

	for x := xStart; x < xMax; x += tileSize {
		for y := yStart; y < yMax; y += tileSize {
			t.frames = append(t.frames, pixel.R(x, y, x+tileSize, y+tileSize))
		}
	}

}

func newChunk(area pixel.Rect) *chunk {

	c := &chunk{
		area:    area,
		sprites: sprites,
	}

	c.loadBatches()

	return c

}

func (c *chunk) Build() {

	var x int
	var y int

	for x = 0; x < 100; x++ {

		for y = 0; y < 100; y++ {

			t := tile{
				pos: pixel.V(float64(x*tileSize), float64(y*tileSize)),
			}

			ttype := c.SurfaceType[x][y]
			tshape := c.SurfaceShape[x][y]

			switch ttype {

			case Water:

				switch tshape {

				case Full:

					t.tile = pixel.NewSprite(c.batches.waterBatch.sprite, c.batches.waterBatch.frames[rand.Intn(len(c.batches.waterBatch.frames))])
					c.batches.waterBatch.tiles = append(c.batches.waterBatch.tiles, t)

				case NorthSide:

					t.tile = pixel.NewSprite(c.batches.waterSideBatch.sprite, c.batches.waterSideBatch.frames[1+(rand.Intn(7)*4)])
					c.batches.waterSideBatch.tiles = append(c.batches.waterSideBatch.tiles, t)

				case EastSide:

					t.tile = pixel.NewSprite(c.batches.waterSideBatch.sprite, c.batches.waterSideBatch.frames[0+(rand.Intn(7)*4)])
					c.batches.waterSideBatch.tiles = append(c.batches.waterSideBatch.tiles, t)

				case SouthSide:

					t.tile = pixel.NewSprite(c.batches.waterSideBatch.sprite, c.batches.waterSideBatch.frames[3+(rand.Intn(7)*4)])
					c.batches.waterSideBatch.tiles = append(c.batches.waterSideBatch.tiles, t)

				case WestSide:

					t.tile = pixel.NewSprite(c.batches.waterSideBatch.sprite, c.batches.waterSideBatch.frames[2+(rand.Intn(7)*4)])
					c.batches.waterSideBatch.tiles = append(c.batches.waterSideBatch.tiles, t)

				case NorthEastInner:

					t.tile = pixel.NewSprite(c.batches.waterInnerBatch.sprite, c.batches.waterInnerBatch.frames[1+(rand.Intn(5)*4)])
					c.batches.waterInnerBatch.tiles = append(c.batches.waterInnerBatch.tiles, t)

				case SouthEastInner:

					t.tile = pixel.NewSprite(c.batches.waterInnerBatch.sprite, c.batches.waterInnerBatch.frames[0+(rand.Intn(5)*4)])
					c.batches.waterInnerBatch.tiles = append(c.batches.waterInnerBatch.tiles, t)

				case SouthWestInner:

					t.tile = pixel.NewSprite(c.batches.waterInnerBatch.sprite, c.batches.waterInnerBatch.frames[3+(rand.Intn(5)*4)])
					c.batches.waterInnerBatch.tiles = append(c.batches.waterInnerBatch.tiles, t)

				case NorthWestInner:

					t.tile = pixel.NewSprite(c.batches.waterInnerBatch.sprite, c.batches.waterInnerBatch.frames[2+(rand.Intn(5)*4)])
					c.batches.waterInnerBatch.tiles = append(c.batches.waterInnerBatch.tiles, t)

				case NorthEastOuter:

					t.tile = pixel.NewSprite(c.batches.waterOuterBatch.sprite, c.batches.waterOuterBatch.frames[1+(rand.Intn(5)*4)])
					c.batches.waterOuterBatch.tiles = append(c.batches.waterOuterBatch.tiles, t)

				case SouthEastOuter:

					t.tile = pixel.NewSprite(c.batches.waterOuterBatch.sprite, c.batches.waterOuterBatch.frames[0+(rand.Intn(5)*4)])
					c.batches.waterOuterBatch.tiles = append(c.batches.waterOuterBatch.tiles, t)

				case SouthWestOuter:

					t.tile = pixel.NewSprite(c.batches.waterOuterBatch.sprite, c.batches.waterOuterBatch.frames[3+(rand.Intn(5)*4)])
					c.batches.waterOuterBatch.tiles = append(c.batches.waterOuterBatch.tiles, t)

				case NorthWestOuter:

					t.tile = pixel.NewSprite(c.batches.waterOuterBatch.sprite, c.batches.waterOuterBatch.frames[2+(rand.Intn(5)*4)])
					c.batches.waterOuterBatch.tiles = append(c.batches.waterOuterBatch.tiles, t)

				}

			case Grass:

				if tshape != Full {

					t.tile = pixel.NewSprite(c.batches.dirtBatch.sprite, c.batches.dirtBatch.frames[rand.Intn(len(c.batches.dirtBatch.frames))])
					c.batches.dirtBatch.tiles = append(c.batches.dirtBatch.tiles, t)

				}

				switch tshape {

				case Full:

					t.tile = pixel.NewSprite(c.batches.grassBatch.sprite, c.batches.grassBatch.frames[rand.Intn(len(c.batches.grassBatch.frames))])
					c.batches.grassBatch.tiles = append(c.batches.grassBatch.tiles, t)

				case NorthSide:

					t.tile = pixel.NewSprite(c.batches.grassSideBatch.sprite, c.batches.grassSideBatch.frames[1+(rand.Intn(7)*4)])
					c.batches.grassSideBatch.tiles = append(c.batches.grassSideBatch.tiles, t)

				case EastSide:

					t.tile = pixel.NewSprite(c.batches.grassSideBatch.sprite, c.batches.grassSideBatch.frames[0+(rand.Intn(7)*4)])
					c.batches.grassSideBatch.tiles = append(c.batches.grassSideBatch.tiles, t)

				case SouthSide:

					t.tile = pixel.NewSprite(c.batches.grassSideBatch.sprite, c.batches.grassSideBatch.frames[3+(rand.Intn(7)*4)])
					c.batches.grassSideBatch.tiles = append(c.batches.grassSideBatch.tiles, t)

				case WestSide:

					t.tile = pixel.NewSprite(c.batches.grassSideBatch.sprite, c.batches.grassSideBatch.frames[2+(rand.Intn(7)*4)])
					c.batches.grassSideBatch.tiles = append(c.batches.grassSideBatch.tiles, t)

				case NorthEastInner:

					t.tile = pixel.NewSprite(c.batches.grassInnerBatch.sprite, c.batches.grassInnerBatch.frames[1+(rand.Intn(5)*4)])
					c.batches.grassInnerBatch.tiles = append(c.batches.grassInnerBatch.tiles, t)

				case SouthEastInner:

					t.tile = pixel.NewSprite(c.batches.grassInnerBatch.sprite, c.batches.grassInnerBatch.frames[0+(rand.Intn(5)*4)])
					c.batches.grassInnerBatch.tiles = append(c.batches.grassInnerBatch.tiles, t)

				case SouthWestInner:

					t.tile = pixel.NewSprite(c.batches.grassInnerBatch.sprite, c.batches.grassInnerBatch.frames[3+(rand.Intn(5)*4)])
					c.batches.grassInnerBatch.tiles = append(c.batches.grassInnerBatch.tiles, t)

				case NorthWestInner:

					t.tile = pixel.NewSprite(c.batches.grassInnerBatch.sprite, c.batches.grassInnerBatch.frames[2+(rand.Intn(5)*4)])
					c.batches.grassInnerBatch.tiles = append(c.batches.grassInnerBatch.tiles, t)

				case NorthEastOuter:

					t.tile = pixel.NewSprite(c.batches.grassOuterBatch.sprite, c.batches.grassOuterBatch.frames[1+(rand.Intn(5)*4)])
					c.batches.grassOuterBatch.tiles = append(c.batches.grassOuterBatch.tiles, t)

				case SouthEastOuter:

					t.tile = pixel.NewSprite(c.batches.grassOuterBatch.sprite, c.batches.grassOuterBatch.frames[0+(rand.Intn(5)*4)])
					c.batches.grassOuterBatch.tiles = append(c.batches.grassOuterBatch.tiles, t)

				case SouthWestOuter:

					t.tile = pixel.NewSprite(c.batches.grassOuterBatch.sprite, c.batches.grassOuterBatch.frames[3+(rand.Intn(5)*4)])
					c.batches.grassOuterBatch.tiles = append(c.batches.grassOuterBatch.tiles, t)

				case NorthWestOuter:

					t.tile = pixel.NewSprite(c.batches.grassOuterBatch.sprite, c.batches.grassOuterBatch.frames[2+(rand.Intn(5)*4)])
					c.batches.grassOuterBatch.tiles = append(c.batches.grassOuterBatch.tiles, t)

				}

			case Dirt:

				t.tile = pixel.NewSprite(c.batches.dirtBatch.sprite, c.batches.dirtBatch.frames[rand.Intn(len(c.batches.dirtBatch.frames))])
				c.batches.dirtBatch.tiles = append(c.batches.dirtBatch.tiles, t)

			}

		}

	}

}

func (c *chunk) Draw(can *pixelgl.Canvas) {
	can.Clear(color.RGBA{255, 255, 255, 255})

	// Draw Dirt Tiles
	drawBgTiles(c.batches.dirtBatch, c.batches.dirtSideBatch, c.batches.dirtInnerBatch, c.batches.dirtOuterBatch)
	// Draw Grass Tiles
	drawBgTiles(c.batches.grassBatch, c.batches.grassSideBatch, c.batches.grassInnerBatch, c.batches.grassOuterBatch)
	// Draw Water Tiles
	drawBgTiles(c.batches.waterBatch, c.batches.waterSideBatch, c.batches.waterInnerBatch, c.batches.waterOuterBatch)

	c.batches.dirtBatch.batch.Draw(can)
	c.batches.dirtSideBatch.batch.Draw(can)
	c.batches.dirtInnerBatch.batch.Draw(can)
	c.batches.dirtOuterBatch.batch.Draw(can)

	c.batches.grassBatch.batch.Draw(can)
	c.batches.grassSideBatch.batch.Draw(can)
	c.batches.grassInnerBatch.batch.Draw(can)
	c.batches.grassOuterBatch.batch.Draw(can)

	c.batches.waterBatch.batch.Draw(can)
	c.batches.waterSideBatch.batch.Draw(can)
	c.batches.waterInnerBatch.batch.Draw(can)
	c.batches.waterOuterBatch.batch.Draw(can)
}

func (c *chunk) loadBatches() {

	c.batches = &batchManager{}

	// Water Batches
	waterBatch := tileBatch{
		sprite: *c.sprites.waterSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.waterSprite),
	}
	waterBatch.loadFrames()
	c.batches.waterBatch = &waterBatch

	waterSideBatch := tileBatch{
		sprite: *c.sprites.waterSideSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.waterSideSprite),
	}
	waterSideBatch.loadFrames()
	c.batches.waterSideBatch = &waterSideBatch

	waterInnerBatch := tileBatch{
		sprite: *c.sprites.waterInnerSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.waterInnerSprite),
	}
	waterInnerBatch.loadFrames()
	c.batches.waterInnerBatch = &waterInnerBatch

	waterOuterBatch := tileBatch{
		sprite: *c.sprites.waterOuterSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.waterOuterSprite),
	}
	waterOuterBatch.loadFrames()
	c.batches.waterOuterBatch = &waterOuterBatch

	// Grass Batches
	grassBatch := tileBatch{
		sprite: *c.sprites.grassSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.grassSprite),
	}
	grassBatch.loadFullTileFrames()
	c.batches.grassBatch = &grassBatch

	grassSideBatch := tileBatch{
		sprite: *c.sprites.grassSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.grassSprite),
	}
	grassSideBatch.loadSideTileFrames()
	c.batches.grassSideBatch = &grassSideBatch

	grassInnerBatch := tileBatch{
		sprite: *c.sprites.grassSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.grassSprite),
	}
	grassInnerBatch.loadInnerTileFrames()
	c.batches.grassInnerBatch = &grassInnerBatch

	grassOuterBatch := tileBatch{
		sprite: *c.sprites.grassSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.grassSprite),
	}
	grassOuterBatch.loadOuterTileFrames()
	c.batches.grassOuterBatch = &grassOuterBatch

	// Dirt Batches
	dirtBatch := tileBatch{
		sprite: *c.sprites.dirtSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.dirtSprite),
	}
	dirtBatch.loadFullTileFrames()
	c.batches.dirtBatch = &dirtBatch

	dirtSideBatch := tileBatch{
		sprite: *c.sprites.dirtSideSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.dirtSprite),
	}
	dirtSideBatch.loadSideTileFrames()
	c.batches.dirtSideBatch = &dirtSideBatch

	dirtInnerBatch := tileBatch{
		sprite: *c.sprites.dirtInnerSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.dirtSprite),
	}
	dirtInnerBatch.loadInnerTileFrames()
	c.batches.dirtInnerBatch = &dirtInnerBatch

	dirtOuterBatch := tileBatch{
		sprite: *c.sprites.dirtOuterSprite,
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, *c.sprites.dirtSprite),
	}
	dirtOuterBatch.loadOuterTileFrames()
	c.batches.dirtOuterBatch = &dirtOuterBatch

}

func drawBgTiles(tBatches ...*tileBatch) {

	for _, tBatch := range tBatches {

		for _, t := range tBatch.tiles {

			t.tile.Draw(tBatch.batch, pixel.IM.Moved(t.pos))

		}

	}

}
