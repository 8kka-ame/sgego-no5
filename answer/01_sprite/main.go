package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var character *ebiten.Image

func init() {
	var err error
	character, _, err = ebitenutil.NewImageFromFile("./img/gopher.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if err := screen.DrawImage(character, nil); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "01. Sprite"); err != nil {
		log.Fatal(err)
	}
}
