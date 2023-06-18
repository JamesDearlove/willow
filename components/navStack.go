package components

type NavStack struct {
	
	DefaultPage Page
	
	items []Page
}

func (n *NavStack) Push(p *Page) {

	n.items = append(n.items, *p)
}

func (n *NavStack) Pop() {

}

func (n *NavStack) CurrentPage() *Page {
	return &n.items[len(n.items)-1]
}

func (n *NavStack) Create() {
	

}

func (n *NavStack) Draw() {
	n.CurrentPage().Draw()
}

func (n *NavStack) Update() {
	n.CurrentPage().Update()
}

