package specialist

import "math"

//BelongsToTheLinealFunction -
func BelongsToTheLinealFunction(x1, y1, x2, y2, x, y float64) bool {
	if (x == x1 && x == x2) || (y == y1 && y == y2) {
		return true
	}
	return ((y - y1) / (x - x1)) == ((y2 - y1) / (x2 - x1))
}

//IsInsideTheTriangle -
func IsInsideTheTriangle(x1, y1, x2, y2, x3, y3, x, y float64) bool {
	orientation1 := getOrientation(x1, y1, x2, y2, x, y)
	orientation2 := getOrientation(x2, y2, x3, y3, x, y)
	orientation3 := getOrientation(x3, y3, x1, y1, x, y)

	return orientation1 == orientation2 && orientation1 == orientation3
}

func getOrientation(x1, y1, x2, y2, x3, y3 float64) bool {
	return ((x1-x3)*(y2-y3))-((y1-y3)*(x2-x3)) >= 0
}

//GetPerimeterOfTriangle -
func GetPerimeterOfTriangle(x1, y1, x2, y2, x3, y3 float64) float64 {
	var p float64

	p += getDistanceBetweenTwoPoints(x1, y1, x2, y2)
	p += getDistanceBetweenTwoPoints(x2, y2, x3, y3)
	p += getDistanceBetweenTwoPoints(x3, y3, x1, y1)

	return math.Round(p*100) / 100
}

func getDistanceBetweenTwoPoints(x1, y1, x2, y2 float64) float64 {
	d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	return math.Round(d*100) / 100
}
