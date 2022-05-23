package Consumer

import(
	//Local Packages
	"log"
	"sync"

	//This Project Packages
	myMongo "bulkmail/packages/DataAccess/MongoDb"
	//"bulkmail/packages/Data/Models"
	myLogger "bulkmail/packages/Utils/Looger"

	//Git Packages
	amqp "github.com/rabbitmq/amqp091-go"
	//"go.mongodb.org/mongo-driver/mongo"
)

func Consumer(wg *sync.WaitGroup) {
	defer wg.Done()
	//Rabbit Connection
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	myLogger.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	//Rabbit Channel
	ch, err := conn.Channel()
	myLogger.FailOnError(err, "Failed to open a channel")
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
	myLogger.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	myLogger.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	
	go func() {
		for d := range msgs {
			myLogger.PrintData("Received a message => ", d.Body)

			myMongo.GetClient("test","testcol")
			myMongo.InsertOne("test")
		}
	}()

	myLogger.Print("Waiting for messages. To exit press CTRL+C")
	<-forever
}