package userModel

import (
	"deall-package/types"
	"deall-package/utils"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = "user"

func Insert(user types.User) (*mongo.InsertOneResult, error) {
	mongDB := utils.MongoDB
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	return mongDB.Collection(collection).InsertOne(utils.MongoContext, user)
}

func Update(ID string, user types.User) (*mongo.UpdateResult, error) {
	mongDB := utils.MongoDB
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	fmt.Println(ID)
	id, _ := primitive.ObjectIDFromHex(ID)
	return mongDB.Collection(collection).UpdateByID(utils.MongoContext, id, bson.D{
		{"$set", user},
	})
}

func Delete(ID string) (*mongo.DeleteResult, error) {
	mongDB := utils.MongoDB
	id, _ := primitive.ObjectIDFromHex(ID)
	return mongDB.Collection(collection).DeleteOne(utils.MongoContext, bson.D{
		{"_id", id},
	})
}

func Find(params bson.M, limit int64, skip int64) []types.User {
	mongDB := utils.MongoDB
	options := options.FindOptions{}
	options.Limit = &limit
	options.Skip = &skip
	crsr, err := mongDB.Collection(collection).Find(utils.MongoContext, params, &options)
	if err != nil {
		fmt.Println(err)
	}
	var user []types.User
	for crsr.Next(utils.MongoContext) {
		var row types.User
		err := crsr.Decode(&row)
		if err != nil {
			fmt.Println("err Decode", err)
		}
		user = append(user, row)
	}

	return user
}

func CountDocuments(params bson.M) int64 {
	mongDB := utils.MongoDB
	total, err := mongDB.Collection(collection).CountDocuments(utils.MongoContext, params)
	if err != nil {
		fmt.Println(err)
	}

	return total
}

func FindByUsername(username string) types.User {
	mongoDB := utils.MongoDB
	var user types.User

	if err := mongoDB.Collection(collection).FindOne(utils.MongoContext, bson.M{
		"username": username,
	}).Decode(&user); err != nil {
		fmt.Println("err", err)
	}

	return user
}

func FindById(id string) types.User {
	mongoDB := utils.MongoDB
	var customer types.User
	objectId, err := primitive.ObjectIDFromHex(id)

	fmt.Println("id", id, objectId, collection)
	if err != nil {
		fmt.Println("objectId err", err)
	}

	if err := mongoDB.Collection(collection).FindOne(utils.MongoContext, bson.M{
		"_id": objectId,
	}).Decode(&customer); err != nil {
		fmt.Println("err", err)
	}

	return customer
}
