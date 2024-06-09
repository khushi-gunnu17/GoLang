package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_STRING = "mongodb+srv://khushiastrogeek:s2otFs8xQueOcSc0@mongodb-go.rtbazij.mongodb.net/?retryWrites=true&w=majority&appName=mongodb-go"

const DBNAME = "Netflix"

const COLLECTION_NAME = "watchlist"

// Most Important
var collection *mongo.Collection


// connect with mongoDb
// initialization method, only runs for one time and at the very start
func init() {

	// client option
	clientOption := options.Client().ApplyURI(CONNECTION_STRING)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDb connection success.")

	collection = client.Database(DBNAME).Collection(COLLECTION_NAME)

	// collection instance
	fmt.Println("Collection reference is ready.")

}



// helper methods of mongoDb - file

// insert one record
// model from package name
func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted One Movie in DB with id : ", inserted.InsertedID)

}


// update one record
func updateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched" : true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count : ", result.ModifiedCount)

}


// deleting a record
func deleteOneRecord(movieId string)  {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id" : id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count : ", deleteCount)

}


// delete all records fro mongoDb
func deleteAllRecords() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println("number of movies deleted : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
	
}


// get all movies from db
func getAllMovies() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M

		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
	
}


// actual controllers - this file

func GetAllMyMovies(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
	
}


func createMovie(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	
	
	
}