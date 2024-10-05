package controllers

import (
	"api_go_rest/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	response := models.GetAllPersonalities()
	json.NewEncoder(w).Encode(response)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response := models.GetPersonalityById(id)
	json.NewEncoder(w).Encode(response)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result := models.CreatePersonality(personality)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(result.RowsAffected)
	}
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	result := models.DeletePersonality(id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(result.RowsAffected)
	}
}

func PutPersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result := models.PutPersonality(personality)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(result.RowsAffected)
	}
}
