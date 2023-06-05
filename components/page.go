package components

import (
// "github.com/gen2brain/raylib-go/raylib"
)

type Page struct {
	// heading string
	items   []Component
}

func (p *Page) Create() {
	for i := 0; i < len(p.items); i++ {
		p.items[i].Create()
	}
}

func (p *Page) Draw() {
	for i := 0; i < len(p.items); i++ {
		p.items[i].Draw()
	}
}

func (p *Page) Update() {
	for i := 0; i < len(p.items); i++ {
		p.items[i].Update()
	}
}
