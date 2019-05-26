package model

//Galaxy -
type Galaxy struct {
	Planets []*Planet
}

//NewGalaxy returns a Galaxy
func NewGalaxy(planets ...*Planet) *Galaxy {
	return &Galaxy{
		Planets: planets,
	}
}

//Aging makes the galaxy age the amount of days given
func (g *Galaxy) Aging(days float64) {
	for _, p := range g.Planets {
		p.Aging(days)
	}
}
