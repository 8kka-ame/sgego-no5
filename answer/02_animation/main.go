package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480

	characterWidth  = 64
	characterHeight = 108

	interval = 10
)

var (
	frame     = 0
	spriteNum = 0
	character *ebiten.Image
)

func init() {
	var err error
	character, _, err = ebitenutil.NewImageFromFile("./img/gopher.png", ebiten.FilterDefault)
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

	x0 := characterWidth * spriteNum
	y0 := 0
	x1 := x0 + characterWidth
	y1 := characterHeight
	gopher := image.Rect(x0, y0, x1, y1)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if err := screen.DrawImage(character.SubImage(gopher).(*ebiten.Image), nil); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "02. Animation"); err != nil {
		log.Fatal(err)
	}
}
