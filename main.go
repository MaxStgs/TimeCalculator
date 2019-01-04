package main

import (
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	InitDB()
	initTray()

	go RunWebServer()
	go RunSocketServer()

	go trayUpdate()
	go trayMenuHandler()
}

func onExit() {

}
