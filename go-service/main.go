package main

import (
	"context"
	"fmt"
	"go-service/controllers"
	"go-service/database"
	"go-service/kafka"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
	ctx := context.Background()
	kafka.Consume(ctx)
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/get", controllers.GetAllBlogEntries).Methods("GET")
	router.HandleFunc("/create", controllers.CreateBlogEntry).Methods("POST")
	router.HandleFunc("/get/{id}", controllers.GetBlogEntryByID).Methods("GET")
	router.HandleFunc("/update", controllers.UpdateBlogEntry).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteBlogEntryByID).Methods("DELETE")
}

func initDB() {
	var config database.Config
	config.Server = "PCANDRES\\MSSQL"
	config.Port = 1433
	config.User = "sa"
	config.Password = "test123"
	config.DataBaseName = "devdb"
	var connectionString = database.GetConnectionString(config)
	var err = database.Connect(connectionString)
	if err != nil {

		fmt.Printf(fmt.Sprintf("server=%s", err.Error()))
	}

}
