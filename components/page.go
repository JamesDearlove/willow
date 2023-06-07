package components

import (
// "github.com/gen2brain/raylib-go/raylib"
)

type Page struct {
	// heading string
	Items   []Component
}

func (p *Page) Create() {
	for i := 0; i < len(p.Items); i++ {
		p.Items[i].Create()
	}
}

func (p *Page) Draw() {
	for i := 0; i < len(p.Items); i++ {
		p.Items[i].Draw()
	}
}

func (p *Page) Update() {
	for i := 0; i < len(p.Items); i++ {
		p.Items[i].Update()
	}
}
