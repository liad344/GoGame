package main

import (
	utils "github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/jpeg"
	"log"

	game "github.com/hajimehoshi/ebiten"
	input "github.com/hajimehoshi/ebiten/inpututil"
)

const (
	width  = 1280
	height = 720
)

func Control(screen *game.Image) {
	go shoot(screen)
	moveByKeyInput(screen)

}
func shoot(screen *game.Image) {
	shot := input.IsKeyJustReleased(game.KeySpace)
	if shot {
		handleShot(screen)
	}
}
func handleShot(screen *game.Image) {
	bullet, _, err := utils.NewImageFromFile("images/bullet.jpeg", game.FilterDefault)
	if err != nil {
		log.Fatal("Cant load bullet", err)
	}
	op := &game.DrawImageOptions{}
	bx := x
	by := y
	op.GeoM.Translate(bx, by)
	screen.DrawImage(bullet, op)


}
func moveByKeyInput(screen *game.Image) {
	op := &game.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Rotate(r)
	screen.DrawImage(img, op)
	KeyLogic()
}
func KeyLogic() {
	w := input.KeyPressDuration(game.KeyW)
	s := input.KeyPressDuration(game.KeyS)
	a := input.KeyPressDuration(game.KeyA)
	d := input.KeyPressDuration(game.KeyD)
	rFast := input.KeyPressDuration(game.KeyR)
	rSlow := input.KeyPressDuration(game.KeyE)

	y -= float64(w)
	y += float64(s)
	if y > height {
		y = height - float64(PlayerSize.Max.Y)
	}
	if y < 0 {
		y = float64(PlayerSize.Max.Y)
	}
	x += float64(d)
	x -= float64(a)
	if x > width {
		x = width - float64(PlayerSize.Max.X)
	}
	if x < 0 {
		x = float64(PlayerSize.Max.X)
	}
	r += float64(rFast) * 0.001
	r -= float64(rSlow) * 0.001

}
