package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	x, y float64
	color color.Color
	image *ebiten.Image
}

// const (
// 	tileSize   = 80
// 	tileMargin = 4
// )

// var (
// 	tileImage = ebiten.NewImage(tileSize, tileSize)
// )

// func init() {
// 	tileImage.Fill(color.White)
// }

// type Game struct{}

// func (g *Game) Update() error {
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	ebitenutil.DebugPrint(screen, "Hello, World!")
	
// 	// draw box and move it to x:100, y:20
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(100.0, 20.0)
// 	screen.DrawImage(tileImage, op)
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 320, 240
// }