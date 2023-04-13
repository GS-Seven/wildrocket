package wdb

import "gopkg.in/redis.v5"

// InitRedisClient 初始化redis 连接
func InitRedisClient(addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
