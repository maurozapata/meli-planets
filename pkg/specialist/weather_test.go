package specialist_test

import (
	"testing"

	"github.com/meli-planets/pkg/model"
	"github.com/meli-planets/pkg/specialist"
)

func TestGetWeather(t *testing.T) {
	aux := []struct {
		p1, p2, p3    *model.Planet
		expectedValue string
	}{
		{
			p1:            model.NewPlanet("", 1, true, 500, 90),
			p2:            model.NewPlanet("", 1, true, 1000, 270),
			p3:            model.NewPlanet("", 1, false, 2000, 270),
			expectedValue: specialist.DROUGHT,
		},
		{
			p1:            model.NewPlanet("", 1, true, 500, 93),
			p2:            model.NewPlanet("", 1, true, 1000, 204),
			p3:            model.NewPlanet("", 1, true, 2000, 270),
			expectedValue: specialist.OTHER,
		},
		{
			p1:            model.NewPlanet("", 1, true, 500, 0),
			p2:            model.NewPlanet("", 1, true, 1000, 120),
			p3:            model.NewPlanet("", 1, true, 2000, 240),
			expectedValue: specialist.RAIN,
		},
		{
			p1:            model.NewPlanet("", 1, true, 5, 53.13),
			p2:            model.NewPlanet("", 1, true, 5, 306.87),
			p3:            model.NewPlanet("", 1, true, 3, 0),
			expectedValue: specialist.OPTIMALPRESSUREANDTEMPERATURE,
		},
	}

	for _, v := range aux {
		g := model.NewGalaxy(v.p1, v.p2, v.p3)

		if specialist.GetCurrentWeather(g) != v.expectedValue {
			t.Fail()
		}
	}
}

func TestPredictExtendedWeather(t *testing.T) {
	p1 := model.NewPlanet("", 1, true, 5, 53.13)
	p2 := model.NewPlanet("", 5, false, 5, 306.87)
	p3 := model.NewPlanet("", 13, true, 3, 0)

	g := model.NewGalaxy(p1, p2, p3)

	expectedValues := map[int]string{
		194: "otro",
		158: "lluvia",
		174: "otro",
		179: "presion y temperatura optima",
		35:  "pico de lluvia",
		39:  "lluvia",
		79:  "otro",
	}

	weatherPerDay := specialist.PredictExtendedWeather(g, 200)

	for day, weather := range expectedValues {
		if weatherPerDay[day] != weather {
			t.Fail()
		}
	}
}
