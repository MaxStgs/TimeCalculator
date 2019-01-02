package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/pkg/browser"
	"io/ioutil"
)

var generalPage, myselfStart, myselfStop, exit *systray.MenuItem

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	firstInit()
	initTrayMenu()
	enableTrayMenu()

	go RunWebServer()
	go update()
	go updateClick()
}

func initTrayMenu() {
	generalPage = systray.AddMenuItem(getText("General page"), getText("General page"))
	generalPage.Hide()

	myselfStart = systray.AddMenuItem(getText("Start myself timer"), getText("Myself timer start"))
	myselfStart.Hide()

	myselfStop = systray.AddMenuItem(getText("Stop myself timer"), getText("Myself timer stop"))
	myselfStop.Hide()

	systray.AddSeparator()

	exit = systray.AddMenuItem(getText("Exit"), getText("Exit"))
}

func enableTrayMenu() {
	generalPage.Show()
	myselfStart.Show()
}

func firstInit() {
	systray.SetTitle(getText("Time Calculator"))
	systray.SetTooltip(getText("Time Calculator"))
	icon := getIcon("assets/icon.ico")
	if icon == nil {
		fmt.Println("onReady.getIcon is nil")
		return
	}
	systray.SetIcon(icon)
}

func update() {

}

func updateClick() {
	for {
		select {
		case <-generalPage.ClickedCh:
			_ = browser.OpenURL(generateIndex())
		case <-myselfStart.ClickedCh:
			handleMyselfStart()
		case <-myselfStop.ClickedCh:
			handleMyselfStop()
		case <-exit.ClickedCh:
			systray.Quit()
		}
	}
}

func onExit() {

}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		errorHandler("main.getIcon() err: "+err.Error(), console)
	}

	return b
}