package RabbitMQ

import(
	//This Project Packages
	myLog "bulkmail/packages/Utils/Logger"

	//Git Packages
	amqp "github.com/rabbitmq/amqp091-go"
)

func AddToQueue(body []byte){
	//RabbitMq connection
	conn, err := amqp.Dial("amqp://root:root@localhost:5672")
	myLog.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	//RabbitMq Channell
	ch, err := conn.Channel()
	myLog.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	//RabbitMq queue declare
	q, err := ch.QueueDeclare(
		"BulkMail",	//name
		false, 		//durable
		false, 		//delete when unused
		false, 		//exclusive
		false, 		//no-wait
		nil, 		//arguments
	)
	myLog.FailOnError(err, "Failed to declare a queue")
	//Publish data
	er := ch.Publish(
		"", 	//exchange
		q.Name, //reouting key
		false, 	//mandatory
		false, 	//immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body: body,
		})
	myLog.FailOnError(er, "Failed to publish a message")
	myLog.PrintData("Sended this data => ", body)//Logger
}