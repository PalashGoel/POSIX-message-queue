package main

import (
	"fmt"
	"log"
	"time"

	"github.com/syucream/posix_mq/src/posix_mq"
)

const maxTickNum = 50

func main() {
	oflag := posix_mq.O_WRONLY | posix_mq.O_CREAT
	mq, err := posix_mq.NewMessageQueue("/posix_mq_example", oflag, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer mq.Close()

	count := 0
	for {
		count++
		mq.Send([]byte(fmt.Sprintf("Hello, World : %d\n", count)), 0)
		log.Println("Sent a new message")

		if count >= maxTickNum {
			break
		}

		time.Sleep(1 * time.Second)
	}
}
