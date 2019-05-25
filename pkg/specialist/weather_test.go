package specialist_test

import (
	"testing"

	"github.com/meli-planets/pkg/model"
	"github.com/meli-planets/pkg/specialist"
)

func TestGetWeather(t *testing.T) {
	p1 := model.NewPlanet("", 1, true, 500, 90)
	p2 := model.NewPlanet("", 1, true, 1000, 270)
	p3 := model.NewPlanet("", 1, true, 2000, 270)
	g := model.NewGalaxy(p1, p2, p3)

	if specialist.GetCurrentWeather(g) != specialist.DROUGHT {
		t.Fail()
	}
}
