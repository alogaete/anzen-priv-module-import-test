package main

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/alogaete/anzen-messaging-module"
)

func main() {
	uri, err := url.Parse(os.Getenv("CLOUDMQTT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "test"
	}

	go messaging.Listen(uri, topic)

	client := messaging.Connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, t.String())
	}
}