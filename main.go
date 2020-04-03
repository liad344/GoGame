package main

import (
	log "github.com/sirupsen/logrus"
	_ "image/png"
	"liad/game/Controllers"

	game "github.com/hajimehoshi/ebiten"
	utils "github.com/hajimehoshi/ebiten/ebitenutil"
)

var Player Character

func init() {
	var err error
	Controllers.Img, _, err = utils.NewImageFromFile("images/mario.png", game.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	Controllers.PlayerSize = Controllers.Img.Bounds()
	if err != nil {
		log.Fatal(err)
	}
	Player = Character{
		Items:      item{},
		Controller: Controllers.KeyController{},
	}
}

func update(screen *game.Image) error {
	if game.IsDrawingSkipped() {
		return nil
	}
	Player.Controller.Start(screen)
	return nil
}

func main() {
	if err := game.Run(update, 1280, 720, 1, "Geometry Matrix"); err != nil {
		log.Fatal(err)
	}
}
