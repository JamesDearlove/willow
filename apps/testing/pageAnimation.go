package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/jamesdearlove/willow/components"
)

func makeAnimatePAge() components.Page {
	h := components.Page{}

	page := WritablePage{
		X:      20,
		Y:      20,
		Width:  screenHeight,
		Height: screenHeight,
	}

	infoBar := components.InfoBar{
		X:      0,
		Y:      float32(screenHeight) - statusBarHeight,
		Width:  float32(screenWidth),
		Height: statusBarHeight,

		Text:       "Compose",
		TextHeight: 20,
	}

	h.Items = []components.Component{&page, &infoBar}

	return h
}

type WritablePage struct {
	X, Y, Width, Height float32

	edgePadding float32

	pageRect         rl.Rectangle
	animationCounter int32

	text string
}

func (p *WritablePage) Create() {
	p.animationCounter = 0
	p.text = ""
	p.edgePadding = 5

	p.pageRect = rl.Rectangle{
		X:      20,
		Y:      screenHeight,
		Width:  360,
		Height: 220,
	}
}

func (p *WritablePage) Update() {
	// TODO: Need to convert all input keys to text
	if rl.IsKeyPressed(rl.KeyQ) {
		p.text += "Q"
	} else if rl.IsKeyPressed(rl.KeyW) {
		p.text += "W"
	} else if rl.IsKeyPressed(rl.KeyE) {
		p.text += "E"
	} else if rl.IsKeyPressed(rl.KeyR) {
		p.text += "R"
	} else if rl.IsKeyPressed(rl.KeyT) {
		p.text += "T"
	} else if rl.IsKeyPressed(rl.KeyY) {
		p.text += "Y"
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		p.text += "\n"
	} else if rl.IsKeyPressed(rl.KeyBackspace) {
		if len(p.text) > 0 {
			p.text = p.text[0 : len(p.text)-1]
		}
	}

	if p.pageRect.Y != p.Y {
		if p.animationCounter > 10 {

			diff := p.pageRect.Y - p.Y
			p.pageRect.Y = p.pageRect.Y - diff/10
		}
	}

	p.animationCounter++
}

func (p *WritablePage) Draw() {
	rl.DrawRectangleRounded(p.pageRect, 1, 0, rl.White)
	rl.DrawRectangleRoundedLines(p.pageRect, 0.05, 8, 2, rl.Black)

	// TODO: Yeah this needs to wrap
	rl.DrawText(p.text, int32(p.pageRect.X+p.edgePadding), int32(p.pageRect.Y+p.edgePadding), 16, rl.Black)
}
