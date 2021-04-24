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
	var (
		err error
	)

	db, err = sql.Open("sqlite3", "./foo.db")

	if err != nil {
		log.Fatal(err)
	}

	q := `
    DROP TABLE IF EXISTS user;
    create table user (name varchar(50) not null PRIMARY KEY,passwd varchar(50) not null,sex varchar(2) not null);

    insert into user (name,passwd,sex) values ('foo','a1313','M');
    insert into user (name,passwd,sex) values ('foo1','1bb313','M');
    insert into user (name,passwd,sex) values ('foo2','13d13','M');
    insert into user (name,passwd,sex) values ('bar','13313','W');
    insert into user (name,passwd,sex) values ('bar1','13uu13','W');
    insert into user (name,passwd,sex) values ('bar2','131ii3','W');
    `

	_, err = db.Exec(q)
	if err != nil {
		panic(err)
	}
}
