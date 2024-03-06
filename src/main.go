package main

import (
	"context"
	"eelsAAA/components"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("cannot connect to DB")
		os.Exit(1)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("eelsAAAId")
		showTest := err != nil
		components.Base(showTest).Render(context.Background(), w)
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		testPage(w, r, db)
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		handleResult(r, db)
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		showData(w, r, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = http.ListenAndServe(":"+port, nil)

}
