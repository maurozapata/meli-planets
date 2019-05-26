package model

import (
	"math"

	"github.com/meli-planets/pkg/geo"
)

//Planet -
type Planet struct {
	Name            string
	AngularVelocity float64
	Clockwise       bool
	Radio           float64
	Angle           float64
}

//NewPlanet -
func NewPlanet(name string, angularVelocity float64, clockwise bool, radio float64, angle float64) *Planet {
	return &Planet{
		Name:            name,
		AngularVelocity: angularVelocity,
		Clockwise:       clockwise,
		Radio:           radio,
		Angle:           angle,
	}
}

//Aging -
func (p *Planet) Aging(days float64) {
	var total float64

	if p.Clockwise {
		total = p.Angle - (days * p.AngularVelocity)
	} else {
		total = p.Angle + (days * p.AngularVelocity)
	}

	angle := math.Mod(total, 360)

	if angle >= 0 {
		p.Angle = angle
	} else {
		p.Angle = angle + 360
	}
}

//GetCoordinates -
func (p *Planet) GetCoordinates() geo.Point {
	x := p.Radio * math.Cos(p.Angle*math.Pi/180)
	y := p.Radio * math.Sin(p.Angle*math.Pi/180)

	return geo.NewPoint(math.Round(x*100)/100, math.Round(y*100)/100)

}
