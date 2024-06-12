package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_STRING = "mongodb+srv://khushiastrogeek:khushicosmos9@mongodb-go.rtbazij.mongodb.net/?retryWrites=true&w=majority&appName=mongodb-go"

const DBNAME = "Netflix"

const COLLECTION_NAME = "watchlist"

// Most Important
var collection *mongo.Collection


// connect with mongoDb
// initialization method, only runs for one time and at the very start, special method in golang
func init() {

	// client option
	clientOption := options.Client().ApplyURI(CONNECTION_STRING)

	// connecting to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDb connection success.")

	collection = client.Database(DBNAME).Collection(COLLECTION_NAME)

	// collection instance
	fmt.Println("Collection reference/instance is ready.")

}



// helper methods of mongoDb - should be in another file

// insert one record
// lowercase functions as we are not exporting them
func insertOneMovie(movie model.Netflix) {			// model name from package name

	// you have to always pass the context, whenever you are doing any database operations
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted One Movie in DB with id : ", inserted.InsertedID)

}



// update one record
func updateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)			// converts string into the valid ObjectID

	// bson.M = Not case sensitive 
	// bson.D = for ordered elements
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched" : true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count : ", result.ModifiedCount)

}



// deleting a record
func deleteOneMovie(movieId string)  {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id" : id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count : ", deleteCount)

}



// delete all records from mongoDb
func deleteAllMovies() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)		// nil for options

	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println("number of movies deleted : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
	
}



// get all movies from db
func getAllMovies() []primitive.M {

	// cursor is a big gigantic object on which you can loop through and get the data.
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	// the loop will keep on giving you values as long as it has values
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




// actual controllers - should be in this file only

func GetAllMyMovies(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
	
}



func CreateMovie(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	
	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)

}



func MarkAsWatched(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
	
}



func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}



func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode(count)

}