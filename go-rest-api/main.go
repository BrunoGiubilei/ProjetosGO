package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/brunogiubilei/go-rest-api/handlers"
	"github.com/brunogiubilei/go-rest-api/middleware"
	"github.com/gorilla/mux"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func migrations() {
	// Conectar ao banco de dados SQLite
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicializar o driver de migração
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Criar uma nova instância de migração
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Aplicar as migrações
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migrações aplicadas com sucesso")

	db.Close()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func routes() {
	route := mux.NewRouter()

	route.Use(loggingMiddleware)
	route.Use(middleware.Cors)

	route.HandleFunc("/api", handlers.GetAPI).Methods("GET")
	route.HandleFunc("/api/greeting", handlers.GetGreeting).Methods("GET")
	route.HandleFunc("/api/adduser", handlers.AddUser).Methods("POST", "OPTIONS")
	route.HandleFunc("/api/getusers", handlers.GetUsers).Methods("GET")
	route.HandleFunc("/api/removeuser/{id}", handlers.DeleteUser).Methods("DELETE", "OPTIONS")
	route.HandleFunc("/api/alteruser/{id}", handlers.UpdateUser).Methods("PUT", "OPTIONS")

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}

func main() {
	migrations()
	routes()
}
