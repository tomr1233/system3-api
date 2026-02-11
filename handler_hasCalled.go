package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerHasCalled(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Slug string `json:"slug"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, err = cfg.db.SetHasCalled(r.Context(), params.Slug)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error setting has called"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	return
}
