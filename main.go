package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	Message string `json:"msg"`
}

type Response struct {
	Answer string `json:"answer"`
}

type UserCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserPets struct {
	Pets string `json:"pet"`
}

var info map[string]UserCredentials
var petss map[string]UserPets

func main() {
	info = map[string]UserCredentials{}
	r := mux.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	RegisterHandlers(r)
	log.Println("Starti ng server...")
	http.ListenAndServe("0.0.0.0:8000", r)
}

func HandleDataRequest(w http.ResponseWriter, r *http.Request) {
	req := Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("ERROR failed to decode request: " + err.Error())
		return
	}
	if req.Message == "" {
		http.Error(w, "Empty request", http.StatusBadRequest)
		log.Println("ERROR empty request")
		return
	}
	answ := Response{Answer: "Thanks for suplying: " + req.Message + ":"}
	err = json.NewEncoder(w).Encode(answ)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("ERROR failed to encode request")
		return
	}
	log.Println("INFO processed data")
}

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/", HandleDataRequest).Methods("POST")
	r.HandleFunc("/user/", Username_and_Password).Methods("POST")
	r.HandleFunc("/pets/", petsss).Methods("POST")
	r.HandleFunc("/contacts/", ContactInfo)
}
func ContactInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("+9965555555 \n golang@gmail.com \n https://wa.me/1654613213")
}
