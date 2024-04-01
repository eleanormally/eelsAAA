package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type response struct {
	Correct bool `json:"correct"`
	Id      int  `json:"id"`
	Time    int  `json:"time"`
	Word    bool `json:"word"`
}

func EnterTestResult(r *http.Request, db *pgxpool.Pool) {
	//body response
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading body")
		return
	}
	var value response
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

	res, err := db.Query(context.Background(),
		`INSERT INTO results ("user", correct, time, pair_id, word) VALUES ($1, $2, $3, $4, $5)`,
		intId,
		value.Correct,
		value.Time,
		value.Id,
		value.Word,
	)

	if err != nil {
		fmt.Println("Error Setting DB Result: " + err.Error())
	}
	defer res.Close()

}
