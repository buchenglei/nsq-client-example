package main

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

func producer(topic string) {
	// 创建一个生产者，直接连接的是nsqd
	producer, err := nsq.NewProducer(NsqdHost, nsq.NewConfig())
	defer producer.Stop()
	if err != nil {
		panic(err.Error())
	}

	if err = producer.Ping(); err != nil {
		panic(err.Error())
	}

	var publishMsg string

	for publishMsg != "exit" {
		fmt.Print("Publish Message: ")
		fmt.Scan(&publishMsg)

		err = producer.Publish(topic, []byte(publishMsg))
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("> [%s] publish success\n", publishMsg)
	}
}
