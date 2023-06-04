package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	// "math"
)

type control interface {
	draw()
}

type list struct {
	xPos, yPos    float32
	items         []listItem
	selectedIndex int

	width, itemHeight float32
}

type listItem struct {
	xPos, yPos, width, height float32
	text                      string
	selected                  bool
}

func (l *list) createItems(strings []string) {
	l.selectedIndex = 0

	for i := 0; i < len(strings); i++ {
		newItem := listItem{
			xPos:   l.xPos,
			yPos:   float32(i) * l.itemHeight,
			width:  l.width,
			height: l.itemHeight,

			text: strings[i],
		}

		l.items = append(l.items, newItem)
	}
}

func (l list) draw() {
	for i := 0; i < len(l.items); i++ {
		l.items[i].selected = i == l.selectedIndex

		l.items[i].draw()
	}
}

func (l listItem) draw() {
	itemRect := rl.Rectangle{
		X:      l.xPos,
		Y:      l.yPos,
		Width:  l.width,
		Height: l.height,
	}

	borderColour := rl.White
	if l.selected {
		borderColour = rl.Black
	}

	rl.DrawRectangleRec(itemRect, rl.White)
	rl.DrawRectangleLinesEx(itemRect, 2, borderColour)
	rl.DrawText(l.text, 10, int32(l.yPos)+10, 20, rl.Black)
}
