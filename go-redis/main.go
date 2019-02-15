package main

import (
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland2", "artist", "Jimi Hendrix2", "price", 4.952, "likes", 82)
	if resp.Err != nil {
		log.Fatal(resp.Err)
	}
	fmt.Println("Electric Ladyland added!")
}
