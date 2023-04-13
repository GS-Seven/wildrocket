package wdb

import (
	"fmt"
	"testing"
)

func TestInitRedisClient(t *testing.T) {
	client, err := InitRedisClient("127.0.0.1:7379", "123456", 0)
	if err != nil {
		t.Fatal(err)
	}

	err = client.Set("a", 1, 0).Err()
	if err != nil {
		t.Fatal(err)
	}

	result, err := client.Get("a").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
