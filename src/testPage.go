package main

import (
	"context"
	"eelsAAA/components"
	"log"
	"math/rand"
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
			Choice:  rand.Int()%2 + 1,
		},
		{
			Word:    "future",
			NonWord: "heru",
			Id:      -1,
			Choice:  rand.Int()%2 + 1,
		},
		{
			Word:    "invitation",
			NonWord: "masler",
			Id:      -1,
			Choice:  rand.Int()%2 + 1,
		},
		{
			Word:    "ranch",
			NonWord: "plapforb",
			Id:      -1,
			Choice:  rand.Int()%2 + 1,
		},
		{
			Word:    "scale",
			NonWord: "strofe",
			Id:      -1,
			Choice:  rand.Int()%2 + 1,
		},
		{
			Word:    "verse",
			NonWord: "zobe",
			Id:      -1,
			Choice:  rand.Int()%2 + 1,
		},
	}
	value, err := db.Query(context.Background(), `SELECT word, nonword, id, choice from (
(SELECT * FROM (SELECT *, NTILE(2) OVER ( ORDER BY RANDOM() ) as choice FROM
	(
    SELECT * FROM "wordPairs" as wp WHERE aoa = 'early' and freq = 'high' ORDER BY RANDOM() LIMIT 16) as wp
 )as wp) union
 (SELECT * FROM (SELECT *, NTILE(2) OVER ( ORDER BY RANDOM() ) as choice FROM
	(
    SELECT * FROM "wordPairs" as wp WHERE aoa = 'late' and freq = 'high' ORDER BY RANDOM() LIMIT 16) as wp
 )as wp) union
 (SELECT * FROM (SELECT *, NTILE(2) OVER ( ORDER BY RANDOM() ) as choice FROM
	(
    SELECT * FROM "wordPairs" as wp WHERE aoa = 'early' and freq = 'low' ORDER BY RANDOM() LIMIT 16) as wp
 )as wp) union
 (SELECT * FROM (SELECT *, NTILE(2) OVER ( ORDER BY RANDOM() ) as choice FROM
	(
    SELECT * FROM "wordPairs" as wp WHERE aoa = 'late' and freq = 'low' ORDER BY RANDOM() LIMIT 16) as wp
 )as wp)
) as final ORDER BY RANDOM()`)
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer value.Close()
	for value.Next() {
		var newWordPair components.WordPair
		value.Scan(&newWordPair.Word, &newWordPair.NonWord, &newWordPair.Id, &newWordPair.Choice)
		base = append(base, newWordPair)
	}

	components.Tester(id, base).Render(context.Background(), w)
}
