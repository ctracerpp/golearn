package mongodb

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri_mongo = "mongodb://localhost:27017/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&3t.uriVersion=3&3t.connection.name=localhost-mongo-27017&3t.alwaysShowAuthDB=true&3t.alwaysShowDBFromUserRole=true"

const uri = "mongodb://127.0.0.1:27017/?maxPoolSize=20"

func TestName(t *testing.T) {
	fmt.Println("test mongodb ")
	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

// Language: go

/*
   Define my document struct
*/
type Post struct {
	// Id    int64  `bson:"_id"`
	Guid  string `bson:"guid,omitempty"`
	Level int32  `bson:"level,omitempty"`
}

func TestMongoDbOrm(t *testing.T) {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*60)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   Get my collection instance
	*/
	collection := client.Database("cloud_dw").Collection("ods_oss_divis_quantity_list_record")

	/*
	   Insert documents
	*/
	//docs := []interface{}{
	//	bson.D{{"_id", 4}, {"title", "World"}, {"body", "Hello World"}},
	//	bson.D{{"_id", 5}, {"title", "Mars"}, {"body", "Hello Mars"}},
	//	bson.D{{"_id", 9123456789012345678}, {"title", "Pluto"}, {"body", "Hello Pluto"}},
	//}
	//
	//res, insertErr := collection.InsertMany(ctx, docs)
	//if insertErr != nil {
	//	log.Fatal(insertErr)
	//}
	//fmt.Println(res)
	/*
	   Iterate a cursor and print it
	*/
	cur, currErr := collection.Find(ctx, bson.D{}, options.Find().SetLimit(10000000).SetBatchSize(1000))

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	//var posts []Post
	for cur.Next(ctx) {
		// var post Post
		// err := cur.Decode(&post)
		//if err != nil {
		//	panic(err)
		//}
		// fmt.Println(post)
		//posts = append(posts, post)
	}
	//if err = cur.All(ctx, &posts); err != nil {
	//	panic(err)
	//}
	//fmt.Println(len(posts))
}

// 创建10W个数据库及对应的测试collection
func TestCreat100kDB(t *testing.T) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*60)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	for i := 0; i < 100000; i++ {
		// client.Database(fmt.Sprintf("test_%d", i)).CreateCollection(ctx, "test")
		client.Database(fmt.Sprintf("test_%d", i)).Drop(ctx)
	}
	// output: 566.38s
}
