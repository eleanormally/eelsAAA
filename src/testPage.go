package main

import (
	"context"
	"eelsAAA/components"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func testPage(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	_ = r

	var id string
	err := db.QueryRow(context.Background(), "insert into users (name) values ($1) returning id", "test").Scan(&id)
	if err != nil {
		log.Print(err.Error())
		return
	}
	w.Header().Add("set-cookie", "eelsAAAId="+string(id)+";")

	// getting word list
	base := []components.WordPair{
		{
			Word:    "blade",
			NonWord: "charrel",
			Id:      -1,
		},
		{
			Word:    "future",
			NonWord: "heru",
			Id:      -1,
		},
		{
			Word:    "invitation",
			NonWord: "masler",
			Id:      -1,
		},
		{
			Word:    "ranch",
			NonWord: "plapforb",
			Id:      -1,
		},
		{
			Word:    "scale",
			NonWord: "strofe",
			Id:      -1,
		},
		{
			Word:    "verse",
			NonWord: "zobe",
			Id:      -1,
		},
	}
	value, err := db.Query(context.Background(), "SELECT word, nonword, id FROM \"wordPairs\" ORDER BY RANDOM() LIMIT 50")
	if err != nil {
		log.Print(err.Error())
		return
	}
	for value.Next() {
		var newWordPair components.WordPair
		value.Scan(&newWordPair.Word, &newWordPair.NonWord, &newWordPair.Id)
		base = append(base, newWordPair)
	}

	components.Tester(id, base).Render(context.Background(), w)
}
