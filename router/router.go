package router

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Handlefunc("/")
	router.Handlefunc("/")
	router.Handlefunc("/")
	router.Handlefunc("/")
	router.Handlefunc("/")
}
