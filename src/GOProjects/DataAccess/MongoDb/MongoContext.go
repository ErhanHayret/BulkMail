package MongoDb

import(
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//primitive "go.mongodb.org/mongo-driver/bson/primitive"

	model "bulkmail/packages/Data/Models"
)

var client *mongo.Client
var database *mongo.Database
var ctx=context.TODO()
var err error
var result model.StatusResult

func GetClient(databaseName string, collectionName string) (*mongo.Collection, model.StatusResult){
	connectionOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client,err = mongo.Connect(ctx, connectionOptions)
	if err != nil {
		result.Error = err
		result.Message = "Mongo connection failed"
		result.Status = false
		return nil, result
	}
	
	result.Error = err
	result.Message = "Success"
	result.Status = true
	collection := client.Database(databaseName).Collection(collectionName)
	return collection, result
}

func InsertOne(collection *mongo.Collection, data any) model.StatusResult{
	_, er := collection.InsertOne(ctx, data)
	if er != nil {
		result.Error = er
		result.Message = "Mongo - InsertOne = InsetOne Failed"
		result.Status = false
		return result
	}

	result.Error = nil
	result.Message = "Success"
	result.Status = true
	return result
}

func FindAll(collection *mongo.Collection) ([]*model.Mail, model.StatusResult){
	var mails []*model.Mail

	filter := bson.D{{}}
	cur, er := collection.Find(ctx, filter)
	if er != nil {
		result.Error = er
		result.Message = "Mongo - FindAll = Find Failed"
		result.Status = false
		return mails, result
	}

	for cur.Next(ctx) {
		var mail model.Mail
		er := cur.Decode(&mail)
		if er != nil {
			result.Error = er
			result.Message = "Mongo - FindAll = Decode Failed"
			result.Status = false
			return mails, result
		}
		mails = append(mails, &mail)
	}

	if len(mails) == 0 {
		result.Error = errors.New("Database empty")
		result.Message = "Mongo - FindAll = DB Empty"
		result.Status = false
		return mails, result
	}

	result.Error = nil
	result.Message = "Success"
	result.Status = true
	return mails, result
}

func FindById(collection *mongo.Collection, id any) (model.Mail, model.StatusResult){
	var mail model.Mail

	filter := bson.D{{"id",id}}
	proj := bson.D{{"id",1}}
	opts := options.FindOne().SetProjection(proj)
	er := collection.FindOne(ctx, filter, opts).Decode(&mail)
	if er != nil {
		result.Error = er
		result.Message = "Mongo - FindById = FindOne Failed"
		result.Status = false
		return mail, result
	}

	result.Error = nil
	result.Message = "Success"
	result.Status = true
	return mail, result
}

func FindBySender(collection *mongo.Collection, adress string) ([]*model.Mail, model.StatusResult){
	var mails []*model.Mail

	filter := bson.M{"senderemail":adress}
	cur, er := collection.Find(ctx, filter)
	if er != nil {
		result.Error = er
		result.Message = "Mongo - FindBySender = Find Failed"
		result.Status = false
		return mails, result
	}

	for cur.Next(ctx) {
		var mail model.Mail
		er := cur.Decode(&mail)
		if er != nil {
			result.Error = er
			result.Message = "Mongo - FindBySender = Decode Failed"
			result.Status = false
			return mails, result
		}
		mails = append(mails, &mail)
	}

	if len(mails) == 0 {
		result.Error = errors.New("Database empty")
		result.Message = "Mongo - FindBySender = DB Empty"
		result.Status = false
		return mails, result
	}

	result.Error = nil
	result.Message = "Success"
	result.Status = true
	return mails, result
}