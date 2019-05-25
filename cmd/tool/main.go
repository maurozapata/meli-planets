package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/meli-planets/pkg/model"
	"github.com/meli-planets/pkg/redis"
	"github.com/meli-planets/pkg/specialist"
)

func main() {
	fmt.Println("Tool started")

	if godotenv.Load(".env") != nil {
		fmt.Println("Error loading .env file")
		return
	}

	db, err := initRedis()
	if err != nil {
		fmt.Println("Unable to connect redisdb")
		return
	}

	predictWeather(db)

}

func predictWeather(db *redis.Client) {
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

		err := db.Set(strconv.Itoa(i), weather)
		if err != nil {
			fmt.Println(err)
		}
	}
	periods[yesterdaysweather]++

	fmt.Println(periods)
}

func initRedis() (*redis.Client, error) {
	database, _ := strconv.Atoi(os.Getenv("DATABASE"))
	params := redis.Params{
		Address:  os.Getenv("ADDRESS"),
		Password: os.Getenv("PASSWORD"),
		Database: database,
	}

	return redis.Init(params)
}
