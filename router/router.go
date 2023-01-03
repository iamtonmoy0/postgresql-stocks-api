package router

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Handlefunc("/api/stock/{id}", middle.GetStock).Methods("GET", "OPTIONS")
	router.Handlefunc("/api/stock", middle.GetAllStock).Methods("GET", "OPTIONS")
	router.Handlefunc("/api/newstock", middle.CreateStock).Methods("POST", "OPTIONS")
	router.Handlefunc("/api/stock/{id}", middle.UpdateStock).Methods("PUT", "OPTIONS")
	router.Handlefunc("/api/deletestock/{id}", middle.DeleteStock).Methods("DELETE", "OPTIONS")
}
