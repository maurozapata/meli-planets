package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/meli-planets/pkg/redis"
	"github.com/meli-planets/pkg/server"
)

func main() {
	fmt.Println("Server started")

	if godotenv.Load(".env") != nil {
		panic("Error loading .env file")
	}

	db, err := initRedis()
	if err != nil {
		panic("Connect: unable to connect redisdb")
	}

	server := server.GetEngine(db)
	if server.Run(":"+os.Getenv("PORT")) != nil {
		panic("Error running server")
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
