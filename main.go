package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

const statusBarHeight float32 = 40
const listItemHeight = 40

const screenWidth = 400
const screenHeight = 240

func drawGrid() {
	rl.PushMatrix()
	rl.Translatef(0, 25*50, 0)
	rl.Rotatef(90, 1, 0, 0)
	rl.DrawGrid(100, 50)
	rl.PopMatrix()
}

func easeOutQuart(x float64) float64 {
	return 1 - math.Pow(1 - x, 4)
}

func minInt(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {

	rl.InitWindow(screenWidth, screenHeight, "Willow Emulator")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	hp := homePage{}

	var selected element = &hp

	selected.create()

	for !rl.WindowShouldClose() {
		selected.tick()

		// DRAW
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		selected.draw()

		rl.EndDrawing()
	}

}
