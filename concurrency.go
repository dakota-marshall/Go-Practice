package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)
	go count("sheep", channel)

	for msg_from_channel := range channel {
		fmt.Println(msg_from_channel)
	}

}

func count(thing string, channel chan string) {
	for count := 1; count <= 10; count++ {
		message := thing
		channel <- message
		time.Sleep(time.Millisecond * 500)
	}

	close(channel)

}
