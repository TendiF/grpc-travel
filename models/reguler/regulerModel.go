package regulerModel

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

var collection = "refuler"

func Insert(reguler types.Reguler) (*mongo.InsertOneResult, error) {
	mongDB := database.MongoDB
	reguler.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	return mongDB.Collection(collection).InsertOne(database.MongoContext, reguler)
}

func Update(ID string, reguler types.Reguler) (*mongo.UpdateResult, error) {
	mongDB := database.MongoDB
	reguler.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	fmt.Println(ID)
	id, _ := primitive.ObjectIDFromHex(ID)
	return mongDB.Collection(collection).UpdateByID(database.MongoContext, id, bson.D{
		{"$set", reguler},
	})
}

func Delete(ID string) (*mongo.DeleteResult, error) {
	mongDB := database.MongoDB
	id, _ := primitive.ObjectIDFromHex(ID)
	return mongDB.Collection(collection).DeleteOne(database.MongoContext, bson.D{
		{"_id", id},
	})
}

func FindOne(params bson.M) types.Reguler {
	mongoDB := database.MongoDB
	var reguler types.Reguler

	if err := mongoDB.Collection(collection).FindOne(database.MongoContext, params).Decode(&reguler); err != nil {
		fmt.Println("err", err)
	}

	return reguler
}

func Find(params bson.M, limit int64, skip int64) []types.Reguler {
	mongDB := database.MongoDB
	options := options.FindOptions{}
	options.Limit = &limit
	options.Skip = &skip
	crsr, err := mongDB.Collection(collection).Find(database.MongoContext, params, &options)
	if err != nil {
		fmt.Println(err)
	}
	var reguler []types.Reguler
	for crsr.Next(database.MongoContext) {
		var row types.Reguler
		err := crsr.Decode(&row)
		if err != nil {
			fmt.Println("err Decode", err)
		}
		reguler = append(reguler, row)
	}

	return reguler
}

func CountDocuments(params bson.M) int64 {
	mongDB := database.MongoDB
	total, err := mongDB.Collection(collection).CountDocuments(database.MongoContext, params)
	if err != nil {
		fmt.Println(err)
	}

	return total
}

func FindById(id string) types.Reguler {
	mongoDB := database.MongoDB
	var customer types.Reguler
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
