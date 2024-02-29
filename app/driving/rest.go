package driving

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restAPIServer/app/driven/mongo"
	"restAPIServer/app/service"
	"strings"
)

func Start(path string) {
	client := GetDbClientMongo()

	router := mux.NewRouter()

	repositoryDb := mongo.NewRepositoryDb(client)
	h := Handlers{service.NewService(repositoryDb)}

	router.HandleFunc("/home", home).Methods(http.MethodGet)
	router.HandleFunc("/api/1.0/users", getBody).Methods(http.MethodPost)
	router.HandleFunc("/mongotest", h.MongoTest).Methods(http.MethodGet)
	router.HandleFunc("/mongotest/id", h.MongoTestById).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func HeadersToString(headers http.Header) string {
	var headerStrings []string

	for key, values := range headers {
		for _, value := range values {
			headerStrings = append(headerStrings, fmt.Sprintf("%s: %s", key, value))
		}
	}

	return strings.Join(headerStrings, "\n")
}
