package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://rashisharma:CaWvYi0kg7lNjc3D@cluster0.ebgvgj3.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbname = "Netflixdb"
const colName = "watchlist"

var collection *mongo.Collection

// connect woth mongodb
// special method which run only one time
func init() {
	//client connection
	clientoption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		panic(err)
	}
	collection = client.Database(dbname).Collection(colName)
	fmt.Println("db connected ")
}

func insertOnemovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted in db ", inserted.InsertedID)
}

// updateOnemovie
func updateonemovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified ", result.ModifiedCount)

}

// delete one
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete successdully ", deletecount)

}

// delete all
func deleteallMovie() int64 {
	// filter := bson.M{}
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete all", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies
func getallmovies() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer curr.Close(context.Background())
	return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencod")
	allmovies := getallmovies()
	json.NewEncoder(w).Encode(allmovies)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencod")
	w.Header().Set("Allow-control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	insertOnemovie(movie)
	fmt.Println("after insert")
	json.NewEncoder(w).Encode(movie)

}
func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencod")
	w.Header().Set("Allow-control-Allow-Methods", "PUT")
	params := chi.URLParam(r, "id")
	updateonemovie(params)
	json.NewEncoder(w).Encode(params)

}
func DeleteAmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencod")
	w.Header().Set("Allow-control-Allow-Methods", "DELETE")
	params := chi.URLParam(r, "id")
	deleteOneMovie(params)
	json.NewEncoder(w).Encode(params)
}
func DeleteAllmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencod")
	w.Header().Set("Allow-control-Allow-Methods", "DELETE")

	count := deleteallMovie()
	json.NewEncoder(w).Encode(count)
}
