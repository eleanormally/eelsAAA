package main

import (
	"context"
	"eelsAAA/endpoints"
	"eelsAAA/views"
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
		views.Homepage(w)
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		views.TestPage(w, r, db)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		views.UserInfo(w, db)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		endpoints.EnterTestResult(r, db)
	})

	http.HandleFunc("/postUser", func(w http.ResponseWriter, r *http.Request) {
		endpoints.EnterUserData(w, r, db)
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		views.Admin(w, db)
	})
	http.HandleFunc("/writeup", func(w http.ResponseWriter, r *http.Request) {
		views.Writeup(w)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = http.ListenAndServe(":"+port, nil)

}
