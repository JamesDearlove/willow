package components

import (
	"github.com/gen2brain/raylib-go/raylib"
	// "math"
)

type List struct {
	X, Y float32

	Items         []ListItem
	SelectedIndex int

	TextStrings []string

	Width, ItemHeight float32
}

type ListItem struct {
	X, Y float32

	Width, Height float32
	Text          string
	Selected      bool
}

func (l *List) Create() {
	l.SelectedIndex = 0

	for i := 0; i < len(l.TextStrings); i++ {
		newItem := ListItem{
			X:      l.X,
			Y:      l.Y + float32(i)*l.ItemHeight,
			Width:  l.Width,
			Height: l.ItemHeight,

			Text: l.TextStrings[i],
		}

		l.Items = append(l.Items, newItem)
	}
}

func (l *ListItem) Create() {}

func (l *List) Update()     {}
func (l *ListItem) Update() {}

func (l *List) Draw() {
	for i := 0; i < len(l.Items); i++ {
		l.Items[i].Selected = i == l.SelectedIndex

		l.Items[i].Draw()
	}
}

func (l *ListItem) Draw() {
	itemRect := rl.Rectangle{
		X:      l.X,
		Y:      l.Y,
		Width:  l.Width,
		Height: l.Height,
	}

	borderColour := rl.White
	if l.Selected {
		borderColour = rl.Black
	}

	rl.DrawRectangleRec(itemRect, rl.White)
	rl.DrawRectangleLinesEx(itemRect, 2, borderColour)
	rl.DrawText(l.Text, 10, int32(l.Y)+10, 20, rl.Black)
}
