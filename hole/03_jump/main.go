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

	tileWidth  = 32
	tileHeight = 32

	interval = 10
)

var (
	frame     = 0
	spriteNum = 0

	character *ebiten.Image
	tile      *ebiten.Image
)

func init() {
	var err error
	character, _, err = ebitenutil.NewImageFromFile("./img/gopher.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: タイル画像を読み込む
}

func update(screen *ebiten.Image) error {
	frame++

	if frame%interval == 0 {
		spriteNum++
	}
	if spriteNum > 2 {
		spriteNum = 0
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// TODO: ループでタイルを描画する
	// 移動しているように見せるには、「1~10フレームはそのまま表示し、11~20フレームはタイル画像を半分ずらす」を繰り返す

	characterOp := &ebiten.DrawImageOptions{}
	// TODO: キャラクターをタイルの上に移動させる

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
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "03. Running"); err != nil {
		log.Fatal(err)
	}
}
