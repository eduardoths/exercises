package ex4

type Position struct {
	X float64
	Y float64
}

func (p Position) Offset(x, y float64) Position {
	pCopy := p
	pCopy.X += x
	pCopy.Y += y
	return pCopy
}
