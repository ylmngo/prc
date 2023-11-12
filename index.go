package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var dsn = "postgres://usermessages_user:baAHOs1dyGNtrRHDynaAZvlqAN7nB4vC@dpg-cl7j1iquuipc73ehmc60-a.singapore-postgres.render.com/usermessages"

type application struct {
	db *sql.DB
}

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	name := "Aamer"
	query := `INSERT INTO test (name) 
			  VALUES ($1)`

	_ = app.db.QueryRow(query, name)
}

func main() {
	app := &application{}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	app.db = db

	http.HandleFunc("/", app.index)
	http.ListenAndServe(":8000", nil)
}
