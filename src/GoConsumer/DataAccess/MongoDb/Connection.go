package MongoDb

import(
	//Local Packages
	"context"
	"log"
	"time"
	//Online Packages
	//bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

func GetConnection() **mongo.Client{
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!= nil{
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err!= nil{
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	return &client
}

func CreateDatabase(clnt *mongo.Client) *mongo.Database{
	return clnt.Database("test") 
}

func CreateCollection(db *mongo.Database) *mongo.Collection{
	return db.Collection("testcollect")
}