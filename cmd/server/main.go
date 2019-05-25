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
		fmt.Println("Error loading .env file")
		return
	}

	db, err := initRedis()
	if err != nil {
		fmt.Println("Unable to connect redisdb")
		return
	}

	server := server.GetEngine(db)
	if server.Run(":"+os.Getenv("PORT")) != nil {
		fmt.Println("Error running server")
		return
	}
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
