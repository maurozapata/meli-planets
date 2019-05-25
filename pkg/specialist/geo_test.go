package specialist_test

import (
	"fmt"
	"testing"

	"github.com/meli-planets/pkg/specialist"
)

func TestBelongsToTheLinealFunction(t *testing.T) {
	aux := []struct {
		x1, x2, x, y1, y2, y float64
		expectedValue        bool
	}{
		{x1: -1, y1: -1, x2: 0, y2: 0, x: 1, y: 1, expectedValue: true},
		{x1: -3, y1: 0, x2: 0, y2: 0, x: 7, y: 0, expectedValue: true},
		{x1: 0, y1: -3, x2: 0, y2: 0, x: 0, y: 4, expectedValue: true},
		{x1: -3, y1: -2, x2: -1, y2: 0, x: 1, y: 2, expectedValue: true},
		{x1: -1, y1: -1, x2: 1, y2: -1, x: 10, y: 0, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x: 0, y: 10, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x: 3, y: 3, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x: -3, y: -3, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x: -0, y: -3, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x: -1, y: 1, expectedValue: false},
	}

	for _, v := range aux {
		if specialist.BelongsToTheLinealFunction(v.x1, v.y1, v.x2, v.y2, v.x, v.y) != v.expectedValue {
			t.Fail()
		}
	}
}

func TestIsInsideTheTriangle(t *testing.T) {
	aux := []struct {
		x1, x2, x3, x, y1, y2, y3, y float64
		expectedValue                bool
	}{
		{x1: -1, y1: -1, x2: 1, y2: -1, x3: 1, y3: 1, x: 0, y: 0, expectedValue: true},
		{x1: -3, y1: -3, x2: 3, y2: -3, x3: 3, y3: 3, x: 1, y: 1, expectedValue: true},
		{x1: -1, y1: -1, x2: 1, y2: -1, x3: 1, y3: 1, x: 10, y: 0, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x3: 1, y3: 1, x: 0, y: 10, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x3: 1, y3: 1, x: 3, y: 3, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x3: 1, y3: 1, x: -3, y: -3, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x3: 1, y3: 1, x: -0, y: -3, expectedValue: false},
		{x1: -1, y1: -1, x2: 1, y2: -1, x3: 1, y3: 1, x: -1, y: 1, expectedValue: false},
	}

	for _, v := range aux {
		if specialist.IsInsideTheTriangle(v.x1, v.y1, v.x2, v.y2, v.x3, v.y3, v.x, v.y) != v.expectedValue {
			t.Fail()
		}
	}
}

func TestGetPerimeterOfTriangle(t *testing.T) {
	aux := []struct {
		x1, x2, x3, y1, y2, y3 float64
		expectedValue          float64
	}{
		{x1: -1, y1: -1, x2: 1, y2: -1, x3: 1, y3: 1, expectedValue: 6.83},
		{x1: -3, y1: -3, x2: 3, y2: -3, x3: 3, y3: 3, expectedValue: 20.49},
		{x1: -7, y1: -1, x2: 5, y2: -1, x3: 3, y3: 1, expectedValue: 25.03},
		{x1: -4, y1: 0, x2: 9, y2: -2, x3: 4, y3: 8, expectedValue: 35.64},
	}

	for _, v := range aux {
		if specialist.GetPerimeterOfTriangle(v.x1, v.y1, v.x2, v.y2, v.x3, v.y3) != v.expectedValue {
			fmt.Println(specialist.GetPerimeterOfTriangle(v.x1, v.y1, v.x2, v.y2, v.x3, v.y3))
			t.Fail()
		}
	}
}
