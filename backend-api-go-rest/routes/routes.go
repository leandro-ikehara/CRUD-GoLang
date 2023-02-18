package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/leandro-ikehara/Estudos-GoLang/go-rest-api/controller"
	"github.com/leandro-ikehara/Estudos-GoLang/go-rest-api/routes/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controller.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controller.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controller.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controller.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

// CONTINUAR EM ALURA - 05 Go Middleware e Front: desenvolvendo uma API Rest
