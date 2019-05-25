package service

//Response -
type Response struct {
	StatusCode int
	Body       interface{}
	Error      error
}
