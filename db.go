package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var db *sql.DB

func InitDB() {
	val, err := sql.Open("sqlite3", "./TC.sqlite")
	db = val
	if err != nil {
		panic(err.Error())
	}
}

func handleEvent(eventType EventType, state EventState) {
	/*
		switch eventType {
		case unknownEventType:
		default:
			errorHandler("db.handleEvent(): " +
				"EventType: " + strconv.Itoa(int(eventType)) + " " +
				"EventState: " + strconv.Itoa(int(state)), 1)
		case creatureEventType:
		case workEventType:
		}
	*/

	addEvent(time.Now(), eventType, state)
}

func addEvent(time time.Time, eventType EventType, state EventState) int64 {
	stmt, err := db.Prepare("INSERT INTO Events(moment_workTime, type, state) values(?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	res, err := stmt.Exec(time, eventType, state)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return id
}
