package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/internal/app/db"
	"server/internal/app/todolist/list"
)

var dbConfiguration *db.ConfigurationDB

func main() {
	initializeDatabase()
	r := configureRoutes()
	serve(r)

}

func initializeDatabase() {
	dbConfiguration = &db.ConfigurationDB{}
	dbConfiguration.InitializeDatabase()
	dbConfiguration.Repository = make([]db.Repository, 0)
	dbConfiguration.Repository = append(dbConfiguration.Repository, &list.ListRepository{
		Connection: dbConfiguration,
		Model:      list.TodoListModel{},
	})
	for _, value := range dbConfiguration.Repository {
		value.Migrate()
	}

}

func configureRoutes() *mux.Router {
	listRepository := &list.ListRepository{
		Connection: dbConfiguration,
		Model:      list.TodoListModel{},
	}

	listService := list.New(listRepository)

	r := mux.NewRouter()
	r.HandleFunc("/todo", listService.ListAll).Methods("GET")
	r.HandleFunc("/todo", listService.Save).Methods("POST")
	r.HandleFunc("/todo/{id}", listService.Delete).Methods("DELETE")
	r.HandleFunc("/todo/{id}", listService.Update).Methods("PUT")
	return r
}
func serve(r *mux.Router) {

	allAllowedMethods := make([]string, 0)
	allAllowedMethods = append(allAllowedMethods, "GET", "POST", "PUT", "PATCH", "DELETE")
	r.Methods("GET", "POST", "PUT", "PATCH", "DELETE")

	methods := handlers.AllowedMethods(allAllowedMethods)
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-request-with", "Access-Control-Allow-Origin", "Content-Type"})

	srv := &http.Server{
		Handler: handlers.CORS(methods, origins, headers)(r),
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())
}
