package endpoints

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

func EnterUserData(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	err := r.ParseForm()
	if err != nil {
		log.Print("error parsing form: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	lang := r.FormValue("l1")
	age, err := strconv.Atoi(r.FormValue("age"))
	id := r.FormValue("id")
	if err != nil {
		log.Print("invalid age")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rows, err := db.Query(context.Background(), "UPDATE users SET name = $1, age = $2, lang = $3 where id = $4", name, age, lang, id)
	defer rows.Close()
	if err != nil {
		log.Print("error updating user:" + err.Error())
	}
	http.Redirect(w, r, "/test", http.StatusPermanentRedirect)

}
