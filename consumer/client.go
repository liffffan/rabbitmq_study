package main

import (
	"fmt"
	"log"
	"rabbitmq_study/config"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	forever := make(chan bool)

	for routine := 0; routine < 5; routine++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()

			q, err := ch.QueueDeclare(
				config.QUEUENAME,
				false, //durable
				false,
				false,
				false,
				nil,
			)

			failOnError(err, "Failed to declare a queue")

			msgs, err := ch.Consume(
				q.Name,
				"MsgWorkConsumer",
				false, //Auto Ack
				false,
				false,
				false,
				nil,
			)

			if err != nil {
				log.Fatal(err)
			}

			for msg := range msgs {
				log.Printf("In %d consume a message: %s\n", 0, msg.Body)
				log.Printf("Done")
				msg.Ack(false) //Ack
			}

		}(routine)
	}

	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
