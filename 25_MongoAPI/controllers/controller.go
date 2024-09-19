package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sahil/MongoAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://<username>:<password>@<clustername>.<clusterlocation>.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client option
	clientOption := options.Client().ApplyURI(coonnectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)  // when you are making request to diff. machine you need to pass context 
	if err != nil {
		lof.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ready")
}

//MONGODB Helpers - file

// insert 1 record
func inserOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted One record in mongo db with id: ", inserted.InsertedID)
}

// update 1 movie
func updateOneMovie(movieId string) {
	id, _ := primitve.ObjectIDFromHex(movieId)
	filter := bson.D{"_id", id}
	update := bson.M{"$set", bson.M{"watched", true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Updated one movie with id: ", result.ModifiedCount)
}

// delete 1 movie
func deleteOneMovie() {
	id, _ := primitve.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: ", deleteCount)
}

// delete all records
func deleteAllMovie() int64{
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database
func getAllMovies() []primitive.M{
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}
	var movies []primitive.M
	
	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// Actual Controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// Create Movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// Mark a movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete a movie
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}

