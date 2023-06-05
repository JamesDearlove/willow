package components

import "github.com/gen2brain/raylib-go/raylib"

func drawGrid() {
	rl.PushMatrix()
	rl.Translatef(0, 25*50, 0)
	rl.Rotatef(90, 1, 0, 0)
	rl.DrawGrid(100, 50)
	rl.PopMatrix()
}
