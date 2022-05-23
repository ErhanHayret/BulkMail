package main

import(
	//Local Packages
	"log"
	//This Project Packages
	db "bulkmail/packages/DataAccess/MongoDb"
	"bulkmail/packages/Data/Models"
	//Git Packages
	amqp "github.com/rabbitmq/amqp091-go"
	//"go.mongodb.org/mongo-driver/mongo"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	//Rabbit Connection
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	//Rabbit Channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	//Rabbit Queue
	q, err := ch.QueueDeclare(
		"BulkMail", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	var mail Models.Mail
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	
	//var clnt *mongo.Client
	clnt:= db.GetConnection()
	mongodb:=db.CreateDatabase(*clnt)
	col:=db.CreateCollection(*mongodb)
	

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}