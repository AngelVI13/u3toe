package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


const (
	ScreenWidth = 400
	ScreenHeight = 400
	// -----------
	tileSize   = 80
	tileMargin = 4
	// -----------
	rowSize = 3
	boardSize = 9
)

var (
	tileImage = ebiten.NewImage(tileSize, tileSize)
	blueColor = color.RGBA{0, 0, 255, 255}
	boardOffsetX = (ScreenWidth - rowSize*(tileSize + tileMargin) - tileMargin) / 2
	boardOffsetY = (ScreenHeight - rowSize*(tileSize + tileMargin) - tileMargin) / 2
)

func init() {
	tileImage.Fill(color.White)
}

type Game struct{
	tiles []*Tile
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{}
	// var err error

	for i := 0; i < boardSize; i++{
		row := i / rowSize
		col := i % rowSize

		tile := Tile{
			x: float64(boardOffsetX + row*(tileSize + tileMargin)),
			y: float64(boardOffsetY + col*(tileSize + tileMargin)),
			color: blueColor,
			image: tileImage,
			marked: NoMark,
		}
		fmt.Println(tile)
		g.tiles = append(g.tiles, &tile)
	} 
	
	return g, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	// draw box and move it to x:100, y:20
	for _, tile := range g.tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(tile.x, tile.y)
		tile.image.Fill(tile.color)
		screen.DrawImage(tile.image, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}