package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func petsss(w http.ResponseWriter, r *http.Request) {
	pet := UserPets{}
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("ERROR failed to decode request: " + err.Error())
		return
	}
	if pet.Pets == "" {
		http.Error(w, "Empty request", http.StatusBadRequest)
		log.Println("ERROR empty credentials")
		return
	}
	log.Println("Pet created")
	petss[pet.Pets] = pet
	w.WriteHeader(201)
	fmt.Println(w, "Pet was created with a name of: "+pet.Pets)
}
