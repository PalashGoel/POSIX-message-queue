package main

import (
	"log"
        "time"
	"github.com/syucream/posix_mq/src/posix_mq"
)

const maxTickNum = 50

func main() {
	time.Sleep(5 * time.Second)
	oflag := posix_mq.O_RDONLY
	mq, err := posix_mq.NewMessageQueue("/posix_mq_example", oflag, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer mq.Close()

	log.Println("Start receiving messages")

	count := 0
	for {
		count++

		msg, _, err := mq.Receive()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(string(msg))

		if count >= maxTickNum {
			break
		}
	}
}
