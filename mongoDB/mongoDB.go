package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	err    error
)

func DemoMongoDB() {
	connect()
	testing()
	defer disconnect()
}

func testing() {
	dbns := showDatabaseNames()
	if len(dbns) > 0 {
		// find a database to use
		dbName := dbns[0]
		colName := "testing"
		createCollection(dbName, colName)
		showAllCollectionName(dbName)
		col := getCollection(dbName, colName)
		data := bson.M{"name": "demo", "index": 0}
		dataArr := []interface{}{
			bson.M{"name": "test1", "index": 1},
			bson.M{"name": "test2", "index": 2},
			bson.M{"name": "test3", "index": 3},
		}
		filter := bson.M{"name": "demo"}
		createData(col, data)
		createDatas(col, dataArr, false)
		getData(col, filter)
		deleteData(col, filter)
		deleteCollection(dbName, colName)
	}
}

func connect() {
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017/?connect=direct")
	client, err = mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
}

func disconnect() {
	client.Disconnect(context.TODO())
}

func showDatabaseNames() []string {
	dbns, _ := client.ListDatabaseNames(context.TODO(), bson.M{})
	fmt.Print("[showDatabaseNames] ")
	fmt.Println(dbns)
	return dbns
}

func showAllCollectionName(dbName string) []string {
	db := client.Database(dbName)
	cns, _ := db.ListCollectionNames(context.TODO(), bson.M{})
	fmt.Print("[showAllCollectionName] ")
	fmt.Println(cns)
	return cns
}

func createCollection(dbName, colName string) {
	client.Database(dbName).CreateCollection(context.TODO(), colName)
}

func getCollection(dbName, colName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}

func deleteCollection(dbName, colName string) {
	client.Database(dbName).Collection(colName).Drop(context.TODO())
}

func createData(col *mongo.Collection, data interface{}) {
	_, err := col.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("[createData] data:" + fmt.Sprint(data) + " happen err: " + err.Error())
	}
}

func createDatas(col *mongo.Collection, data []interface{}, insertFailedThenStop bool) {
	opt := options.InsertMany().SetOrdered(insertFailedThenStop)
	col.InsertMany(context.TODO(), data, opt)
}

func getData(col *mongo.Collection, filter interface{}) *mongo.SingleResult {
	ref := col.FindOne(context.TODO(), filter)
	data := &bson.M{}
	ref.Decode(data)
	fmt.Println("[getData] filter:" + fmt.Sprint(filter) + " ref: " + fmt.Sprint(data))
	return ref
}

func getDatas(col *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {
	ref, err := col.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("[getDatas] filter:" + fmt.Sprint(filter) + " happen err: " + err.Error())
		return nil, err
	}
	return ref, nil
}

func deleteData(col *mongo.Collection, filter interface{}) {
	col.DeleteOne(context.TODO(), filter)
}

func deleteDatas(col *mongo.Collection, filter interface{}) {
	col.DeleteMany(context.TODO(), filter)
}
