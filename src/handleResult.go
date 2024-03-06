package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Response struct {
	Correct bool `json:"correct"`
	Id      int  `json:"id"`
	Time    int  `json:"time"`
}

func handleResult(r *http.Request, db *pgxpool.Pool) {
	//body response
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading body")
		return
	}
	var value Response
	err = json.Unmarshal(body, &value)

	//checking for valid cookie id
	userId, err := r.Cookie("eelsAAAId")
	if err != nil {
		fmt.Println("error getting user id")
		return
	}
	intId, err := strconv.Atoi(userId.Value)
	if err != nil {
		fmt.Println("invalid id")
	}

	_, err = db.Query(context.Background(), "INSERT INTO results (\"user\", correct, time, pair_id) VALUES ($1, $2, $3, $4)", intId, value.Correct, value.Time, value.Id)
	if err != nil {
		fmt.Println("Error Setting DB Result: " + err.Error())
	}

}
