package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
)

type Book struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Date_Bought time.Time  `json:"date_bought"`
	Status      int        `json:"status"`
	Date_Read   *time.Time `json:"date_read"`
	Description *string    `json:"description"`
}

func main() {
	// DB Connection
	connStr := "postgres://postgres:postgres@localhost:5432/mybooks-db?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("unable to connecto to database", err)
	}
	defer conn.Close(context.Background())

	//Create router
	r := chi.NewRouter()

	// Enable CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Define Endpoints
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("New API"))
	})

	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		rows, err := conn.Query(context.Background(),
			"SELECT id, title, author, date_bought, status, date_read, description FROM public.books")
		if err != nil {
			http.Error(w, "Query Failed", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var books []Book
		for rows.Next() {
			var b Book
			if err := rows.Scan(&b.Id, &b.Title, &b.Author, &b.Date_Bought, &b.Status, &b.Date_Read, &b.Description); err != nil {
				http.Error(w, "Row scan failed: "+err.Error(), http.StatusInternalServerError)
				return
			}
			books = append(books, b)
			fmt.Println("Books: ", len(books))
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Row iteration failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	})

	http.ListenAndServe(":8000", r)
}
