package service

import (
	"net/http"

	"github.com/meli-planets/pkg/redis"
	"github.com/meli-planets/pkg/utils"
)

//GetHealthStatus -
func GetHealthStatus(db redis.IClient) Response {
	return Response{}
}

//GetWeather -
func GetWeather(db redis.IClient, day string) Response {
	val, err := db.Get(utils.GetWeatherKey(day))

	if err != nil {
		return Response{
			StatusCode: http.StatusNotFound,
			Error:      err,
		}
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       val,
	}
}
