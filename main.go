package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tomr1233/system3-api/internal/database"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	AgentId    string    `json:"agent_id"`
	VisitCount int32     `json:"visit_count"`
	HasCalled  bool      `json:"has_called"`
}

type apiConfig struct {
	db *database.Queries
}

func corsMiddleware(next http.Handler) http.Handler {
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		origin = "*"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	_ = godotenv.Load() // optional: .env not present in Docker (env injected by compose)
	// Get the connection string from the environment variable
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		fmt.Fprintf(os.Stderr, "DATABASE_URL not set\n")
		os.Exit(1)
	}
	// Connect to the database
	conn, err := sql.Open("pgx", connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connection established")
	cfg := apiConfig{
		db: database.New(conn),
	}
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: corsMiddleware(mux),
	}
	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	//  returns { name, agent_id, slug, visit_count, has_called }
	mux.HandleFunc("GET /api/visitors/{slug}", cfg.handlerGetVisitor)
	mux.HandleFunc("POST /api/visitors", cfg.handlerCreateVisitor)
	// accepts { slug }, returns { visit_count }
	mux.HandleFunc("POST /api/visitors/increment-visit", cfg.handlerIncrementVisit)
	// accepts { slug }
	mux.HandleFunc("POST /api/visitors/update-has-called", cfg.handlerHasCalled)
	log.Fatal(server.ListenAndServe())
}
