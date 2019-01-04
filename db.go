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

func handleEvent(idTimer int, state EventState) {
	/*
		switch idTimer {
		case unknownEventType:
		default:
			errorHandler("db.handleEvent(): " +
				"EventType: " + strconv.Itoa(int(idTimer)) + " " +
				"EventState: " + strconv.Itoa(int(state)), 1)
		case creatureEventType:
		case workEventType:
		}
	*/

	addEvent(time.Now(), idTimer, state)
}

func addEvent(time time.Time, idTimer int, state EventState) int64 {
	stmt, err := db.Prepare("INSERT INTO Events(moment_workTime, type_timer_id, state) values(?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	res, err := stmt.Exec(time, idTimer, state)
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

func LoadAppData() AppData {
	rows, err := db.Query("SELECT * FROM AppData;")
	if err != nil {
		fmt.Println(err.Error())
		return AppData{}
	}
	rows.Next()
	appData := AppData{}
	err = rows.Scan(&appData.Title, &appData.Tooltip, &appData.Icon)
	return appData
}

func LoadTimers() []Timer {
	rows, err := db.Query("SELECT * FROM Timers;")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var timers []Timer
	for rows.Next() {
		timer := Timer{}
		err = rows.Scan(&timer.Id, &timer.Name)
		timers = append(timers, timer)
	}
	return timers
}
