package components

type Component interface {
	Create()
	Draw()
	Update()
}


func createComponent(c Component) {
	if c == nil {
		return
	}
	c.Create()
}

func drawComponent(c Component) {
	if c == nil {
		return
	}
	c.Draw()
}

func updateComponent(c Component) {
	if c == nil {
		return
	}
	c.Update()
}
