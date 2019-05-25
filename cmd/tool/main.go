package main

import (
	"fmt"

	"github.com/meli-planets/pkg/model"
	"github.com/meli-planets/pkg/specialist"
)

func main() {
	ferengi := model.NewPlanet("Ferengi", 1, true, 500, 0)
	betasoide := model.NewPlanet("Betasoide", 3, true, 2000, 0)
	vulcano := model.NewPlanet("Vulcano", 5, false, 1000, 0)

	periods := map[string]int{}

	distantGalaxy := model.NewGalaxy(ferengi, vulcano, betasoide)
	yesterdaysweather := ""
	for i := 0; i < 3650; i++ {
		distantGalaxy.Aging(1)

		weather := specialist.GetWeather(distantGalaxy)
		if weather != yesterdaysweather && yesterdaysweather != "" {
			periods[weather]++
		}
		yesterdaysweather = weather
		// if ferengi.Angle == betasoide.Angle || ferengi.Angle == vulcano.Angle || betasoide.Angle == vulcano.Angle {
		// 	fmt.Println(i, "-", ferengi.Angle, betasoide.Angle, vulcano.Angle, weather)
		// }
	}
	periods[yesterdaysweather]++

	fmt.Println(periods)

}
