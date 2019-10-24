package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/workshops", GetWorkshops).Methods("GET")
	return router
}

func GetWorkshops(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)
	w.Write(nil)
}
