package ex4

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type Pen struct {
	Position
	IsDown bool
	Type   int
}

func (p *Pen) Down() {
	p.IsDown = true
}

func (p *Pen) Up() {
	p.IsDown = false
}
func (p *Pen) Move(direction Direction, distance float64) {
	offset := map[Direction]Position{
		NORTH: p.Position.Offset(0, float64(distance)),
		EAST:  p.Position.Offset(float64(distance), 0),
		SOUTH: p.Position.Offset(0, -float64(distance)),
		WEST:  p.Position.Offset(-float64(distance), 0),
	}
	p.Position = offset[direction]
}
