package Controllers

import (
	game "github.com/hajimehoshi/ebiten"
	"image"
)

var Img *game.Image
var PlayerSize image.Rectangle
var X, Y, R float64

type Controller interface {
	Start(screen *game.Image)
	moveUp()
	moveDown()
	moveLeft()
	moveRight()
}
