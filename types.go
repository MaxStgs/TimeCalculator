package main

type Error int

const (
	unknownError Error = 0
	console      Error = 1
	file         Error = 2
)

type Lang int

const (
	en Lang = -1
	ru Lang = 0
	de Lang = 1
)

type EventState string

const (
	unknownState EventState = "U"
	start        EventState = "S"
	end          EventState = "E"
)

type EventType string

const (
	unknownEventType  EventType = "U"
	creatureEventType EventType = "C"
	workEventType     EventType = "W"
)
