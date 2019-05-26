package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/meli-planets/pkg/model"
	"github.com/meli-planets/pkg/redis"
	"github.com/meli-planets/pkg/utils"
	"github.com/meli-planets/pkg/weather"
)

func main() {
	if godotenv.Load(".env") != nil {
		fmt.Println("Error loading .env file")
		return
	}

	db, err := initRedis()
	if err != nil {
		fmt.Println("Connect: unable to connect redisdb")
		return
	}

	ferengi := model.NewPlanet("Ferengi", 1, true, 500, 0)
	betasoide := model.NewPlanet("Betasoide", 3, true, 2000, 0)
	vulcano := model.NewPlanet("Vulcano", 5, false, 1000, 0)

	distantGalaxy := model.NewGalaxy(ferengi, betasoide, vulcano)

	weatherPerDay := weather.PredictExtendedWeather(distantGalaxy, 3650)
	for key, value := range weatherPerDay {
		if db.Set(utils.GetWeatherKey(strconv.Itoa(key)), value) != nil {
			panic("Set: unable to connect redisdb")
		}
	}

}

func initRedis() (*redis.Client, error) {
	database, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	params := redis.Params{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: database,
	}

	return redis.Init(params)
}
