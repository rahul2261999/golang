package controller

import (
	"25mongodb/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectString string = "mongodb+srv://rahulsaini:Password1@mongo-cluster.v7o8ni7.mongodb.net/?retryWrites=true&w=majority"
const dbName string = "netflix"
const collextionName string = "watchlist"

var collection *mongo.Collection

func InitDb() {
	clientOption := options.Client().ApplyURI(connectString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected!")

	collection = client.Database(dbName).Collection(collextionName)

	fmt.Println("Collection 'watchlist' initialized!")
}

func insertMovie(movie model.Netflix) {
	data, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted movie with ID: ", data.InsertedID)
}

func updateMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	println("Updated movie with ID: ", movieId, " with result: ", result.MatchedCount)
}

func deleteMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted movie with ID: ", movieId, " with result: ", result.DeletedCount)
}

func deleteAllMovies() {
	result, err := collection.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all movies with result: ", result.DeletedCount)
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})

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

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)

}

func InsertMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methdss", "POST")

	var movie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(movie)

	insertMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methdss", "PUT")

	params := mux.Vars(r)

	updateMovie(params["id"])

	json.NewEncoder(w).Encode("Movie updated successfully!")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methdss", "DELETE")

	params := mux.Vars(r)

	deleteMovie(params["id"])

	json.NewEncoder(w).Encode("Movie deleted successfully!")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methdss", "DELETE")

	deleteAllMovies()

	json.NewEncoder(w).Encode("All movies deleted successfully!")
}
