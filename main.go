package main

import (
	"fmt"
	"rabbitmq_study/config"
	"rabbitmq_study/rabbitmq"
)

type TestPro struct {
	msgContent string
}

// 实现发送者
func (t *TestPro) MsgContent() string {
	return t.msgContent
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

func main() {

	fmt.Println("da")
	fmt.Println(config.Conf)

	return
	msg := fmt.Sprintf("这是测试任务")
	t := &TestPro{
		msg,
	}
	queueExchange := &rabbitmq.QueueExchange{
		"test.rabbit",
		"rabbit.key",
		"test.rabbit.mq",
		"direct",
	}
	mq := rabbitmq.New(queueExchange)
	mq.RegisterProducer(t)
	mq.RegisterReceiver(t)
	mq.Start()
	/*
		// Dial接受AMQP URI格式的字符串，并使用PlainAuth返回新的TCP连接。 默认情况下，服务器心跳间隔为10秒，并将握手期限设置为30秒。
		conn, err := amqp.Dial(config.RMQADDR)
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		// 创建信道
		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		// 创建队列
		q, err := ch.QueueDeclare(
			config.QUEUENAME,		//Queue name
			false,			// durable
			false,		// delete when unused
			false,// exclusive
			false,            // no-wait
			nil,              // arguments
		)

		failOnError(err, "Failed to declare a queue")

		for i := 0; i < 500; i++ {
			msgBody := fmt.Sprintf("Message_%d", i)

			err = ch.Publish(
				"",     //exchange
				q.Name, //routing key
				false,
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent, //Msg set as persistent
					ContentType:  "text/plain",
					Body:         []byte(msgBody),
				})

			log.Printf(" [x] Sent %s", msgBody)
			failOnError(err, "Failed to publish a message")
		}

		log.Println("All messages sent!!!!")

	*/
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
