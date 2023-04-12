package main

import (
	"fmt"
	"log"
	"time"

	cache "vscode2/example"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second*5)
	userId, err := cache.Get("userId")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userId)

	time.Sleep(time.Second * 4)

	userId, err = cache.Get("userId")

	if err != nil {
		log.Fatal(err)
	}
}
