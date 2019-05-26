package weather

import (
	"fmt"

	"github.com/meli-planets/pkg/geo"
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
	//IMPOSSIBLETOPREDICT -
	IMPOSSIBLETOPREDICT = "imposible predecir"
)

//PredictExtendedWeather -
func PredictExtendedWeather(galaxy *model.Galaxy, days int) map[int]string {
	periods := map[string]int{}

	weatherPerDay := map[int]string{}

	var maxPerimeter float64
	maxRainDays := []int{}

	yesterdaysweather := ""
	for i := 1; i <= days; i++ {
		galaxy.Aging(1)

		weather := GetCurrentWeather(galaxy)
		if weather != yesterdaysweather && yesterdaysweather != "" {
			periods[weather]++
		}

		if weather == RAIN {
			perimeter := getPerimeterOfGalaxy(galaxy)
			if perimeter > maxPerimeter {
				maxPerimeter = perimeter
				maxRainDays = []int{i}
			} else if perimeter == maxPerimeter {
				maxRainDays = append(maxRainDays, i)
			}
		}

		weatherPerDay[i] = weather

		yesterdaysweather = weather
	}
	periods[yesterdaysweather]++

	for _, v := range maxRainDays {
		weatherPerDay[v] = MAXRAIN
	}

	for weather, cant := range periods {
		fmt.Println(weather, "-", cant)
	}
	fmt.Println("max rain days -", maxRainDays)
	fmt.Println("max perimeter -", maxPerimeter)

	return weatherPerDay
}

func getPerimeterOfGalaxy(galaxy *model.Galaxy) float64 {

	if len(galaxy.Planets) < 3 {
		return 0
	}

	p1 := galaxy.Planets[0].GetCoordinates()
	p2 := galaxy.Planets[1].GetCoordinates()
	p3 := galaxy.Planets[2].GetCoordinates()

	return geo.GetPerimeterOfTriangle(p1, p2, p3)
}

//GetCurrentWeather -
func GetCurrentWeather(galaxy *model.Galaxy) string {

	if len(galaxy.Planets) < 3 {
		return IMPOSSIBLETOPREDICT
	}

	p1 := galaxy.Planets[0].GetCoordinates()
	p2 := galaxy.Planets[1].GetCoordinates()
	p3 := galaxy.Planets[2].GetCoordinates()

	if geo.BelongsToTheLinealFunction(p1, p2, p3) {
		if geo.BelongsToTheLinealFunction(p1, p2, geo.NewPoint(0, 0)) {
			return DROUGHT
		}
		return OPTIMALPRESSUREANDTEMPERATURE
	} else if geo.IsInsideTheTriangle(p1, p2, p3, geo.NewPoint(0, 0)) {
		return RAIN
	}

	return OTHER
}
