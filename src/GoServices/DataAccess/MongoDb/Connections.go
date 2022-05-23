package MongoDb

import(
	//Local Packages
	"context"
	//This Project Packages
	myLogger "bulkmail/packages/Utils/Logger"
	//Online Packages
	//bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database
var collection *mongo.Collection
var ctx=context.TODO()
var err error

func GetClient(databaseName string, collectionName string){
	connectionOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/")
	client,err = mongo.Connect(ctx, connectionOptions)
	if err != nil {
		myLogger.FailOnError(err, "MONGO CONNECTION FAILED")
	}
	collection = client.Database(databaseName).Collection(collectionName)
}

func InsertOne(data string){
	_, er := collection.InsertOne(ctx, data)
	if er != nil {
		myLogger.FailOnError(err, "INSERT ONE FAILED")
	}
	myLogger.Print("inserted")
}

// func GetConnection() **mongo.Client{
// 	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err!= nil{
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err!= nil{
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	return &client
// }

// func GetDatabase(clnt *mongo.Client) *mongo.Database{
// 	return clnt.Database("test") 
// }

// func GetCollection(db *mongo.Database) *mongo.Collection{
// 	return db.Collection("testcollect")
// }