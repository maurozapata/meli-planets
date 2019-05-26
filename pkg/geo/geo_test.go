package geo_test

import (
	"fmt"
	"testing"

	"github.com/meli-planets/pkg/geo"
)

func TestBelongsToTheLinealFunction(t *testing.T) {
	aux := []struct {
		p1, p2, p     geo.Point
		expectedValue bool
	}{
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(0, 0), p: geo.NewPoint(1, 1), expectedValue: true},
		{p1: geo.NewPoint(-3, 0), p2: geo.NewPoint(0, 0), p: geo.NewPoint(7, 0), expectedValue: true},
		{p1: geo.NewPoint(0, -3), p2: geo.NewPoint(0, 0), p: geo.NewPoint(0, 4), expectedValue: true},
		{p1: geo.NewPoint(-3, -2), p2: geo.NewPoint(-1, 0), p: geo.NewPoint(1, 2), expectedValue: true},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p: geo.NewPoint(10, 0), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p: geo.NewPoint(0, 10), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p: geo.NewPoint(3, 3), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p: geo.NewPoint(-3, -3), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p: geo.NewPoint(-0, -3), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p: geo.NewPoint(-1, 1), expectedValue: false},
	}

	for _, v := range aux {
		if geo.BelongsToTheLinealFunction(v.p1, v.p2, v.p) != v.expectedValue {
			t.Fail()
		}
	}
}

func TestIsInsideTheTriangle(t *testing.T) {
	aux := []struct {
		p1, p2, p3, p geo.Point
		expectedValue bool
	}{
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p3: geo.NewPoint(1, 1), p: geo.NewPoint(0, 0), expectedValue: true},
		{p1: geo.NewPoint(-3, -3), p2: geo.NewPoint(3, -3), p3: geo.NewPoint(3, 3), p: geo.NewPoint(1, 1), expectedValue: true},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p3: geo.NewPoint(1, 1), p: geo.NewPoint(10, 0), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p3: geo.NewPoint(1, 1), p: geo.NewPoint(0, 10), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p3: geo.NewPoint(1, 1), p: geo.NewPoint(3, 3), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p3: geo.NewPoint(1, 1), p: geo.NewPoint(-3, -3), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p3: geo.NewPoint(1, 1), p: geo.NewPoint(-0, -3), expectedValue: false},
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p3: geo.NewPoint(1, 1), p: geo.NewPoint(-1, 1), expectedValue: false},
	}

	for _, v := range aux {
		if geo.IsInsideTheTriangle(v.p1, v.p2, v.p3, v.p) != v.expectedValue {
			t.Fail()
		}
	}
}

func TestGetPerimeterOfTriangle(t *testing.T) {
	aux := []struct {
		p1, p2, p3    geo.Point
		expectedValue float64
	}{
		{p1: geo.NewPoint(-1, -1), p2: geo.NewPoint(1, -1), p3: geo.NewPoint(1, 1), expectedValue: 6.83},
		{p1: geo.NewPoint(-3, -3), p2: geo.NewPoint(3, -3), p3: geo.NewPoint(3, 3), expectedValue: 20.49},
		{p1: geo.NewPoint(-7, -1), p2: geo.NewPoint(5, -1), p3: geo.NewPoint(3, 1), expectedValue: 25.03},
		{p1: geo.NewPoint(-4, 0), p2: geo.NewPoint(9, -2), p3: geo.NewPoint(4, 8), expectedValue: 35.64},
	}

	for _, v := range aux {
		if geo.GetPerimeterOfTriangle(v.p1, v.p2, v.p3) != v.expectedValue {
			fmt.Println(geo.GetPerimeterOfTriangle(v.p1, v.p2, v.p3))
			t.Fail()
		}
	}
}
