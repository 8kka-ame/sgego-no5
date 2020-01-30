package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480

	characterWidth  = 64
	characterHeight = 108

	interval = 10 // スプライトの変更間隔
)

var (
	frame     = 0 // 現在のフレーム数
	spriteNum = 0 // gopher画像の番号 (左から0, 1, 2, 3)
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

	// TODO: 10フレーム間隔で画像を変更する

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// TODO: x0, y0, x1, y1を指定してスプライト画像を矩形で区切る
	// x0 := ...
	// y0 := ...
	// x1 := ...
	// y1 := ...
	// gopher := image.Rect(x0, y0, x1, y1)

	// TODO: SubImageを使用して矩形領域の画像を表示する

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "02. Animation"); err != nil {
		log.Fatal(err)
	}
}
