package specialist

import (
	"fmt"

	"github.com/meli-planets/pkg/model"
)

const (
	//RAIN -
	RAIN = "lluvia"
	//MAXRAIN -
	MAXRAIN = "pico de lluvia"
	//DROUGHT -
	DROUGHT = "sequia"
	//OTHER -
	OTHER = "otro"
	//OPTIMALPRESSUREANDTEMPERATURE -
	OPTIMALPRESSUREANDTEMPERATURE = "presion y temperatura optima"
)

//PredictExtendedWeather -
func PredictExtendedWeather(galaxy *model.Galaxy, days int) map[int]string {
	periods := map[string]int{}

	weatherPerDay := map[int]string{}

	var maxPerimeter float64
	maxPerimeterDays := []int{}

	yesterdaysweather := ""
	for i := 0; i < days; i++ {
		galaxy.Aging(1)

		weather := GetCurrentWeather(galaxy)
		if weather != yesterdaysweather && yesterdaysweather != "" {
			periods[weather]++
		}

		if weather == RAIN {
			perimeter := getPerimeterOfGalaxy(galaxy)
			if perimeter > maxPerimeter {
				maxPerimeter = perimeter
				maxPerimeterDays = []int{i}
			} else if perimeter == maxPerimeter {
				maxPerimeterDays = append(maxPerimeterDays, i)
			}
		}

		weatherPerDay[i] = weather

		yesterdaysweather = weather
	}
	periods[yesterdaysweather]++

	for _, v := range maxPerimeterDays {
		weatherPerDay[v] = MAXRAIN
	}

	fmt.Println(periods)

	return weatherPerDay
}

func getPerimeterOfGalaxy(galaxy *model.Galaxy) float64 {
	x1, y1 := galaxy.Planets[0].GetCoordinates()
	x2, y2 := galaxy.Planets[1].GetCoordinates()
	x3, y3 := galaxy.Planets[2].GetCoordinates()

	return GetPerimeterOfTriangle(x1, y1, x2, y2, x3, y3)
}

//GetCurrentWeather -
func GetCurrentWeather(galaxy *model.Galaxy) string {
	x1, y1 := galaxy.Planets[0].GetCoordinates()
	x2, y2 := galaxy.Planets[1].GetCoordinates()
	x3, y3 := galaxy.Planets[2].GetCoordinates()

	if BelongsToTheLinealFunction(x1, y1, x2, y2, x3, y3) {
		if BelongsToTheLinealFunction(x1, y1, x2, y2, 0, 0) {
			return DROUGHT
		}
		return OPTIMALPRESSUREANDTEMPERATURE
	} else if IsInsideTheTriangle(x1, y1, x2, y2, x3, y3, 0, 0) {
		return RAIN
	}

	return OTHER
}
