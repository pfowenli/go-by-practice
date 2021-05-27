package main

import (
	"fmt"
	"time"
)

func main() {
	cache1 := NewRedisCache("localhost:6379", "", 5*time.Minute)
	key := "breakfast"

	cache1.SetValue(key, "oat")
	fmt.Printf("value(%v) = %v\n", key, cache1.GetValue(key))
	cache1.AppendValue(key, "ly")
	fmt.Printf("value(%v) = %v\n", key, cache1.GetValue(key))

	cache2 := NewRedisCache("localhost:6379", "", 2021*time.Second)
	key = "pandenmic"

	cache2.SetValue(key, "COVID")
	fmt.Printf("value(%v) = %v\n", key, cache2.GetValue(key))
	cache2.AppendValue(key, 19)
	fmt.Printf("value(%v) = %v\n", key, cache2.GetValue(key))
}
