package rover

type direction interface {
	forward(p *position) *position
	backward(p *position) *position
	left(p *position) *position
	right(p *position) *position
}

type Directions struct {
	North direction
	East  direction
	West  direction
	South direction
}

type North struct{}

type East struct{}

type West struct{}

type South struct{}

func (n North) forward(p *position) *position {
	return p
}
func (n North) backward(p *position) *position {
	return p
}
func (n North) left(p *position) *position {
	return p
}
func (n North) right(p *position) *position {
	return p
}

func (e East) forward(p *position) *position {
	return p
}

func (e East) backward(p *position) *position {
	return p
}
func (e East) left(p *position) *position {
	return p
}
func (e East) right(p *position) *position {
	return p
}

func (s South) forward(p *position) *position {
	return p
}
func (s South) backward(p *position) *position {
	return p
}
func (s South) left(p *position) *position {
	return p
}
func (s South) right(p *position) *position {
	return p
}

func (w West) forward(p *position) *position {
	return p
}
func (w West) backward(p *position) *position {
	return p
}
func (w West) left(p *position) *position {
	return p
}
func (w West) right(p *position) *position {
	return p
}
