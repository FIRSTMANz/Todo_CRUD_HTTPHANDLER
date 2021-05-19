package main

import (
	"net/http"
	"projectfirsty/controller"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getList", controller.GetListData).Methods(http.MethodGet)
	r.HandleFunc("/add", controller.Insert).Methods(http.MethodPost)
	r.HandleFunc("/update", controller.Update).Methods(http.MethodPost)
	r.HandleFunc("/delete", controller.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/show", controller.Show).Methods(http.MethodGet)
	http.ListenAndServe(":8081", r)
}
