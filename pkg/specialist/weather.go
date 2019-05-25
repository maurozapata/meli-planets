package specialist

import (
	"fmt"

	"github.com/meli-planets/pkg/model"
)

//PredictExtendedWeather -
func PredictExtendedWeather(galaxy *model.Galaxy, days int) map[int]string {
	periods := map[string]int{}
	weatherPerDay := map[int]string{}

	yesterdaysweather := ""
	for i := 0; i < days; i++ {
		galaxy.Aging(1)

		weather := GetCurrentWeather(galaxy)
		if weather != yesterdaysweather && yesterdaysweather != "" {
			periods[weather]++
		}

		weatherPerDay[i] = weather

		yesterdaysweather = weather
	}
	periods[yesterdaysweather]++

	fmt.Println(periods)

	return weatherPerDay
}

//GetCurrentWeather -
func GetCurrentWeather(galaxy *model.Galaxy) string {
	x1, y1 := galaxy.Planets[0].GetCoordinates()
	x2, y2 := galaxy.Planets[1].GetCoordinates()
	x3, y3 := galaxy.Planets[2].GetCoordinates()

	if BelongsToTheLinealFunction(x1, y1, x2, y2, x3, y3) {
		if BelongsToTheLinealFunction(x1, y1, x2, y2, 0, 0) {
			return "sequia"
		}
		return "presion y temperatura optima"
	}

	if IsInsideTheTriangle(x1, y1, x2, y2, x3, y3, 0, 0) {
		return "lluvia"
	}

	return "otro"
}
