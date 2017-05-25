package main

import nsq "github.com/nsqio/go-nsq"

func consumer(topic, channel string) {
	consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		panic(err.Error())
	}

	handler := new(MessageHandler)
	handler.msgchan = make(chan *nsq.Message, 100)
	consumer.AddHandler(nsq.HandlerFunc(handler.HandleMessage))
	err = consumer.ConnectToNSQLookupd(LookupdHost)
	if err != nil {
		panic(err.Error())
	}

	handler.Process()
}
