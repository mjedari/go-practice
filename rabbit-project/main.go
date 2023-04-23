package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	err = ch.ExchangeDeclare("test-exchange", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	firstQueue, err := ch.QueueDeclare("first-queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare firstQueue: %v", err)
	}

	secondQueue, err := ch.QueueDeclare("second-queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare firstQueue: %v", err)
	}

	err = ch.QueueBind(firstQueue.Name, "", "test-exchange", false, nil)
	if err != nil {
		log.Fatalf("Failed to bind firstQueue: %v", err)
	}
	err = ch.QueueBind(secondQueue.Name, "", "test-exchange", false, nil)
	if err != nil {
		log.Fatalf("Failed to bind firstQueue: %v", err)
	}

	// run publisher here
	go publisher(ch)

	// run consumer here
	consumer(ch, firstQueue)

}

func consumer(ch *amqp.Channel, queue amqp.Queue) {
	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	for msg := range msgs {
		log.Println("received message:", string(msg.Body))
	}
}

func publisher(ch *amqp.Channel) {

	for {
		select {
		case <-time.Tick(time.Second * 5):
			sendQueue(ch)
		}
	}
}

func sendQueue(ch *amqp.Channel) {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello, World!"),
	}
	err := ch.Publish("test-exchange", "", false, false, message)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
}
