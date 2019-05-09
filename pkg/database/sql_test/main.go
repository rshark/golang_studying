package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db sql.DB

func main() {
	// 连接数据库
	db, err := sql.Open("mysql", "fengbo:FengBo12.@tcp(127.0.0.1:3306)/Hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		name string
	)

	// 使用Query查询
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	// prepared statement
	stmt, err := db.Prepare("select id, name from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	rows.Close()
	stmt.Close()

	// single query row
	err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	stmt, err = db.Prepare("select name from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
	stmt.Close()

	// statement modify data
	rows, err = db.Query("select count(*) from users")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("INSERT INTO users(id, name) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	var insertedID int
	for rows.Next() {
		err = rows.Scan(&insertedID)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()

	res, err := stmt.Exec(insertedID+1, "Dolly")
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	stmt.Close()

	// working with Transactions
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	/*
		stmt, err = tx.Prepare("INSERT INTO users VALUES (?)")
		if err != nil {
			log.Fatal(err)
		}

		stmt.Close()
	*/
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// error handling
	err = db.QueryRow("select name from users where id = ?", 10).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			println("there were no rows, but otherwise no error occured")
		} else {
			log.Fatal(err)
		}
	}

	// working with nulls
	rows, err = db.Query("select * from users")
	for rows.Next() {
		var (
			id      int
			name    sql.NullString
			address sql.NullString
		)
		rows.Scan(&id, &name, &address)
		if address.Valid {
			fmt.Printf("id %d address is not null \n", id)
		} else {
			fmt.Printf("id %d address is null \n", id)
		}
	}
	rows.Close()

	// working with unknown columns

}
