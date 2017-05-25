package main

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

type MessageHandler struct {
	msgchan chan *nsq.Message
}

func (m *MessageHandler) HandleMessage(message *nsq.Message) error {
	m.msgchan <- message
	return nil
}

func (m *MessageHandler) Process() {
	for {
		message := <-m.msgchan
		fmt.Printf(
			"ID: %s\tBody: %s\n",
			message.ID,
			string(message.Body),
		)
	}
}
