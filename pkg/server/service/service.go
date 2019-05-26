package service

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/meli-planets/pkg/keyfactory"
	"github.com/meli-planets/pkg/redis"
)

//GetHealthStatus returns the status of the server
func GetHealthStatus(db redis.IClient) Response {

	if err := db.Ping(); err != nil {
		return Response{
			StatusCode: http.StatusServiceUnavailable,
			Error:      errors.New("Service unavailable"),
		}
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       map[string]string{"Body": "Live long and prosper!"},
	}
}

//GetWeather returns the weather of the given day
func GetWeather(db redis.IClient, day string) Response {
	if err := validateParamDay(day); err != nil {
		return Response{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}

	value, err := db.Get(keyfactory.GetWeather(day))
	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
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

func validateParamDay(day string) error {
	if day == "" {
		return errors.New("Param:day can't be empty")
	}

	if _, err := strconv.Atoi(day); err != nil {
		return errors.New("Param:day must be integer")
	}

	return nil
}
