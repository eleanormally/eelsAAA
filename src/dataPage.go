package main

import (
	"context"
	"eelsAAA/components"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func userData(db *pgxpool.Pool) (int, int, int, error) {
	var totalUsers int
	var inactiveUsers int
	err := db.QueryRow(context.Background(), "SELECT count(*) from users").Scan(&totalUsers)
	if err != nil {
		return 0, 0, 0, err
	}
	err = db.QueryRow(context.Background(), "SELECT count(*) from users as u WHERE NOT EXISTS (select * from results as r where r.user = u.id ) ").Scan(&inactiveUsers)
	if err != nil {
		return 0, 0, 0, err
	}
	var complete int
	err = db.QueryRow(context.Background(), "SELECT count(*) from (SELECT count(*) from results group by user) as r where r.count >= 64").Scan(&complete)
	if err != nil {
		return 0, 0, 0, err
	}
	return totalUsers, inactiveUsers, complete, nil
}

func showData(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {

	_, err := r.Cookie("eelsAAAAdmin")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	totalUsers, inactiveUsers, complete, err := userData(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	var resultCount int
	db.QueryRow(context.Background(), "SELECT count(*) from results").Scan(&resultCount)

	var entries []components.Entry
	values, err := db.Query(context.Background(),
		`SELECT wp.word, r.time, wp.freq, wp.aoa from results as r inner join "wordPairs" as wp on wp.id = r.pair_id where r.correct = true and r.word = true`,
	)
	defer values.Close()
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for values.Next() {
		var entry components.Entry
		values.Scan(&entry.Word, &entry.Time, &entry.Freq, &entry.Aoa)
		entries = append(entries, entry)
	}

	var incorrects []components.Entry
	values, err = db.Query(context.Background(),
		`SELECT wp.nonword, r.time, wp.freq, wp.aoa from results as r inner join "wordPairs" as wp on wp.id = r.pair_id where r.correct = true and r.word = false`,
	)
	defer values.Close()
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for values.Next() {
		var entry components.Entry
		values.Scan(&entry.Word, &entry.Time, &entry.Freq, &entry.Aoa)
		incorrects = append(incorrects, entry)
	}

	components.Data(
		resultCount,
		totalUsers,
		inactiveUsers,
		complete,
		components.EntryVisualizer(incorrects, "Non Word Results", 400),
		components.EntryVisualizer(entries, "Correct Results", 400),
	).Render(context.Background(), w)

}
