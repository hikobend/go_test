package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// model
type ToDo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func Create(db *sql.DB, t *ToDo) error {
	_, err := db.Exec("INSERT INTO todo(title) VALUES ( ? )", t.Title)
	return err
}

func Read(db *sql.DB, id int) (*ToDo, error) {
	t := &ToDo{}
	err := db.QueryRow("SELECT id, title FROM todo WHERE id = ?", id).Scan(&t.Id, &t.Title)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func Update(db *sql.DB, t *ToDo) error {
	_, err := db.Exec("UPDATE todo SET title = ? WHERE id = ?", t.Title, t.Id)
	return err
}

func Delete(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
	return err
}

func main() {
	d, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	todo := &ToDo{Title: "read books"}
	err = Create(d, todo)
	if err != nil {
		log.Fatal(err)
	}
}
