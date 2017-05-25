package main

import (
	"flag"
)

var (
	LookupdHost string
	NsqdHost    string
)

func main() {
	mode := flag.Int("mode", 0, "指明是\n\t 1 - 生产者 \n\t 0 - 消费者(默认)")
	topic := flag.String("topic", "default", "指明要操作的topic")
	channel := flag.String("channel", "default", "指明要操作的channel")
	lookupd := flag.String("lookupd", "127.0.0.1:4161", "指明lookupd的地址，默认127.0.0.1:4161")
	nsqd := flag.String("nsqd", "127.0.0.1:4150", "指明nsqd的地址，默认127.0.0.1:4150")

	flag.Parse()

	LookupdHost = *lookupd
	NsqdHost = *nsqd

	if *mode == 0 {
		consumer(*topic, *channel)
	} else {
		producer(*topic)
	}
}
