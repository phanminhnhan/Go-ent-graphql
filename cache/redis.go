package cache
import (
	"encoding/json"
	"time"
	"github.com/go-redis/redis"
)

var (
	Client = &RedisClient{}
)

type RedisClient struct {
	c *redis.Client
}

//GetClient get the redis client
func Initialize() *RedisClient {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	Client.c = c
	return Client
}

//GetKey get key
func (client RedisClient) GetKey(key string, src interface{}) error {
	val, err := client.c.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}

//SetKey set key
func (client RedisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}