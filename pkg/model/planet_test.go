package model_test

import (
	"testing"

	"github.com/meli-planets/pkg/model"
)

func TestAgingClockwise(t *testing.T) {
	planet := model.NewPlanet("Ferengi", 1, true, 500, 0)

	aux := []struct {
		days          float64
		expectedValue float64
	}{
		{
			days:          1,
			expectedValue: 359,
		},
		{
			days:          7,
			expectedValue: 352,
		},
		{
			days:          30,
			expectedValue: 322,
		},
		{
			days:          360,
			expectedValue: 322,
		},
	}

	for _, v := range aux {
		planet.Aging(v.days)

		if planet.Angle != v.expectedValue {
			t.Fail()
		}
	}
}

func TestAgingCounterclockwise(t *testing.T) {
	planet := model.NewPlanet("Ferengi", 4, false, 500, 3)

	aux := []struct {
		days          float64
		expectedValue float64
	}{
		{
			days:          3,
			expectedValue: 15,
		},
		{
			days:          7,
			expectedValue: 43,
		},
		{
			days:          30,
			expectedValue: 163,
		},
		{
			days:          360,
			expectedValue: 163,
		},
	}

	for _, v := range aux {
		planet.Aging(v.days)

		if planet.Angle != v.expectedValue {
			t.Fail()
		}
	}
}
