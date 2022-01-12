package query

import (
	"context"
	"go-service/database"
	"go-service/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var blogEntriesCollection = database.MongoDB().Database("devDB").Collection("blogEntries") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateBlogEntry(blogEntry entity.BlogEntry) *mongo.InsertOneResult {
	insertResult, err := blogEntriesCollection.InsertOne(context.TODO(), blogEntry)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}

func GetByID(id int) primitive.M {
	var result primitive.M
	err := blogEntriesCollection.FindOne(context.TODO(), bson.D{{Key: "ID", Value: id}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func UpdateBlogEntry(blogEntry entity.BlogEntry) primitive.M {

	filter := bson.D{{Key: "ID", Value: blogEntry.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "title", Value: blogEntry.Title}, {Key: "content", Value: blogEntry.Content}, {Key: "createddate", Value: blogEntry.CreatedDate}}}}
	updateResult := blogEntriesCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)
	var result primitive.M
	_ = updateResult.Decode(&result)
	return result
}

func DeleteBlogEntry(id int) *mongo.DeleteResult {
	opts := options.Delete().SetCollation(&options.Collation{})
	res, err := blogEntriesCollection.DeleteOne(context.TODO(), bson.D{{Key: "ID", Value: id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func GetAll() []primitive.M {
	var results []primitive.M
	cur, err := blogEntriesCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem) // appending document pointed by Next()
	}
	return results
}
