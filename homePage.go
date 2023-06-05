package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jamesdearlove/willow/components"
)

type homePage struct {
	camera rl.Camera2D

	cameraMove      float32
	moveDelay       int
	lastMoveCounter int

	listC    components.List
	listText []string
}

func (h *homePage) Create() {
	h.camera = rl.Camera2D{}
	h.camera.Zoom = 1.0

	h.cameraMove = float32(0)

	h.moveDelay = 8
	h.lastMoveCounter = h.moveDelay

	h.listText = []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9", "Item 10"}

	h.listC = components.List{
		X:             0,
		Y:             0,
		Width:         float32(screenWidth),
		ItemHeight:    listItemHeight,
		SelectedIndex: 0,
		TextStrings:   h.listText,
	}
	
	h.listC.Create()
}

func (h *homePage) Update() {
	if h.lastMoveCounter == h.moveDelay {
		if rl.IsKeyDown(rl.KeyDown) {
			h.listC.SelectedIndex = minInt(h.listC.SelectedIndex+1, len(h.listText)-1)
			h.lastMoveCounter = 0
		} else if rl.IsKeyDown(rl.KeyUp) {
			h.listC.SelectedIndex = maxInt(h.listC.SelectedIndex-1, 0)
			h.lastMoveCounter = 0
		}
	} else {
		h.lastMoveCounter += 1
	}

	// Check camera can check bounds of list
	viewportYHeight := float32(screenHeight - int32(statusBarHeight))
	currentLoc := float32(h.listC.SelectedIndex * listItemHeight)

	if currentLoc >= h.camera.Target.Y+viewportYHeight {
		h.cameraMove = 1
	} else if currentLoc < h.camera.Target.Y {
		h.cameraMove = -1
	}

	// TODO: This animation is sketch
	if h.lastMoveCounter < 8 {
		x := float64(h.lastMoveCounter) / 4
		h.camera.Target.Y += float32(easeOutQuart(x)) * h.cameraMove * 6.45
	} else if h.lastMoveCounter == 8 {
		h.camera.Target.Y = float32(math.Round(float64(h.camera.Target.Y)))
		h.cameraMove = 0
	}
}

func (h *homePage) Draw() {
	rl.BeginMode2D(h.camera)

	// Draw the 3d grid, rotated 90 degrees and centered around 0,0 for the XY Plane
	// if grid {
	// 	drawGrid()
	// }

	// drawList(listItems, SelectedIndex)
	h.listC.Draw()

	rl.EndMode2D()

	topBar := rl.Rectangle{
		X:      0,
		Y:      float32(screenHeight) - statusBarHeight,
		Width:  float32(screenWidth),
		Height: statusBarHeight,
	}

	rl.DrawRectanglePro(topBar, rl.Vector2{X: 0, Y: 0}, 0, rl.Black)

	// rl.DrawRectangleLinesEx(topBar, 2, rl.DarkGray)
	rl.DrawText("A List", 10, int32(topBar.Y)+10, 20, rl.White)
}
