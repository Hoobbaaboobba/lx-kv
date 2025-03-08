package client

import (
	"context"
	"fmt"
	"log"
	"testing"
)

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
