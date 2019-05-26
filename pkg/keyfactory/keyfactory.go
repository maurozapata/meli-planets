package keyfactory

var (
	baseWeatherKey = "distantGalaxy:weather:day:"
)

//GetWeather return a weater key
func GetWeather(key string) string {
	return baseWeatherKey + key
}
