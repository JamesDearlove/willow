package system

import (
	// "github.com/gen2brain/raylib-go/raylib"
	"github.com/jamesdearlove/willow/components"
)

const statusBarHeight float32 = 30
const listItemHeight = 40

const screenWidth = 400
const screenHeight = 240


func MakeHomeScreen() components.Page {
	h := components.Page{}

	listText := []string{"Beeper", "Weather", "Settings"}

	list := components.List{
		X:             0,
		Y:             0,
		Width:         float32(screenWidth),
		Height:        float32(screenHeight - statusBarHeight),
		ItemHeight:    listItemHeight,
		SelectedIndex: 0,
		TextStrings:   listText,
	}

	infoBar := components.InfoBar{
		X:      0,
		Y:      float32(screenHeight) - statusBarHeight,
		Width:  float32(screenWidth),
		Height: statusBarHeight,

		Text:       "Some Text",
		TextHeight: 20,
	}

	h.Items = []components.Component{&list, &infoBar}

	return h
}
