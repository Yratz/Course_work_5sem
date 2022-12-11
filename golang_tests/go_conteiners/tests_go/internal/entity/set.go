package entity

import "github.com/jackc/pgx/v5"

var db *pgx.Conn

func SetDb(db_new *pgx.Conn) {
	db = db_new
}

type error_res struct {
	Err_s string `json: "error_lol"`
}

type Status struct {
	Status string `json:"status"`
}
