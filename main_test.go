package main

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {

	t.Run(
		"Createが成功するケース",
		func(t *testing.T) {
			// Arrange
			todo := &ToDo{
				Title: "testToDo",
			}
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO todo(title) VALUES ( ? )")).
				WithArgs(todo.Title).
				WillReturnResult(sqlmock.NewResult(1, 1))

			// Act
			err = Create(db, todo)

			// Assert
			if err != nil {
				t.Error(err.Error())
			}
		},
	)

	t.Run(
		"Createが失敗するケース",
		func(t *testing.T) {
			// Arrange
			todo := &ToDo{
				Title: "testToDo",
			}
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO todo(title) VALUES ( ? )")).
				WithArgs(todo.Title).
				WillReturnResult(sqlmock.NewErrorResult(errors.New("ERROR!!!"))).
				WillReturnError(errors.New("INSERT FAILED!!!"))

			// Act
			err = Create(db, todo)

			// Assert
			if err == nil {
				t.Error("An error should have occurred.")
			}
		},
	)

}

func TestRead(t *testing.T) {

	t.Run(
		"Readが成功するケース",
		func(t *testing.T) {
			// Arrange
			id := 1
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title FROM todo WHERE id = ?")).
				WithArgs(id).
				WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(1, "testToDo"))

			// Act
			_, err = Read(db, id)

			// Assert
			if err != nil {
				t.Error(err.Error())
			}
		},
	)

	t.Run(
		"Readが失敗するケース(QueryRowでエラー)",
		func(t *testing.T) {
			// Arrange
			id := 1
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title FROM todo WHERE id = ?")).
				WithArgs(id).
				WillReturnError(errors.New("SELECT FAILED!!!"))

			// Act
			_, err = Read(db, id)

			// Assert
			if err == nil {
				t.Error("An error should have occurred.")
			}
		},
	)

	t.Run(
		"Readが失敗するケース(Scanでエラー)",
		func(t *testing.T) {
			// Arrange
			id := 1
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title FROM todo WHERE id = ?")).
				WithArgs(id).
				WillReturnRows(sqlmock.NewRows([]string{"hogehoge"}).AddRow("fugafuga"))

			// Act
			_, err = Read(db, id)

			// Assert
			if err == nil {
				t.Error("An error should have occurred.")
			}
		},
	)

}

func TestUpdate(t *testing.T) {

	t.Run(
		"Updateが成功するケース",
		func(t *testing.T) {
			// Arrange
			todo := &ToDo{
				Id:    1,
				Title: "testToDo",
			}
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("UPDATE todo SET title = ? WHERE id = ?")).
				WithArgs(todo.Title, todo.Id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			// Act
			err = Update(db, todo)

			// Assert
			if err != nil {
				t.Error(err.Error())
			}
		},
	)

	t.Run(
		"Updateが失敗するケース",
		func(t *testing.T) {
			// Arrange
			todo := &ToDo{
				Id:    1,
				Title: "testToDo",
			}
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("UPDATE todo SET title = ? WHERE id = ?")).
				WithArgs(todo.Title, todo.Id).
				WillReturnResult(sqlmock.NewErrorResult(errors.New("ERROR!!!"))).
				WillReturnError(errors.New("UPDATE FAILED!!!"))

			// Act
			err = Update(db, todo)

			// Assert
			if err == nil {
				t.Error("An error should have occurred.")
			}
		},
	)

}

func TestDelete(t *testing.T) {

	t.Run(
		"Deleteが成功するケース",
		func(t *testing.T) {
			// Arrange
			id := 1
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo WHERE id = ?")).
				WithArgs(id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			// Act
			err = Delete(db, id)

			// Assert
			if err != nil {
				t.Error(err.Error())
			}
		},
	)

	t.Run(
		"Deleteが失敗するケース",
		func(t *testing.T) {
			// Arrange
			id := 1
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo WHERE id = ?")).
				WithArgs(id).
				WillReturnResult(sqlmock.NewErrorResult(errors.New("ERROR!!!"))).
				WillReturnError(errors.New("DELETE FAILED!!!"))

			// Act
			err = Delete(db, id)

			// Assert
			if err == nil {
				t.Error("An error should have occurred.")
			}
		},
	)

}
