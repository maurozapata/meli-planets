package redis

import (
	"github.com/go-redis/redis"
)

//IClient -
type IClient interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}

//Client -
type Client struct {
	db *redis.Client
}

//Params -
type Params struct {
	Address  string
	Password string
	Database int
}

//Init -
func Init(params Params) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     params.Address,
		Password: params.Password,
		DB:       params.Database,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Client{db: client}, nil
}

//Get -
func (c *Client) Get(key string) (interface{}, error) {
	return c.db.Get(key).Result()
}

//Set -
func (c *Client) Set(key string, value interface{}) error {
	return c.db.Set(key, value, 0).Err()
}
