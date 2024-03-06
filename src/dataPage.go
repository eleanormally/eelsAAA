package main

import (
	"context"
	"eelsAAA/components"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func showData(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {

	_, err := r.Cookie("eelsAAAAdmin")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var resultCount int
	db.QueryRow(context.Background(), "SELECT count(*) from results").Scan(&resultCount)
	var totalUsers int
	var inactiveUsers int
	db.QueryRow(context.Background(), "SELECT count(*) from users").Scan(&totalUsers)
	db.QueryRow(context.Background(), "SELECT count(*) from users as u WHERE NOT EXISTS (select * from results as r where r.user = u.id)").Scan(&inactiveUsers)
	var complete int
	db.QueryRow(context.Background(), "SELECT count(*) from (SELECT count(*) from results as r group by r.user) where count >= 50").Scan(&complete)

	components.Data(resultCount, totalUsers, inactiveUsers, complete).Render(context.Background(), w)
}
