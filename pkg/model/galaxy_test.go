package model_test

import (
	"testing"

	"github.com/meli-planets/pkg/model"
)

func TestAgingGalaxy(t *testing.T) {
	planet1 := model.NewPlanet("", 1, true, 500, 0)
	planet2 := model.NewPlanet("", 1, false, 2600, 200)
	galaxy := model.NewGalaxy(planet1, planet2)

	galaxy.Aging(21)

	if planet1.Angle != 339 || planet2.Angle != 221 {

		t.Fail()
	}
}
