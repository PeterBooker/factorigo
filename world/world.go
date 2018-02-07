package world

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/sixthgear/noise"
)

const (
	tileSize  = 32
	chunkSize = 1024 // 32 x 32
)

// World ...
type World struct {
	seed          int
	SurfaceType   [][]SurfaceTileType
	SurfaceShape  [][]SurfaceTileShape
	ResourceType  [][]uint8
	ResourceValue [][]uint8
	EntityType    [][]uint8
	EntityShape   [][]uint8
	Chunks        [][]*chunk
	LiveChunks    []*chunk
}

type neighbours struct {
	n, ne, e, se, s, sw, w, nw float64
	t                          int
}

const (
	wL = -0.35
	gL = 0.45
)

// SurfaceTileType describes the type of tile present, e.g. water, grass, sand, etc.
type SurfaceTileType uint8 // 0-255

const (
	Water SurfaceTileType = iota
	Grass
	Dirt
	Sand
)

// SurfaceTileShape describes the shape of tile present, e.g. full, north side, south east inner, etc.
type SurfaceTileShape uint8 // 0-255

const (
	Full SurfaceTileShape = iota
	NorthSide
	EastSide
	SouthSide
	WestSide
	NorthEastInner
	SouthEastInner
	SouthWestInner
	NorthWestInner
	NorthEastOuter
	SouthEastOuter
	SouthWestOuter
	NorthWestOuter
	NorthDoubleInner
	EastDoubleInner
	SouthDoubleInner
	WestDoubleInner
)

// Chunk splits the World into Chunks
func (w *World) Chunk() error {

	var x int
	var y int

	chunk := newChunk(pixel.R(0, 0, 100, 100))

	chunk.SurfaceType = make([][]SurfaceTileType, 100)
	for i := range chunk.SurfaceType {
		chunk.SurfaceType[i] = make([]SurfaceTileType, 100)
	}

	chunk.SurfaceShape = make([][]SurfaceTileShape, 100)
	for i := range chunk.SurfaceShape {
		chunk.SurfaceShape[i] = make([]SurfaceTileShape, 100)
	}

	for x = 0; x < 100; x++ {

		for y = 0; y < 100; y++ {

			chunk.SurfaceType[x][y] = w.SurfaceType[x][y]
			chunk.SurfaceShape[x][y] = w.SurfaceShape[x][y]

		}

	}

	//w.Chunks[0][0] = chunk

	chunk.Build()

	w.LiveChunks = append(w.LiveChunks, chunk)

	return nil

}

// Save saves the current world to file.
func (w *World) Save() error {

	return nil

}

// Load loads the current World from file.
func Load() *World {

	return &World{}

}

// New returns a new World based on the width and height given.
func New(width, height int) *World {

	octaves := 4
	p := 0.25
	scale := 1.0 / 42

	grid := generateProceduralGrid(width, height, octaves, p, scale)

	surTypes, surShapes := parseSurfaceTiles(grid)

	// Init Empty Grid
	emptyGrid := make([][]uint8, width)
	for i := range emptyGrid {
		emptyGrid[i] = make([]uint8, height)
	}

	// Fill the Level with empty defaults
	return &World{
		SurfaceType:   surTypes,
		SurfaceShape:  surShapes,
		ResourceType:  emptyGrid,
		ResourceValue: emptyGrid,
		EntityType:    emptyGrid,
		EntityShape:   emptyGrid,
	}

}

func generateProceduralGrid(width, height int, octaves int, p, scale float64) [][]float64 {

	grid := make([][]float64, width)
	for i := range grid {
		grid[i] = make([]float64, height)
	}

	for x := 0.; x < float64(width); x++ {

		for y := 0.; y < float64(height); y++ {

			ns := noise.OctaveNoise2d(x, y, octaves, p, scale)
			grid[int(x)][int(y)] = ns

		}

	}

	return grid

}

func parseSurfaceTiles(oldGrid [][]float64) ([][]SurfaceTileType, [][]SurfaceTileShape) {

	typeGrid := make([][]SurfaceTileType, len(oldGrid))
	for i := range typeGrid {
		typeGrid[i] = make([]SurfaceTileType, len(oldGrid[0]))
	}

	shapeGrid := make([][]SurfaceTileShape, len(oldGrid))
	for i := range shapeGrid {
		shapeGrid[i] = make([]SurfaceTileShape, len(oldGrid[0]))
	}

	for x := 0; x < len(oldGrid); x++ {

		for y := 0; y < len(oldGrid[0]); y++ {

			t := oldGrid[x][y]

			switch {

			case t < wL:

				nbs := getNeighbours(oldGrid, x, y, wL)
				typeGrid[x][y] = Water
				shapeGrid[x][y] = parseSurfaceShapes(nbs)

			case t >= wL && t < gL:

				nbs := getNeighbours(oldGrid, x, y, gL)
				typeGrid[x][y] = Grass
				shapeGrid[x][y] = parseSurfaceShapes(nbs)

			case t >= gL:
				// Not using sand Shapes yet
				//nbs := getNeighbours(oldGrid, x, y, wL)
				typeGrid[x][y] = Dirt
				shapeGrid[x][y] = Full

			}

		}

	}

	return typeGrid, shapeGrid

}

