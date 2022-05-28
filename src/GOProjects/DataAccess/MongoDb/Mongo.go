package MongoDb

import(
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//primitive "go.mongodb.org/mongo-driver/bson/primitive"

	model "bulkmail/packages/Data/Models"
	dto "bulkmail/packages/Data/Dtos"

)

var client *mongo.Client
var database *mongo.Database
var ctx=context.TODO()
var err error
var result model.ResultModel

func GetClient(databaseName string, collectionName string) (*mongo.Collection, model.ResultModel){
	connectionOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client,err = mongo.Connect(ctx, connectionOptions)
	if err != nil {
		result.StatusCode = 500
		result.Message = "Mongo connection failed"
		result.Status = false
		return nil, result
	}
	
	result.StatusCode = 200
	result.Message = "Success"
	result.Status = true
	collection := client.Database(databaseName).Collection(collectionName)
	return collection, result
}

func InsertOne(collection *mongo.Collection, data any) model.ResultModel{
	_, er := collection.InsertOne(ctx, data)
	if er != nil {
		result.StatusCode = 500
		result.Message = "InsertOne failed"
		result.Status = false
		return result
	}

	result.StatusCode = 200
	result.Message = "Success"
	result.Status = true
	return result
}

//Mail Methods

func FindAllMails(collection *mongo.Collection) ([]*model.MailModel, model.ResultModel){
	var mails []*model.MailModel

	filter := bson.D{{}}
	cur, er := collection.Find(ctx, filter)
	if er != nil {
		result.StatusCode = 500
		result.Message = "FindAllMail-Find failed"
		result.Status = false
		return mails, result
	}

	for cur.Next(ctx) {
		var mail model.MailModel
		er := cur.Decode(&mail)
		if er != nil {
			result.StatusCode = 500
			result.Message = "FindAllMail-Decode failed"
			result.Status = false
			return mails, result
		}
		mails = append(mails, &mail)
	}

	if len(mails) == 0 {
		result.StatusCode = 404
		result.Message = "Db Empty"
		result.Status = false
		return mails, result
	}

	result.StatusCode = 200
	result.Message = "Success"
	result.Status = true
	return mails, result
}

func FindMailById(collection *mongo.Collection, id any) (model.MailModel, model.ResultModel){
	var mail model.MailModel

	filter := bson.D{{"id",id}}
	proj := bson.D{{"id",1}}
	opts := options.FindOne().SetProjection(proj)
	er := collection.FindOne(ctx, filter, opts).Decode(&mail)
	if er != nil {
		result.StatusCode = 500
		result.Message = "FindMailById-FinOne failed"
		result.Status = false
		return mail, result
	}

	result.StatusCode = 200
	result.Message = "Success"
	result.Status = true
	return mail, result
}

func FindBySender(collection *mongo.Collection, adress string) ([]*model.MailModel, model.ResultModel){
	var mails []*model.MailModel

	filter := bson.M{"senderemail":adress}
	cur, er := collection.Find(ctx, filter)
	if er != nil {
		result.StatusCode = 500
		result.Message = "FindBySender-Find failed"
		result.Status = false
		return mails, result
	}

	for cur.Next(ctx) {
		var mail model.MailModel
		er := cur.Decode(&mail)
		if er != nil {
			result.StatusCode = 500
			result.Message = "FindBySender-Decode failed"
			result.Status = false
			return mails, result
		}
		mails = append(mails, &mail)
	}

	if len(mails) == 0 {
		result.StatusCode = 404
		result.Message = "Sender not found"
		result.Status = true
		return mails, result
	}

	result.StatusCode = 200
	result.Message = "Success"
	result.Status = true
	return mails, result
}

//User Methods

func FindUser(collection *mongo.Collection, userDto dto.UserDto) (model.UserModel, model.ResultModel){
	var user model.UserModel

	filter := bson.D{{"username",userDto.UserName},{"password",userDto.Password}}
	proj := bson.D{{"username",1},{"password",1}}
	opts := options.FindOne().SetProjection(proj)
	er := collection.FindOne(ctx, filter, opts).Decode(&user)
	if er != nil {
		result.StatusCode = 500
		result.Message = "FinUser-FinOne failed"
		result.Status = false
		return user, result
	}

	result.StatusCode = 200
	result.Message = "Success"
	result.Status = true
	return user, result
}