package main

import (
	"database/sql"
	"encoding/json"
	"github.com/tomr1233/system3-api/internal/database"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerCreateVisitor(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Slug    string `json:"slug"`
		Name    string `json:"name"`
		AgentID string `json:"agent_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(500)
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Slug: params.Slug,
		Name: params.Name,
		AgentID: sql.NullString{
			String: params.AgentID,
			Valid:  params.AgentID != "",
		},
	})
	if err != nil {
		log.Printf("Error creating user: %s", err)
		w.WriteHeader(500)
		w.Write([]byte("Couldn't create user"))
		return
	}
	resp := User{
		ID:        int(user.ID),
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
		Slug:      user.Slug,
		Name:      user.Name,
		AgentId:   user.AgentID.String,
	}
	dat, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(dat)
}
