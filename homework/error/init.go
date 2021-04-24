package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func init() {
	db, err := sql.Open("sqlite3", "./foo.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := `
    DROP TABLE IF EXISTS user;
    create table user (name varchar(50) not null,sex varchar(2) not null);

    insert into user (name,sex) values ('foo','M');
    insert into user (name,sex) values ('foo1','M');
    insert into user (name,sex) values ('foo2','M');
    insert into user (name,sex) values ('bar','W');
    insert into user (name,sex) values ('bar1','W');
    insert into user (name,sex) values ('bar2','W');
        `

	_, err = db.Exec(q)
	if err != nil {
		panic(err)
	}
}
