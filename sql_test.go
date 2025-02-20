package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)


func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES ('assyaufi', 'Assyaufi')"

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer ORDER BY id DESC"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id:", id, ", name:", name)
	}

	defer rows.Close()
}

func TestDataTypeColumn(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		emailString := "NULL"
		if email.Valid {
			emailString = email.String
		}

		birthDateString := "NULL"
		if birthDate.Valid {
			birthDateString = birthDate.Time.Format("2006-01-02 15:04:05")
		}

		message := `---
id: %s,
name: %s,
email: %v,
balance: %d,
rating: %f,
birth_date: %v,
married: %t,
created_at: %s
`

		fmt.Printf(message, id, name, emailString, balance, rating, birthDateString, married, createdAt)
	}

	defer rows.Close()
}

func TestQuerySqlWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin"
	password := "secret"

	ctx := context.Background()

	query := "SELECT username FROM user WHERE username = '" + username +
	         "' AND password = '" + password +
			 "' LIMIT 1"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Welcome, %s!\n", username)
	} else {
		fmt.Println("Failed to login")
	}

	defer rows.Close()
}

func TestQuerySqlWithSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin'; #" // SQL injection from user input
	password := "wrong"     // wrong password

	ctx := context.Background()

	// BAD PRACTICE!
	query := "SELECT username FROM user WHERE username = '" + username +
	         "' AND password = '" + password +
			 "' LIMIT 1"

    fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Welcome, %s!\n", username)
	} else {
		fmt.Println("Failed to login")
	}

	defer rows.Close()
}

func TestQuerySqlWithSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin'; #" // SQL injection from user input
	password := "wrong"     // wrong password

	ctx := context.Background()

	// BAD PRACTICE!
	// query := "SELECT username FROM user WHERE username = '" + username +
	//          "' AND password = '" + password +
	// 		    "' LIMIT 1"

	// GOOD PRACTICE
	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

    fmt.Println(query)

	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Welcome, %s!\n", username)
	} else {
		fmt.Println("Failed to login")
	}

	defer rows.Close()
}

func TestExecSqlWithSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "rizqi'; DROP TABLE user; #"
	password := "secret"

	ctx := context.Background()

	query := "INSERT INTO user(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "rizqi@gmail.com"
	comment := "Hello! Saya sedang belajar bahasa pemrograman Go."

	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "rizqi_" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini adalah komentar ke " + strconv.Itoa(i)
		
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id:", id)
	}
}
