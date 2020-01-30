package main

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var character *ebiten.Image

func init() {
	// TODO: gopher.pngを読み込み、characterに格納する
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// TODO: 画像を描画する

	return nil
}

func main() {
	// TODO: ebiten.Runを実行する
}
