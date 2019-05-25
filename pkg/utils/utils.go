package utils

var (
	baseWeatherKey = "distantGalaxy:weather:day:"
)

//GetWeatherKey -
func GetWeatherKey(key string) string {
	return baseWeatherKey + key
}
