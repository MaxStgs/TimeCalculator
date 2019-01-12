package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type IndexData struct {
	Timers []Timer
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
		indexData.Timers = append(indexData.Timers, v)
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

func settingsSaveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(302)
}

func RunWebServer() {
	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler)
	routes.HandleFunc("/api/getTimeToday", handlerGetTimeToday)
	routes.HandleFunc("/menu", indexHandler)
	routes.HandleFunc("/settingsSave", settingsSaveHandler).Methods("POST")
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

func handlerGetTimeToday(w http.ResponseWriter, r *http.Request) {
	idTimer := r.FormValue("idTimer")
	if idTimer == "" {
		http.Error(w, "Not found param idTimer", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idTimer)
	if err != nil {
		http.Error(w, "Server error", http.StatusBadRequest)
		fmt.Println("Can't convert idTimer from " + idTimer + " to int" + err.Error())
		return
	}
	calcTime := calculateDates(Timer{id, ""})
	_, err = w.Write([]byte(calcTime))
	if err != nil {
		http.Error(w, "Server error", http.StatusBadRequest)
		fmt.Println("Can't sent response to client: " + calcTime + " error: " + err.Error())
		return
	}
	return
}

func calculateDates(timer Timer) (countToday string) {
	events := LoadEventsByTimerId(timer.Id)
	if events == nil {
		return
	}
	countToday = calculateToday(events)
	return
}

func calculateToday(events []Event) (countToday string) {
	Count := 0.0
	var CountTime time.Time
	NextEvent := StartState
	var LastEventTime time.Time
	for _, v := range events {
		// 2006 - longYear, 1 - stdNumMonth, 2 - stdDay
		Moment, err := time.Parse("2006-01-02T15:04:05.99999Z07:00", v.Moment)
		if err != nil {
			fmt.Println(err.Error())
		}

		result := time.Now().Sub(Moment)
		if result > 0 {
			if err != nil {
				fmt.Print(err.Error())
				continue
			}

			// Today event handler
			if v.State == NextEvent {
				if NextEvent == StartState {
					NextEvent = EndState
					LastEventTime = Moment
				} else {
					NextEvent = StartState
					diff := Moment.Sub(LastEventTime)
					CountTime = CountTime.Add(diff)
					Count += diff.Seconds()
					if DebugLevel == DebugLevelFull {
						fmt.Printf("Added %f seconds\n", diff.Seconds())
					}
				}
			} else {
				continue
			}
		}
	}
	countToday = CountTime.String()
	fmt.Printf("Today timer: %f\n", Count)
	fmt.Println("Today returned:", CountTime)
	return
}
