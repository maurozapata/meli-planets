package geo

import "math"

//Point -
type Point struct {
	X, Y float64
}

//NewPoint returns a Point
func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

//BelongsToTheLinealFunction returns true if p belongs to the line formed by p1 and p2
func BelongsToTheLinealFunction(p1, p2, p Point) bool {
	if (p.X == p1.X && p.X == p2.X) || (p.Y == p1.Y && p.Y == p2.Y) {
		return true
	}
	return ((p.Y - p1.Y) / (p.X - p1.X)) == ((p2.Y - p1.Y) / (p2.X - p1.X))
}

//IsInsideTheTriangle returns true if p is inside the triangle formed by p1, p2 and p3
func IsInsideTheTriangle(p1, p2, p3, p Point) bool {
	orientation1 := getOrientation(p1, p2, p)
	orientation2 := getOrientation(p2, p3, p)
	orientation3 := getOrientation(p3, p1, p)

	return orientation1 == orientation2 && orientation1 == orientation3
}

func getOrientation(p1, p2, p3 Point) bool {
	return ((p1.X-p3.X)*(p2.Y-p3.Y))-((p1.Y-p3.Y)*(p2.X-p3.X)) >= 0
}

//GetPerimeterOfTriangle returns the perimeter of the triangle formed by points p1, p2 and p3
func GetPerimeterOfTriangle(p1, p2, p3 Point) float64 {
	var p float64

	p += getDistanceBetweenTwoPoints(p1, p2)
	p += getDistanceBetweenTwoPoints(p2, p3)
	p += getDistanceBetweenTwoPoints(p3, p1)

	return math.Round(p*100) / 100
}

//PolarToCartesian converts a polar coordinate to a cartesian coordinate
func PolarToCartesian(angle, radio float64) Point {
	x := radio * math.Cos(angle*math.Pi/180)
	y := radio * math.Sin(angle*math.Pi/180)

	return NewPoint(math.Round(x*100)/100, math.Round(y*100)/100)
}

func getDistanceBetweenTwoPoints(p1, p2 Point) float64 {
	d := math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))

	return math.Round(d*100) / 100
}
