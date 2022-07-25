package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Username_and_Password(w http.ResponseWriter, r *http.Request) {
	ul := UserCredentials{}
	err := json.NewDecoder(r.Body).Decode(&ul)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("ERROR failed to decode request: " + err.Error())
		return
	}
	if ul.Login == "" || ul.Password == "" {
		http.Error(w, "Empty request", http.StatusBadRequest)
		log.Println("ERROR empty credentials")
		return
	}
	dt := time.Now()
	log.Println(dt.String() + "account with nickname:" + ul.Login + " was created")
	info[ul.Login] = ul
	w.WriteHeader(201)
}
