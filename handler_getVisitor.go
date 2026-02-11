package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerGetVisitor(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	user, err := cfg.db.GetVisitorBySlug(r.Context(), slug)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte("Couldn't get visitor"))
		return
	}

	resp := User{
		ID:         int(user.ID),
		CreatedAt:  user.CreatedAt.Time,
		UpdatedAt:  user.UpdatedAt.Time,
		Slug:       user.Slug,
		Name:       user.Name,
		AgentId:    user.AgentID.String,
		VisitCount: user.VisitCount.Int32,
		HasCalled:  user.HasCalled.Bool,
	}
	dat, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)

}
