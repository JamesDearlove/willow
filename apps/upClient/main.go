package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jamesdearlove/willow/apps/upClient/api"
	"github.com/joho/godotenv"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/jamesdearlove/willow/components"
)

const screenWidth = 400
const screenHeight = 240
const statusBarHeight float32 = 30
const listItemHeight = 40

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	upToken := os.Getenv("UP_TOKEN")
	accountId := os.Getenv("UP_ACCOUNT_ID")
	items, err := api.MakeTransactionListRequest(upToken, accountId)

	if err != nil {
		// idk handle this at some point
	}
	app(items)
}

func app(items api.Transaction) {
	rl.InitWindow(screenWidth, screenHeight, "Willow Emulator")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	// TODO: for now
	var itemStrings []string
	for i := 0; i < len(items.Data); i++ {
		transaction := items.Data[i]
		value := float32(transaction.Attributes.Amount.ValueAsBase) / 100

		var valueString string
		if value < 0 {
			valueString = fmt.Sprintf("$%.2f", value*-1)
		} else {
			valueString = fmt.Sprintf("+$%.2f", value)
		}

		newItem := fmt.Sprintf("%s - %s", transaction.Attributes.Description, valueString)
		itemStrings = append(itemStrings, newItem)
	}

	homePage := makeHomePage(itemStrings)

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

func makeHomePage(listText []string) components.Page {
	h := components.Page{}

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

		Text:       "Up Transactions",
		TextHeight: 20,
	}

	h.Items = []components.Component{&list, &infoBar}

	return h
}
