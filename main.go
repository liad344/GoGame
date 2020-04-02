package main

import (
	"image"
	_ "image/png"
	log "github.com/sirupsen/logrus"

	game "github.com/hajimehoshi/ebiten"
	utils "github.com/hajimehoshi/ebiten/ebitenutil"
)

var img *game.Image
var PlayerSize image.Rectangle

func init() {
	var err error
	img, _, err = utils.NewImageFromFile("images/mario.png", game.FilterDefault)
	if err != nil{
		log.Fatal(err)
	}
	PlayerSize = img.Bounds()
	if err != nil {
		log.Fatal(err)
	}
}

var x, y , r float64
func update(screen *game.Image) error {
	if game.IsDrawingSkipped() {
		return nil
	}
	Control(screen)

	return nil
}



func main() {
	if err := game.Run(update, 1280, 720, 1, "Geometry Matrix"); err != nil {
		log.Fatal(err)
	}
}