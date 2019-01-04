package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type IndexData struct {
	Timers []TimerData
}

type TimerData struct {
	TimerInfo  Timer
	CountToday string
	CountWeek  string
	CountMonth string
}

type SettingsData struct {
	AppData AppData
	Timers  []Timer
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	indexData := &IndexData{}
	timers := LoadTimers()
	for _, v := range timers {
		countToday, countWeek, countMonth := calculateDates(v)
		indexData.Timers = append(indexData.Timers,
			TimerData{v, countToday, countWeek, countMonth},
		)
	}
	err := t.Execute(w, indexData)
	if err != nil {
		return
	}
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/settings.html")
	settingsData := SettingsData{}
	settingsData.AppData = LoadAppData()
	err := t.Execute(w, settingsData)
	if err != nil {
		return
	}
}

func RunWebServer() {
	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler)
	routes.HandleFunc("/menu", indexHandler)
	routes.HandleFunc("/settings", settingsHandler)

	http.Handle("/", routes)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8593", nil)
	if err != nil {
		fmt.Println("It's Server.RunWebServer() we got problem")
	}
}

func calculateDates(timer Timer) (countToday, countWeek, countMonth string) {
	events := LoadEventsByTimerId(timer.Id)
	if events == nil {
		return
	}
	/*lastId := -1
	isLastStart := false
	for k, v := range events {
		if isLastStart == (v.State == start) {
			lastId = k
		} else if lastId != -1 {
			if lastId < v.Id {

			} else {

			}
		} else {
			continue
		}
	}*/
	countToday = calculateToday(events)
	return
}

func calculateToday(events []Event) (countToday string) {
	count := time.Now()
	for _, v := range events {
		// 2006 - longYear, 1 - stdNumMonth, 2 - stdDay
		moment, err := time.Parse("2006-01-02T15:04:05.99999Z07:00", v.Moment)
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println(moment.Sub(time.Now()).Nanoseconds())
		if time.Now().Sub(moment).Nanoseconds() > 0 {
			fmt.Println("Got today events")
		}
	}
	// 15 - Hours, 4 - Minutes, 5 - Seconds
	countToday = count.Format("15:4:5")
	return
}
