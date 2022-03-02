package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-memdb"
	"github.com/rs/cors"
)

func main() {
	db, err := memdb.NewMemDB(newSchema())
	if err != nil {
		log.Fatalln(err)
	}
	store := NewTodoStore(db)

	controller := NewTodoController(store)

	router := mux.NewRouter().StrictSlash(true)

	globalPath := os.Getenv("GLOBAL_PATH")

	router.Name("List").Methods(http.MethodGet).Path(globalPath + "/todos").Handler(logger(controller.List(), "List"))
	router.Name("Get").Methods(http.MethodGet).Path(globalPath + "/todos/{id}").Handler(logger(controller.Get(), "Get"))
	router.Name("Insert").Methods(http.MethodPost).Path(globalPath + "/todos").Handler(logger(controller.Insert(), "Insert"))
	router.Name("Delete").Methods(http.MethodDelete).Path(globalPath + "/todos/{id}").Handler(logger(controller.Delete(), "Delete"))

	c := cors.New(cors.Options{
		ExposedHeaders: []string{"Authorization", "Content-Disposition"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	log.Fatalln(http.ListenAndServe(":55555", c.Handler(router)))
}
