package routes

import (
	"api_go_rest/controllers"
	"api_go_rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personality/{id}", controllers.GetPersonalityById).Methods("GET")
	r.HandleFunc("/api/personality", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personality/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personality", controllers.PutPersonality).Methods("PUT")

	// CORS configuration to accept requests from any source
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
