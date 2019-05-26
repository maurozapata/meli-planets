package weather_test

import (
	"testing"

	"github.com/meli-planets/pkg/model"
	"github.com/meli-planets/pkg/weather"
)

func TestGetWeather(t *testing.T) {
	aux := []struct {
		planets       []*model.Planet
		expectedValue string
	}{
		{
			planets:       []*model.Planet{model.NewPlanet("", 1, true, 500, 90), model.NewPlanet("", 1, true, 1000, 270), model.NewPlanet("", 1, false, 2000, 270)},
			expectedValue: weather.DROUGHT,
		},
		{
			planets:       []*model.Planet{model.NewPlanet("", 1, true, 500, 93), model.NewPlanet("", 1, true, 1000, 204), model.NewPlanet("", 1, true, 2000, 270)},
			expectedValue: weather.OTHER,
		},
		{
			planets:       []*model.Planet{model.NewPlanet("", 1, true, 500, 0), model.NewPlanet("", 1, true, 1000, 120), model.NewPlanet("", 1, true, 2000, 240)},
			expectedValue: weather.RAIN,
		},
		{
			planets:       []*model.Planet{model.NewPlanet("", 1, true, 5, 53.13), model.NewPlanet("", 1, true, 5, 306.87), model.NewPlanet("", 1, true, 3, 0)},
			expectedValue: weather.OPTIMALPRESSUREANDTEMPERATURE,
		},
		{
			planets:       []*model.Planet{model.NewPlanet("", 1, true, 5, 53.13), model.NewPlanet("", 1, true, 3, 0)},
			expectedValue: weather.IMPOSSIBLETOPREDICT,
		},
	}

	for _, v := range aux {
		g := model.NewGalaxy(v.planets...)

		if weather.GetCurrentWeather(g) != v.expectedValue {
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

	weatherPerDay := weather.PredictExtendedWeather(g, 200)

	for day, weather := range expectedValues {
		if weatherPerDay[day] != weather {
			t.Fail()
		}
	}
}
