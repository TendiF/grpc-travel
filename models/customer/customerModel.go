package customerModel

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

var collection = "customer"

func Insert(customer types.Customer) (*mongo.InsertOneResult, error) {
	mongDB := database.MongoDB

	customer.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	return mongDB.Collection(collection).InsertOne(database.MongoContext, customer)
}

func InsertMany(customers []interface{}) (*mongo.InsertManyResult, error) {
	mongDB := database.MongoDB

	return mongDB.Collection(collection).InsertMany(database.MongoContext, customers)
}

func Count(params bson.M) int64 {
	mongDB := database.MongoDB

	number, error := mongDB.Collection(collection).CountDocuments(database.MongoContext, params)
	if error != nil {
		fmt.Println("[err count]", error)
	}
	return number
}

func Find(params bson.M, limit int64, skip int64) []types.Customer {
	mongDB := database.MongoDB

	options := options.FindOptions{}
	if limit > 0 {
		options.Limit = &limit
	}

	if skip > 0 {
		options.Skip = &skip
	}

	crsr, err := mongDB.Collection(collection).Find(database.MongoContext, params, &options)
	if err != nil {
		fmt.Println(err)
	}
	var Customer []types.Customer
	for crsr.Next(database.MongoContext) {
		var row types.Customer
		err := crsr.Decode(&row)
		if err != nil {
			fmt.Println("err Decode", err)
		}
		Customer = append(Customer, row)
	}

	return Customer
}

func FindById(id string) types.Customer {
	mongoDB := database.MongoDB
	var customer types.Customer

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

func Update(id string, params types.Customer) types.Customer {
	mongoDB := database.MongoDB
	var customer types.Customer

	params.ModifiedAt = primitive.NewDateTimeFromTime(time.Now())

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("objectId err", err)
	}

	fmt.Println(objectId, params)

	mongoDB.Collection(collection).FindOneAndReplace(database.MongoContext, bson.M{
		"_id": objectId,
	}, params).Decode(&customer)

	return customer
}

func CountDocuments(params bson.M) int64 {
	mongDB := database.MongoDB
	total, err := mongDB.Collection(collection).CountDocuments(database.MongoContext, params)
	if err != nil {
		fmt.Println(err)
	}

	return total
}
