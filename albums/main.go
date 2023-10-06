package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// making a struct that has fields(details of an album) needed to be stored as data
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

// making a slice that already have some data in it to work with
var albums = []album{
	{ID: "1", Title: "MMLP", Artist: "Eminem", Price: "$300"},
	{ID: "1", Title: "MMLP2", Artist: "Eminem", Price: "$400"},
	{ID: "1", Title: "Kamikaze", Artist: "Eminem", Price: "$500"},
	{ID: "1", Title: "Music to be murdered by", Artist: "Eminem",
		Price: "$600"},
}

var Port = 3000

func main() {

	fmt.Println("Server is running at port", Port)

	r := mux.NewRouter()
	r.HandleFunc("/albums", getAlbums).Methods("GET")
	r.HandleFunc("/album/{id}", getAlbumById).Methods("GET")
	r.HandleFunc("/albums", postAlbums).Methods("POST")

	http.ListenAndServe(":3000", nil)
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	//setting the content type to json
	w.Header().Set("Content-Type", "application/json")

	//encoding the list of albums as json and sending it as the response
	json.NewEncoder(w).Encode(albums)
	// if err := json.NewEncoder(w).Encode(&albums); err != nil {
	//  http.Error(w, err.Error(), http.StatusInternalServerError)
	//  return
	// }
}

func getAlbumById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //created a variable to store id, here r is request object which retrieves route parameters with the help of m (in this case it is id)
	for _, album := range albums {
		if album.ID == params["id"] {
			json.NewEncoder(w).Encode(album)
			return
		}

	}
}

func postAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var new_album album
	//decode is used to create entries(sending data to server)
	_ = json.NewDecoder(r.Body).Decode(&new_album)
	new_album.ID = strconv.Itoa(rand.Intn(1000000000))

	albums = append(albums, new_album)
	//encode is used to get data from server
	json.NewEncoder(w).Encode(albums)

}
