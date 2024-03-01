package handlers

import (
	"goods/internal/services"
	"log"
	"net/http"
)

func InitHandlers(logger *log.Logger, services *services.Services) *http.ServeMux {
	mux := http.NewServeMux()
	// mux.Handle("/good", newGoodHandler(logger))
	mux.HandleFunc("/good/create", createGood)
	mux.HandleFunc("/good/update", updateGood)
	mux.HandleFunc("/good/remove", removeGood)
	mux.HandleFunc("/goods/list", listGoods)
	mux.HandleFunc("/good/reprioritiize", repriotGood)
	return mux
}

func createGood(rw http.ResponseWriter, r *http.Request) {

}

func updateGood(rw http.ResponseWriter, r *http.Request) {

}

func removeGood(rw http.ResponseWriter, r *http.Request) {

}

func listGoods(rw http.ResponseWriter, r *http.Request) {

}

func repriotGood(rw http.ResponseWriter, r *http.Request) {

}
