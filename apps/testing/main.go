package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/jamesdearlove/willow/components"
)

const screenWidth = 400
const screenHeight = 240

func main() {

	rl.InitWindow(screenWidth, screenHeight, "Willow Emulator")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	homePage := makeHomePage()

	var selected components.Component = &homePage

	selected.Create()

	for !rl.WindowShouldClose() {
		selected.Update()

		// DRAW
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		selected.Draw()

		rl.EndDrawing()
	}

}
