package main

type AppData struct {
	Title        string
	Tooltip      string
	Icon         []byte
	IconLocation string
}

type Timer struct {
	Id   int
	Name string
}

type Event struct {
	Id     int
	Moment string
	//Moment			time.Time
	TimerId int
	State   EventState
}
