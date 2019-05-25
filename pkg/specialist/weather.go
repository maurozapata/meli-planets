package specialist

import "github.com/meli-planets/pkg/model"

//GetWeather -
func GetWeather(galaxy *model.Galaxy) string {
	x1, y1 := galaxy.Planets[0].GetCoordinates()
	x2, y2 := galaxy.Planets[1].GetCoordinates()

	x3, y3 := galaxy.Planets[2].GetCoordinates()

	if BelongsToTheLinealFunction(x1, y1, x2, y2, x3, y3) {
		if BelongsToTheLinealFunction(x1, y1, x2, y2, 0, 0) {
			return "SEQUIA"
		}
		return "PRESION Y TEMPERATURA OPTIMA"
	}

	if IsInsideTheTriangle(x1, y1, x2, y2, x3, y3, 0, 0) {
		return "LLUVIA"
	}

	return "OTRO"
}
