package redis

import (
    "github.com/garyburd/redigo/redis"
	"time"
	"flag"
)

type RedisCli struct {
    conn redis.Conn
}

var (
	pool *redis.Pool
	redisServer = flag.String("redisServer", ":6379", "")
	redisPassword = flag.String("redisPassword", "foo123", "")
)

func init(){
	flag.Parse()
	pool = newPool(*redisServer, *redisPassword)
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func SetValue(key string, value string, expiration ...interface{}) error {
	redisConnection := pool.Get()
	defer redisConnection.Close()
	_, err := redisConnection.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisConnection.Do("EXPIRE", key, expiration[0])
	}
	return err
}

func GetValue(key string) (interface{}, error) {
	redisConnection := pool.Get()
	defer redisConnection.Close()
	return redisConnection.Do("GET", key)
}
