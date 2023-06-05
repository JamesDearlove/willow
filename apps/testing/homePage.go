package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/jamesdearlove/willow/components"
)

type homePage struct {
	listC    components.List
	listText []string
}

func (h *homePage) Create() {

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
	h.listC.Update()

}

func (h *homePage) Draw() {

	h.listC.Draw()

	topBar := rl.Rectangle{
		X:      0,
		Y:      float32(screenHeight) - statusBarHeight,
		Width:  float32(screenWidth),
		Height: statusBarHeight,
	}

	rl.DrawRectanglePro(topBar, rl.Vector2{X: 0, Y: 0}, 0, rl.Black)
	rl.DrawText("A List", 10, int32(topBar.Y)+10, 20, rl.White)
}
