package Reporsitory

import(
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"

	model "bulkmail/packages/Data/Models"
)

var ctx=context.TODO()

func FindByFilter(collection *mongo.Collection, filter bson.D)(*mongo.Cursor, model.ResultModel){
	var result model.ResultModel
	
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		result.StatusCode = 500
		result.Message = "FindByFilter-Find failed"
		result.Status = false
		return cursor, result
	}
	result.Status=true
	return cursor, result
}

func Insert(collection *mongo.Collection, data any) model.ResultModel{
	var result model.ResultModel
	
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		result.StatusCode = 500
		result.Message = "InsertOne failed"
		result.Status = false
		return result
	}
	result.Status = true
	return result
}

func Update(collection *mongo.Collection, id bson.M, data bson.M) model.ResultModel{
	var result model.ResultModel
	
	response, err := collection.UpdateMany(ctx, id, data)
	if err != nil {
		result.Status = false
		result.StatusCode = 500
		result.Message = "Update-UpdateOne failed"
		return result
	}
	if response.ModifiedCount == 0 {
		result.Status = false
		result.StatusCode = 404
		result.Message = "Update-UpdateOne not found"
		return result
	}
	result.Status = true
	return result
}

func Delete(collection *mongo.Collection, filter bson.M) model.ResultModel{
	var result model.ResultModel

	response, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		result.Status = false
		result.StatusCode = 500
		result.Message = "Delete failed"
		return result
	}
	if response.DeletedCount == 0 {
		result.Status = false
		result.StatusCode = 404
		result.Message = "Delete item not found"
		return result
	}
	result.Status = true
	return result
}