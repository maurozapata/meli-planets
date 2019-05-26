package keyfactory

var (
	baseWeatherKey = "distantGalaxy:weather:day:"
)

//GetWeather -
func GetWeather(key string) string {
	return baseWeatherKey + key
}
