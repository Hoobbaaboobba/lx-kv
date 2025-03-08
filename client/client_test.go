package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestNewClientRedisClient(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println(rdb)
	fmt.Println("this is working")

	if err := rdb.Set(context.Background(), "key", "value", 0).Err(); err != nil {
		panic(err)
	}
}

func TestNewClient1(t *testing.T) {
	c, err := New("localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	time.Sleep(time.Second)
	if err := c.Set(context.TODO(), "foo", 1); err != nil {
		t.Fatal(err)
	}
	val, err := c.Get(context.TODO(), "foo")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("GET =>", val)
}

func TestNewClient(t *testing.T) {
	c, err := New("localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if err := c.Set(context.TODO(), fmt.Sprintf("foo_%d", i), fmt.Sprintf("bar_%d", i)); err != nil {
			t.Fatal(err)
		}
		fmt.Println("SET =>", fmt.Sprintf("bar_%d", i))
		val, err := c.Get(context.TODO(), fmt.Sprintf("foo_%d", i))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("GET =>", val)
	}
}
