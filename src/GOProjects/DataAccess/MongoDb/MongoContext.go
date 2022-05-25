package MongoDb

import(
	//Local Packages
	"context"

	//This Project Packages
	"bulkmail/packages/Data/Models"
	eLog "bulkmail/packages/Utils/Logger"
	
	//Online Packages
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

var client *mongo.Client
var database *mongo.Database
var ctx=context.TODO()
var err error

func GetClient(databaseName string, collectionName string) *mongo.Collection{
	connectionOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client,err = mongo.Connect(ctx, connectionOptions)
	if err != nil {
		eLog.FailOnError(err, "MONGO CONNECTION FAILED")
	}
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}

func InsertOne(collection *mongo.Collection, data any){
	//GetClient("MailDb","Mail")
	_, er := collection.InsertOne(ctx, data)
	if er != nil {
		eLog.FailOnError(er, "InsetOne Failed")
	} else {
		eLog.PrintData("Inserted to db:" , data)
	}
}

func FindAll(collection *mongo.Collection) []*Models.Mail{
	//GetClient("test", "test")
	var mails []*Models.Mail
	filter := bson.D{{}}
	cur, er := collection.Find(ctx, filter)
	if err != nil {
		eLog.FailOnError(er, "FindAll Failed")
		return mails
	}
	for cur.Next(ctx) {
		var mail Models.Mail
		er := cur.Decode(&mail)
		if er != nil {
			eLog.FailOnError(er, "FindAll-Decode Failed")
			return mails
		}
		mails = append(mails, &mail)
	}
	if len(mails) == 0 {
		eLog.ErrorPrint("DB Empty")
	}
	return mails
}

func FindById(collection *mongo.Collection, id any) Models.Mail{
	//GetClient("test", "test")
	var mail Models.Mail
	filter := bson.D{{"id",id}}
	proj := bson.D{{"id",1}}
	opts := options.FindOne().SetProjection(proj)
	er := collection.FindOne(ctx, filter, opts).Decode(&mail)
	if er != nil {
		eLog.FailOnError(er, "FindById Failed")
		return mail
	}
	return mail
}

func FindBySender(collection *mongo.Collection, adress string) []*Models.Mail{
	//GetClient("test", "test")
	var mails []*Models.Mail
	filter := bson.M{"senderemail":adress}
	cur, er := collection.Find(ctx, filter)
	if er != nil {
		eLog.FailOnError(er, "FindBySender Failed")
		return mails
	}
	for cur.Next(ctx) {
		var mail Models.Mail
		er := cur.Decode(&mail)
		if er != nil {
			eLog.FailOnError(er ,"FindBySender-Decode Failed")
			return mails
		}
		mails = append(mails, &mail)
	}
	if len(mails) == 0 {
		eLog.ErrorPrint("DB Empty")
	}
	return mails
}