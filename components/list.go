package components

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/jamesdearlove/willow/utils"
)

const moveDelay = 8
const animatePixels = 5

type List struct {
	X, Y float32

	Items         []ListItem
	SelectedIndex int

	TextStrings []string

	Width, Height, ItemHeight float32
	Active                    bool

	camera rl.Camera2D

	lastMoveCounter int
}

type ListItem struct {
	X, Y float32

	Width, Height float32
	Text          string
	Selected      bool
}

func (l *List) Create() {
	// TODO: Potentially not set these
	l.Active = true
	l.SelectedIndex = 0
	l.lastMoveCounter = 0

	// Create camera system for scrolling
	l.camera = rl.Camera2D{}
	l.camera.Zoom = 1.0

	l.BuildListItems()
}

// Create list items based on TextStrings
func (l *List) BuildListItems() {
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

func (l *List) Update() {

	// Check if moving up or down and delay has past
	if l.Active && l.lastMoveCounter > moveDelay {
		if rl.IsKeyDown(rl.KeyDown) {
			l.SelectedIndex = utils.MinInt(l.SelectedIndex+1, len(l.Items)-1)
			l.lastMoveCounter = 0
		} else if rl.IsKeyDown(rl.KeyUp) {
			l.SelectedIndex = utils.MaxInt(l.SelectedIndex-1, 0)
			l.lastMoveCounter = 0
		}
	} else {
		l.lastMoveCounter += 1
	}

	// TODO: Animations were removed because jank, to be readded with its own system.

	// Update the camera location
	// TODO: Remove the magic numbers
	screenHeight := 210

	middleOriginOffset := l.ItemHeight * (float32(screenHeight) / l.ItemHeight - 1) / 2
	totalListLength := l.ItemHeight * float32(len(l.Items))

	// Calculate the camera target location so the selected item is in the center
	centerLoc := float64(l.ItemHeight * float32(l.SelectedIndex) - middleOriginOffset)

	// Ensure the camera doesn't show the space before or after the list.
	targetBounding := float32(math.Min(math.Max(0, centerLoc), float64(totalListLength - float32(screenHeight))))
	l.camera.Target.Y = targetBounding

}

func (l *ListItem) Update() {}

func (l *List) Draw() {
	rl.BeginMode2D(l.camera)

	for i := 0; i < len(l.Items); i++ {
		l.Items[i].Selected = i == l.SelectedIndex

		l.Items[i].Draw()
	}

	rl.EndMode2D()
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
