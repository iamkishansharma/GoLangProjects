package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iamkishansharma/GoLangProjects/tree/25ApiWithMongoDB/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

//Most Important
var collection *mongo.Collection

// connectin with mongo db
func init() {
	// getting details from .env files
	dbUser := goDotEnvVariable("DB_USER")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbName := goDotEnvVariable("DB_NAME")
	colName := goDotEnvVariable("COL_NAME")

	connectionString := "mongodb+srv://" + dbUser + ":" + dbPassword + "@cluster0.omnwb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	// check for error
	checkError(err)

	fmt.Println("MongoDB connection success.")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Connection instance is ready.")
}

// -------- MongoDB helpers -------- // goes to other file

// insert a movie
func insertAMovie(movie model.Netflix) {
	// add a movie
	inserted, err := collection.InsertOne(context.Background(), movie)

	// check for errors
	checkError(err)
	// success message with movie ID
	fmt.Println("Movie inserted successfully in DB with id: ", inserted.InsertedID)
}

// update a movie
func updateAMovie(movieId string) {
	// getting id that mongo accepts
	id, err := primitive.ObjectIDFromHex(movieId)
	checkError(err)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	checkError(err)

	// message after success
	fmt.Println("Movie update successful...")
	fmt.Println("Modified Count is: ", result.ModifiedCount)
}

// delete a movie
func deleteAMovie(movieId string) {
	// id for mongoDb
	id, err := primitive.ObjectIDFromHex(movieId)
	checkError(err)

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	checkError(err)

	fmt.Println("Movie deleting successful...")
	fmt.Println("Deleteed Item count: ", deleteCount)
}

// delete all movies
func deleteAllMovie() {

	// filter := bson.D{{}}
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	checkError(err)

	fmt.Println("Deleting all movies successful...")
	fmt.Println("Deleted Items count: ", deleteCount.DeletedCount)
}

// gvet all movies from mongoDB and returns
func getAllMovie() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	// check for errors
	checkError(err)

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)

		// check for error
		checkError(err)
		movies = append(movies, movie)
	}

	// executes at end, You remember
	defer cur.Close(context.Background())

	return movies
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Another controller that can be used in other packages
//--------- Actual controller -------- that can be Exported
// other file
func GetAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}

func InsertAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// insert into db
	insertAMovie(movie)
	json.NewEncoder(w).Encode(movie)
	fmt.Println("movie added...")
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// getting id from router
	params := mux.Vars(r)

	// passing and calling update
	updateAMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	fmt.Println("movie updated...")
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteAMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
	fmt.Println("movie deleted...")
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovie()
	fmt.Println("all movie deleted...")
}
