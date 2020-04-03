package Controllers

import (
	game "github.com/hajimehoshi/ebiten"
	input "github.com/hajimehoshi/ebiten/inpututil"
	_ "image/jpeg"
	"log"
)

const (
	width  = 600
	height = 720
)

type KeyController struct {
}

func (k KeyController) Start(screen *game.Image) {
	go k.moveUp()
	go k.moveDown()
	go k.moveLeft()
	go k.moveRight()
	avoidEdge()
	move(screen)

}

func move(screen *game.Image) {
	op := &game.DrawImageOptions{}
	op.GeoM.Translate(X, Y)
	err := screen.DrawImage(Img, op)
	if err != nil {
		log.Fatal(err)
	}
}
func avoidEdge() {
	if Y > height {
		Y = height - float64(PlayerSize.Max.Y)
	}
	if Y < 0 {
		Y = float64(PlayerSize.Max.Y)
	}
	if X > width {
		X = width - float64(PlayerSize.Max.X)
	}
	if X < 0 {
		X = float64(PlayerSize.Max.X)
	}
}

func (k KeyController) moveUp() {
	w := input.KeyPressDuration(game.KeyW)
	Y -= float64(w)
}
func (k KeyController) moveDown() {
	s := input.KeyPressDuration(game.KeyS)
	Y += float64(s)
}
func (k KeyController) moveLeft() {
	a := input.KeyPressDuration(game.KeyA)
	X -= float64(a)
}
func (k KeyController) moveRight() {
	d := input.KeyPressDuration(game.KeyD)
	X += float64(d)
}
