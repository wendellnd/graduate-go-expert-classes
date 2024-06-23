package main

import "github.com/wendellnd/graduate-go-expert-classes/events/pkg/rabbitmq"

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	rabbitmq.Publish(ch, "amq.direct", "Hello, World!")
}
