package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

const statusBarHeight float32 = 40
const listItemHeight = 40

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
	screenWidth := int32(400)
	screenHeight := int32(240)

	rl.InitWindow(screenWidth, screenHeight, "400x240 Screen")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	camera := rl.Camera2D{}
	camera.Zoom = 1.0
	const grid = false

	cameraMove := float32(0)

	const moveDelay = 8
	lastMoveCounter := moveDelay

	listText := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9", "Item 10"}

	listC := list{
		xPos: 0,
		yPos: 0,
		width: float32(screenWidth),
		itemHeight: listItemHeight,
		selectedIndex: 0,
	}

	listC.createItems(listText)

	for !rl.WindowShouldClose() {
		
		if lastMoveCounter == moveDelay {
			if rl.IsKeyDown(rl.KeyDown) {
				listC.selectedIndex = minInt(listC.selectedIndex+1, len(listText)-1)
				lastMoveCounter = 0
			} else if rl.IsKeyDown(rl.KeyUp) {
				listC.selectedIndex = maxInt(listC.selectedIndex-1, 0)
				lastMoveCounter = 0
			}
		} else {
			lastMoveCounter += 1
		}

		// Check camera can check bounds of list
		viewportYHeight := float32(screenHeight - int32(statusBarHeight))
		currentLoc := float32(listC.selectedIndex * listItemHeight)

		if currentLoc >= camera.Target.Y+viewportYHeight {
			cameraMove = 1
		} else if currentLoc < camera.Target.Y {
			cameraMove = -1
		}

		// TODO: This animation is sketch
		if lastMoveCounter < 8 {
			x := float64(lastMoveCounter) / 4
			camera.Target.Y += float32(easeOutQuart(x)) * cameraMove * 6.45
		} else if lastMoveCounter == 8 {
			camera.Target.Y = float32(math.Round(float64(camera.Target.Y)))
			cameraMove = 0
		}

		// DRAW
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		rl.BeginMode2D(camera)

		// Draw the 3d grid, rotated 90 degrees and centered around 0,0 for the XY Plane
		if grid {
			drawGrid()
		}

		// drawList(listItems, selectedIndex)
		listC.draw()

		rl.EndMode2D()

		topBar := rl.Rectangle{
			X:      0,
			Y:      float32(screenHeight) - statusBarHeight,
			Width:  float32(screenWidth),
			Height: statusBarHeight,
		}

		rl.DrawRectanglePro(topBar, rl.Vector2{X: 0, Y: 0}, 0, rl.RayWhite)
		rl.DrawRectangleLinesEx(topBar, 2, rl.DarkGray)
		rl.DrawText("Some App", 10, int32(topBar.Y) + 10, 20, rl.Black)

		rl.EndDrawing()
	}

}
