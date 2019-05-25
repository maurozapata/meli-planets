package utils

var (
	baseWeatherKey = "distantGalaxy:weather:"
)

//GetWeatherKey -
func GetWeatherKey(key string) string {
	return baseWeatherKey + key
}
