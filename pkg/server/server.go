package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meli-planets/pkg/redis"
	"github.com/meli-planets/pkg/server/service"
)

var _db redis.IClient

//GetEngine -
func GetEngine(db redis.IClient) *gin.Engine {
	_db = db

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/health", getHealthStatus)
	r.GET("/weather/:day", getWeather)

	return r
}

func getHealthStatus(c *gin.Context) {
	sendResponse(c, service.GetHealthStatus(_db))
}

func getWeather(c *gin.Context) {
	sendResponse(c, service.GetWeather(_db, c.Param("day")))
}

func sendResponse(c *gin.Context, response service.Response) {
	obj := gin.H{}

	if response.Error != nil {
		obj["error"] = response.Error.Error()
	}
	if response.Body != nil {
		obj["body"] = response.Body
	}
	obj["status"] = http.StatusText(response.StatusCode)

	c.JSON(response.StatusCode, obj)
}
