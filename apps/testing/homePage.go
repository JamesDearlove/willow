package testing

import (
	// "github.com/gen2brain/raylib-go/raylib"
	"github.com/jamesdearlove/willow/components"
)

const statusBarHeight float32 = 30
const listItemHeight = 40


func MakeHomePage() components.Page {
	h := components.Page{}

	listText := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9", "Item 10", "Item 11", "Item 12", "Item 13", "Item 14", "Item 15"}

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
