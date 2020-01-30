package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480

	characterWidth  = 64
	characterHeight = 108

	tileWidth  = 32
	tileHeight = 32
	tileSpeed  = 10

	interval = 10
)

var (
	frame     = 0
	spriteNum = 0

	jump           = false
	jumpHeight     = 0
	jumpFrame      = 0
	jumpFrameLimit = 40
	acceleration   = 4

	character *ebiten.Image
	tile      *ebiten.Image
)

func init() {
	var err error
	character, _, err = ebitenutil.NewImageFromFile("./img/gopher.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	tile, _, err = ebitenutil.NewImageFromFile("./img/tile.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	frame++

	if frame%interval == 0 {
		spriteNum++
	}
	if spriteNum > 2 {
		spriteNum = 0
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		jump = true
	}

	if jump {
		spriteNum = 3
		jumpFrame++
		if jumpFrame < jumpFrameLimit {
			jumpHeight += acceleration
		} else {
			jumpHeight -= acceleration
		}
		if jumpHeight <= 0 {
			jump = false
			jumpFrame = 0
			jumpHeight = 0
			spriteNum = 0
		}
	}

	speedWidth := 0
	if frame%(tileSpeed*2) > tileSpeed {
		speedWidth = tileWidth / 2
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for i := 0; i < screenWidth+tileWidth; i += tileWidth {
		tileOp := &ebiten.DrawImageOptions{}
		tileOp.GeoM.Translate(float64(i-speedWidth), float64(screenHeight-tileHeight))
		if err := screen.DrawImage(tile, tileOp); err != nil {
			return err
		}
	}

	characterOp := &ebiten.DrawImageOptions{}
	characterOp.GeoM.Translate(-float64(characterWidth/2), -float64(characterHeight))
	characterOp.GeoM.Translate(screenWidth/2, screenHeight-tileHeight)
	characterOp.GeoM.Translate(0, -float64(jumpHeight))

	x0 := characterWidth * spriteNum
	y0 := 0
	x1 := x0 + characterWidth
	y1 := characterHeight
	gopher := image.Rect(x0, y0, x1, y1)

	if err := screen.DrawImage(character.SubImage(gopher).(*ebiten.Image), characterOp); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "04. Jump"); err != nil {
		log.Fatal(err)
	}
}
