package main

import (
	codeHelpCoreGen "api-spec/generated/code-help-core"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type errorResponse struct {
	Message string `json:"message"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(albums); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	} else {
	}
}

func postAlbums(writer http.ResponseWriter, request *http.Request) {
	var newAlbum album

	if err := json.NewDecoder(request.Body).Decode(&newAlbum); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	albums = append(albums, newAlbum)

	if err := json.NewEncoder(writer).Encode(newAlbum); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func getAlbumById(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")

	for _, v := range albums {
		if v.ID == id {
			if err := json.NewEncoder(writer).Encode(v); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			writer.WriteHeader(http.StatusOK)
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
	json.NewEncoder(writer).Encode(errorResponse{Message: "Album not found"})
}

func fetchCategories(client *codeHelpCoreGen.ClientWithResponses) {
	if c, err := client.GetCategories(context.Background()); err == nil {
		var category codeHelpCoreGen.Category
		if err := json.NewDecoder(c.Body).Decode(&category); err != nil {
			return
		}

		log.Println(category)
	} else {
		log.Fatal(err)
	}
}

func main() {
	//if client, err := CreateClient(); err != nil {
	//	log.Fatal(err)
	//	return
	//} else {
	//	fetchCategories(client)
	//}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /albums", getAlbums)
	mux.HandleFunc("POST /albums", postAlbums)
	mux.HandleFunc("GET /albums/{id}", getAlbumById)

	server := http.Server{
		Addr:    ":5080",
		Handler: mux,
	}

	log.Println(server.ListenAndServe())
}
