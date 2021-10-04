package databases

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	"github.com/niravparikh05/category-svcs/category"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	colon    string = ":"
	at       string = "@"
	database string = "finman-db"
)

type Mongodb struct {
	Host       string `properties:"mongodb.host,default=localhost"`
	Port       int    `properties:"mongodb.port,default=27017"`
	Username   string `properties:"mongodb.username"`
	Password   string `properties:"mongodb.password"`
	connstring string
}

/*
* ReadDatabaseProps function is used to read the properties file from filepath
* Parameters: filepath of type string
* Returns: Reference to Mongodb instance which has database properties
 */
func ReadDatabaseProps(filepath string) (*Mongodb, error) {
	prop := properties.MustLoadFile(filepath, properties.UTF8)
	mongodb := &Mongodb{}
	err := prop.Decode(mongodb)
	decodedUsername, err := b64.StdEncoding.DecodeString(mongodb.Username)
	mongodb.Username = string(decodedUsername)
	decodedPassword, err := b64.StdEncoding.DecodeString(mongodb.Password)
	mongodb.Password = string(decodedPassword)
	mongodb.connstring = mongodb.connectionString()
	return mongodb, err
}

func (mongodb *Mongodb) connectionString() string {
	return "mongodb://" + mongodb.Username + colon + mongodb.Password + at + mongodb.Host + colon + fmt.Sprint(mongodb.Port)
}

func (mongodb *Mongodb) findOne(dbCollection string, id string) (result *mongo.SingleResult) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := mongodb.connect(ctx)
	defer mongodb.disconnect(ctx, client)
	collection := client.Database(database).Collection(dbCollection)
	qryCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	return collection.FindOne(qryCtx, filter)
}

func (mongodb *Mongodb) findAll(dbCollection string) (results *mongo.Cursor, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := mongodb.connect(ctx)
	defer mongodb.disconnect(ctx, client)
	collection := client.Database(database).Collection(dbCollection)
	return collection.Find(ctx, bson.D{})
}

func (mongodb *Mongodb) insertOne(dbCollection string, document interface{}) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := mongodb.connect(ctx)
	defer mongodb.disconnect(ctx, client)
	collection := client.Database(database).Collection(dbCollection)
	res, err := collection.InsertOne(ctx, document)
	if err != nil {
		log.Fatalln("Error while inserting document ", document, err.Error())
		panic(err)
	}
	return res.InsertedID
}

func (mongodb *Mongodb) deleteOne(dbCollection string, id string) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := mongodb.connect(ctx)
	defer mongodb.disconnect(ctx, client)
	collection := client.Database(database).Collection(dbCollection)
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalln("Error while deleting data for id ", id, err.Error())
		panic(err)
	}
	return res.DeletedCount
}

func (mongodb *Mongodb) updateOne(dbCollection string, id string, document interface{}) (result *mongo.UpdateResult) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := mongodb.connect(ctx)
	defer mongodb.disconnect(ctx, client)
	collection := client.Database(database).Collection(dbCollection)
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	result, err := collection.ReplaceOne(ctx, filter, document)
	if err != nil {
		log.Fatalln("Error while updating data for id ", id, err.Error())
		panic(err)
	}
	return result
}

func (mongodb *Mongodb) connect(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb.connstring))
	if err != nil {
		fmt.Printf("Error connecting database %s\n", err.Error())
		panic(err)
	}
	return client
}

func (mongodb *Mongodb) disconnect(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		fmt.Printf("Error disconnecting database %s\n", err.Error())
		panic(err)
	}
	fmt.Println("Successfully disconnected database.")
}

func (mongodb *Mongodb) GetCategoryById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if len(strings.TrimSpace(id)) == 0 {
		http.Error(writer, "missing / invalid document id", http.StatusBadRequest)
		return
	}
	log.Println("Fetching category for id ", id)
	result := mongodb.findOne("category", id)
	categoryObj := &category.Category{}
	err := result.Decode(categoryObj)
	if err != nil {
		log.Printf("Error fetching document from database %s\n", err.Error())
		http.Error(writer, "Document not found, probably due to invalid input ", http.StatusBadRequest)
	} else {
		json.NewEncoder(writer).Encode(categoryObj)
	}
}

func (mongodb *Mongodb) GetAllCategories(writer http.ResponseWriter, request *http.Request) {
	results, err := mongodb.findAll("category")
	if err != nil {
		log.Printf("Error fetching data %s\n", err.Error())
		panic(err)
	}
	log.Println("Fetching all categories ")
	categories := []category.Category{}
	ctx := context.Background()
	for results.Next(ctx) {
		category := &category.Category{}
		results.Decode(category)
		categories = append(categories, *category)
	}
	json.NewEncoder(writer).Encode(categories)
}

func (mongodb *Mongodb) CreateCategory(writer http.ResponseWriter, request *http.Request) {
	category := category.Category{}
	json.NewDecoder(request.Body).Decode(&category)
	id := mongodb.insertOne("category", category)
	log.Println("Successfully created new category, id is ", id)
	writer.Write([]byte("Successfully Created !"))
}

func (mongodb *Mongodb) UpdateCategory(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if len(strings.TrimSpace(id)) == 0 {
		http.Error(writer, "missing / invalid document id", http.StatusBadRequest)
		return
	}
	category := category.Category{}
	json.NewDecoder(request.Body).Decode(&category)
	updResult := mongodb.updateOne("category", id, category)
	log.Println("Successfully updated category for id ", id, " impacted documents: ", updResult.ModifiedCount)
	writer.Write([]byte("Successfully Updated !"))
}

func (mongodb *Mongodb) DeleteCategory(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if len(strings.TrimSpace(id)) == 0 {
		http.Error(writer, "missing / invalid document id", http.StatusBadRequest)
		return
	}
	cnt := mongodb.deleteOne("category", id)
	log.Println("Successfully deleted document with id ", id, " documents impacted ", cnt)
	writer.Write([]byte("Successfully Deleted !"))
}
