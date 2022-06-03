package MongoDb

import(
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	model "bulkmail/packages/Data/Models"
	dto "bulkmail/packages/Data/Dtos"
	repo "bulkmail/packages/DataAccess/Reporsitory"

)

var client *mongo.Client
var database *mongo.Database
var ctx=context.TODO()

func GetClient(databaseName string, collectionName string) (*mongo.Collection, model.ResultModel){
	var err error
	var result model.ResultModel

	connectionOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client,err = mongo.Connect(ctx, connectionOptions)
	if err != nil {
		result.StatusCode = 500
		result.Message = "Mongo connection failed"
		result.Status = false
		return nil, result
	}
	result.Status = true
	collection := client.Database(databaseName).Collection(collectionName)
	return collection, result
}

//Mail Methods

func FindAllMails(collection *mongo.Collection) ([]*model.MailModel, model.ResultModel){
	var result model.ResultModel
	var mails []*model.MailModel

	filter := bson.D{{}}
	cur, dbResponse := repo.FindByFilter(collection, filter)
	if dbResponse.Status == true {
		for cur.Next(ctx) {
			var mail model.MailModel
			err := cur.Decode(&mail)
			if err != nil {
				result.StatusCode = 500
				result.Message = "FindAllMails-Decode failed"
				result.Status = false
				return mails, result
			}
			mails = append(mails, &mail)
		}
	}
	if len(mails) == 0 {
		result.StatusCode = 404
		result.Message = "Db Empty"
		result.Status = false
		return mails, result
	}
	result.Status = true
	return mails, result
}

func FindMailById(collection *mongo.Collection, id string) (model.MailModel, model.ResultModel){
	var result model.ResultModel
	var mail model.MailModel

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"id", objId}}
	cur, dbResponse := repo.FindByFilter(collection, filter)
	if dbResponse.Status == true {
		for cur.Next(ctx){
			er := cur.Decode(&mail)
			if er != nil {
				result.StatusCode = 500
				result.Message = "FindMailById-Decode failed"
				result.Status = false
				return mail, result
			}
			result.Status = true
			return mail, result
		}
	}
	return mail, dbResponse
}

func FindBySender(collection *mongo.Collection, adress string) ([]*model.MailModel, model.ResultModel){
	var result model.ResultModel
	var mails []*model.MailModel

	filter := bson.D{{"senderemail",adress}}
	cur, dbResponse := repo.FindByFilter(collection, filter)
	if dbResponse.Status == true {
		for cur.Next(ctx) {
			var mail model.MailModel
			err := cur.Decode(&mail)
			if err != nil {
				result.StatusCode = 500
				result.Message = "FindBySender-Decode failed"
				result.Status = false
				return mails, result
			}
			mails = append(mails, &mail)
		}
	}
	if len(mails) == 0 {
		result.StatusCode = 404
		result.Message = "Sender or any row not found"
		result.Status = false
		return mails, result
	}
	result.Status = true
	return mails, result
}

func InsretMail(collection *mongo.Collection, mail model.MailModel) model.ResultModel{
	mail.Id = primitive.NewObjectID().Hex()
	result := repo.Insert(collection, mail)
	return result
}

//User Methods

func FindUser(collection *mongo.Collection, userDto dto.UserDto) (model.UserModel, model.ResultModel){
	var result model.ResultModel
	var user model.UserModel

	filter := bson.D{{"username", userDto.UserName}, {"password", userDto.Password}}
	cur, dbResponse := repo.FindByFilter(collection, filter)
	if dbResponse.Status == true {
		for cur.Next(ctx){
			er := cur.Decode(&user)
			if er != nil {
				result.Status = false
				result.Message = "FindUser-Decode failed"
				result.StatusCode = 500
				return user, result
			}
			result.Status = true
			return user, result
		}
	}
	return user, dbResponse
}

func FindUserById(collection *mongo.Collection, id string) (model.UserModel, model.ResultModel){
	var result model.ResultModel
	var user model.UserModel

	filter := bson.D{{"id", id}}
	cur, dbResponse := repo.FindByFilter(collection, filter)
	if dbResponse.Status == true {
		for cur.Next(ctx){
			er := cur.Decode(&user)
			if er != nil {
				result.StatusCode = 500
				result.Message = "FindUserById-Decode failed"
				result.Status = false
				return user, result
			}
			result.Status = true
			return user, result
		}
	}
	return user, dbResponse
}

func InsertUser(collection *mongo.Collection, user model.UserModel) model.ResultModel{
	user.Id = primitive.NewObjectID().Hex()
	result := repo.Insert(collection, user)
	return result
}

func UpdateUser(collection *mongo.Collection, user model.UserModel) model.ResultModel{
	bsonid := bson.M{"id": bson.M{"$eq": user.Id}}
	data := bson.M{"$set": bson.M{"username": user.UserName,"password": user.Password, "isadmin": user.IsAdmin}}
	result := repo.Update(collection, bsonid, data)
	return result
}

func DeleteUser(collection *mongo.Collection, id string) model.ResultModel{
	bsonid := bson.M{"id": id}
	result := repo.Delete(collection, bsonid)
	return result
}