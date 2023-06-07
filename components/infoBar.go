package components

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const textXPadding = 10

type InfoBar struct {
	X, Y, Width, Height float32

	Text       string
	TextHeight int32

	rect rl.Rectangle
}

func (i *InfoBar) Create() {

	i.rect = rl.Rectangle{
		X:      i.X,
		Y:      i.Y,
		Width:  i.Width,
		Height: i.Height,
	}
}

func (i *InfoBar) Update() {

}

func (i *InfoBar) Draw() {
	textYPadding := (i.Height - float32(i.TextHeight)) / 2

	rl.DrawRectangleRec(i.rect, rl.Black)
	rl.DrawText(
		i.Text, 
		textXPadding, 
		int32(i.Y) + int32(textYPadding), 
		20, 
		rl.White,
	)
}
