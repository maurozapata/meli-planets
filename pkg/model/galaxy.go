package model

//Galaxy -
type Galaxy struct {
	Planets []*Planet
}

//NewGalaxy -
func NewGalaxy(planets ...*Planet) *Galaxy {
	return &Galaxy{
		Planets: planets,
	}
}

//Aging -
func (g *Galaxy) Aging(days float64) {
	for _, p := range g.Planets {
		p.Aging(days)
	}
}