func parseSurfaceShapes(nbs uint8) SurfaceTileShape {

	switch nbrs := fmt.Sprintf("%08b", nbs); nbrs {

	// Inner Textures

	// North East Inner
	case "01000000":
		return NorthEastInner

	// South East Inner
	case "00010000":
		return SouthEastInner

	// South West Inner
	case "00000100":
		return SouthWestInner

	// North West Inner
	case "00000001":
		return NorthWestInner

	// Side Textures

	// North Side
	case "10000000", "10000001", "11000000", "11000001":
		return NorthSide

	// East Side
	case "00100000", "01100000", "00110000", "01110000":
		return EastSide

	// South Side
	case "00001000", "00011000", "00001100", "00011100":
		return SouthSide

	// West Side
	case "00000010", "00000110", "00000011", "00000111":
		return WestSide

	// Outer Textures

	// North East Outer
	case "11100000", "11100001", "11110000", "11110001":
		return NorthEastOuter

	// South East Outer
	case "00111000", "01111000", "00111100", "01111100":
		return SouthEastOuter

	// South West Outer
	case "00001110", "00011110", "00001111", "00011111":
		return SouthWestOuter

	// North West Outer
	case "10000011", "10000111", "11000011", "11000111":
		return NorthWestOuter

	// Double Inner

	// North Double Inner
	case "11100011":
		return NorthDoubleInner

	// East Double Inner
	case "11111000":
		return EastDoubleInner

	// South Double Inner
	case "00111110":
		return SouthDoubleInner

	// West Double Inner
	case "10001111":
		return WestDoubleInner

	// All others are full
	default:
		return Full

	}

}

func getNeighbours(grid [][]float64, x, y int, threshold float64) uint8 {

	var nbs uint8

	size := len(grid)
	rsize := size - 1

	// North
	if y < rsize && grid[x][y+1] >= threshold {
		nbs = setNeighbour(nbs, 7)
	}

	// North East
	if x < rsize && y < rsize && grid[x+1][y+1] >= threshold {
		nbs = setNeighbour(nbs, 6)
	}

	// East
	if x < rsize && grid[x+1][y] >= threshold {
		nbs = setNeighbour(nbs, 5)
	}

	// South East
	if x < rsize && y > 0 && grid[x+1][y-1] >= threshold {
		nbs = setNeighbour(nbs, 4)
	}

	// South
	if y > 0 && grid[x][y-1] >= threshold {
		nbs = setNeighbour(nbs, 3)
	}

	// South West
	if x > 0 && y > 0 && grid[x-1][y-1] >= threshold {
		nbs = setNeighbour(nbs, 2)
	}

	// West
	if x > 0 && grid[x-1][y] >= threshold {
		nbs = setNeighbour(nbs, 1)
	}

	// North West
	if y < rsize && x > 0 && grid[x-1][y+1] >= threshold {
		nbs = setNeighbour(nbs, 0)
	}

	return nbs

}

// Set relevant bit to show neighbour is different
func setNeighbour(n uint8, pos uint) uint8 {
	n |= (1 << pos)
	return n
}

// NewLevel generates a new Level based on the width and height given.
// The surface grid uses Perlin Noise
/*func NewLevelOld(width, height int) [][]uint8 {

	seed := time.Now().UnixNano()

	level, err := genLevel(seed, width, height)
	if err != nil {
		panic(err)
	}

	return level

}

func genLevel(seed int64, width, height int) ([][]uint8, error) {

	surfGrid, err := genSurface(seed, width, height)
	if err != nil {
		panic(err)
	}

	grid := surfGrid

	return grid, nil

}

func genSurface(seed int64, width, height int) ([][]uint8, error) {

	grid := make([][]uint8, width)
	for i := range grid {
		grid[i] = make([]uint8, height)
	}

	r := rand.New(rand.NewSource(seed))

	chanceStartAlive := 0.45

	for x := 0; x < width; x++ {

		for y := 0; y < height; y++ {

			if r.Float64() < chanceStartAlive {

				grid[x][y] = 1

			}

		}

	}

	numSimSteps := 4

	for i := 1; i < numSimSteps; i++ {
		grid = doSimulationStep(grid)
	}

	return grid, nil

}

func doSimulationStep(oldGrid [][]uint8) [][]uint8 {

	birthLimit := 3
	deathLimit := 2

	newGrid := make([][]uint8, len(oldGrid))
	for i := range newGrid {
		newGrid[i] = make([]uint8, len(oldGrid[0]))
	}

	for x := 0; x < len(oldGrid); x++ {

		for y := 0; y < len(oldGrid[0]); y++ {

			nbs := countAliveNeighbours(oldGrid, x, y)

			//The new value is based on our simulation rules
			//First, if a cell is alive but has too few neighbours, kill it.
			if oldGrid[x][y] == 1 {

				if nbs < deathLimit {
					newGrid[x][y] = 0
				} else {
					newGrid[x][y] = 1
				}

			} else {

				if nbs > birthLimit {
					newGrid[x][y] = 1
				} else {
					newGrid[x][y] = 0
				}

			}

		}

	}

	return newGrid

}

func countAliveNeighbours(grid [][]uint8, x, y int) int {

	count := 0

	for rx := -1; rx < 2; rx++ {

		for ry := -1; ry < 2; ry++ {

			// Neighbour Cords
			nx := x + rx
			ny := y + ry

			// Ignore the main tile itself
			if rx == 0 && ry == 0 {
				continue
			}

			if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) {
				count = count + 1
				continue
			}

			//Otherwise, a normal check of the neighbour
			if grid[nx][ny] == 1 {
				count = count + 1
				continue
			}
		}

	}

	return count

}*/
