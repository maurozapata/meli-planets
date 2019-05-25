package service

import (
	"errors"
	"net/http"

	"github.com/meli-planets/pkg/redis"
	"github.com/meli-planets/pkg/utils"
)

//GetHealthStatus -
func GetHealthStatus(db redis.IClient) Response {

	if err := db.Ping(); err != nil {
		return Response{
			StatusCode: http.StatusServiceUnavailable,
			Error:      errors.New("Service unavailable"),
		}
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       "Live long and prosper!",
	}
}

//GetWeather -
func GetWeather(db redis.IClient, day string) Response {
	value, err := db.Get(utils.GetWeatherKey(day))

	if err != nil {
		return Response{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("Wheather not found"),
		}
	}

	return Response{
		StatusCode: http.StatusOK,
		Body: map[string]interface{}{
			"dia":   day,
			"clima": value,
		},
	}
}
