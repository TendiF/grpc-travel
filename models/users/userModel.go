package userModel

import (
	"deall-package/types"
	"deall-package/utils/database"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = "user"

func Insert(user types.User) (*mongo.InsertOneResult, error) {
	mongDB := database.MongoDB
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	return mongDB.Collection(collection).InsertOne(database.MongoContext, user)
}

func Update(ID string, user types.User) (*mongo.UpdateResult, error) {
	mongDB := database.MongoDB
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	fmt.Println(ID)
	id, _ := primitive.ObjectIDFromHex(ID)
	return mongDB.Collection(collection).UpdateByID(database.MongoContext, id, bson.D{
		{"$set", user},
	})
}

func Delete(ID string) (*mongo.DeleteResult, error) {
	mongDB := database.MongoDB
	id, _ := primitive.ObjectIDFromHex(ID)
	return mongDB.Collection(collection).DeleteOne(database.MongoContext, bson.D{
		{"_id", id},
	})
}

func FindOne(params bson.M) types.User {
	mongoDB := database.MongoDB
	var user types.User

	if err := mongoDB.Collection(collection).FindOne(database.MongoContext, params).Decode(&user); err != nil {
		fmt.Println("err", err)
	}

	return user
}

func Find(params bson.M, limit int64, skip int64) []types.User {
	mongDB := database.MongoDB
	options := options.FindOptions{}
	options.Limit = &limit
	options.Skip = &skip
	crsr, err := mongDB.Collection(collection).Find(database.MongoContext, params, &options)
	if err != nil {
		fmt.Println(err)
	}
	var user []types.User
	for crsr.Next(database.MongoContext) {
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
	mongDB := database.MongoDB
	total, err := mongDB.Collection(collection).CountDocuments(database.MongoContext, params)
	if err != nil {
		fmt.Println(err)
	}

	return total
}

func FindByUsername(username string) types.User {
	mongoDB := database.MongoDB
	var user types.User

	if err := mongoDB.Collection(collection).FindOne(database.MongoContext, bson.M{
		"username": username,
	}).Decode(&user); err != nil {
		fmt.Println("err", err)
	}

	return user
}

func FindById(id string) types.User {
	mongoDB := database.MongoDB
	var customer types.User
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("objectId err", err)
	}

	if err := mongoDB.Collection(collection).FindOne(database.MongoContext, bson.M{
		"_id": objectId,
	}).Decode(&customer); err != nil {
		fmt.Println("err", err)
	}

	return customer
}
