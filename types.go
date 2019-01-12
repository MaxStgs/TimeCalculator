package main

import "github.com/getlantern/systray"

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
	UnknownState EventState = "U"
	StartState   EventState = "S"
	EndState     EventState = "E"
)

type EventType string

const (
	unknownEventType  EventType = "U"
	creatureEventType EventType = "C"
	workEventType     EventType = "W"
)

type timerButton struct {
	timerId int
	button  *systray.MenuItem
}

type DebugLevelType int

const (
	DebugLevelNoDebug DebugLevelType = 0
	DebugLevelMinimal DebugLevelType = 1
	DebugLevelNormal  DebugLevelType = 2
	DebugLevelMuch    DebugLevelType = 3
	DebugLevelFull    DebugLevelType = 4
)
