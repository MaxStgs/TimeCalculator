package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/pkg/browser"
)

var (
	generalPage,
	myselfStart,
	myselfStop,
	workStart,
	workStop,
	settings,
	exit *systray.MenuItem
)

func initTray() {
	systray.SetTitle(getText("Time Calculator"))
	systray.SetTooltip(getText("Time Calculator"))
	icon := getIcon("assets/icon.ico")
	if icon == nil {
		fmt.Println("onReady.getIcon is nil")
		return
	}
	systray.SetIcon(icon)
}

func initTrayMenu() {
	generalPage = systray.AddMenuItem(getText("General page"), getText("General page"))
	generalPage.Hide()

	myselfStart = systray.AddMenuItem(getText("Start myself timer"), getText("Myself timer start"))
	myselfStart.Hide()

	myselfStop = systray.AddMenuItem(getText("Stop myself timer"), getText("Myself timer stop"))
	myselfStop.Hide()

	myselfStart = systray.AddMenuItem(getText("Start myself timer"), getText("Work timer start"))
	myselfStart.Hide()

	myselfStop = systray.AddMenuItem(getText("Stop myself timer"), getText("Work timer stop"))
	myselfStop.Hide()

	workStart = systray.AddMenuItem(getText("Start work timer"), getText("Work timer start"))
	workStart.Hide()

	workStop = systray.AddMenuItem(getText("Stop work timer"), getText("Work timer stop"))
	workStop.Hide()

	systray.AddSeparator()

	settings = systray.AddMenuItem(getText("Settings"), getText("Settings"))
	exit = systray.AddMenuItem(getText("Exit"), getText("Exit"))
}

func enableTrayMenu() {
	generalPage.Show()
	myselfStart.Show()
	workStart.Show()
}

func trayMenuHandler() {
	for {
		select {
		case <-generalPage.ClickedCh:
			_ = browser.OpenURL(generateIndex())
		case <-myselfStart.ClickedCh:
			handleMyselfStart()
		case <-myselfStop.ClickedCh:
			handleMyselfStop()
		case <-workStart.ClickedCh:
			handleWorkStart()
		case <-workStop.ClickedCh:
			handleWorkStop()
		case <-settings.ClickedCh:
			handleSettings()
		case <-exit.ClickedCh:
			systray.Quit()
		}
	}
}

func trayUpdate() {

}

func handleMyselfStart() {
	myselfStart.Hide()
	myselfStop.Show()
	handleEvent(creatureEventType, start)
}

func handleMyselfStop() {
	myselfStart.Show()
	myselfStop.Hide()
	handleEvent(creatureEventType, end)
}

func handleWorkStart() {
	workStart.Hide()
	workStop.Show()
	handleEvent(workEventType, start)
}

func handleWorkStop() {
	workStart.Show()
	workStop.Hide()
	handleEvent(workEventType, end)
}

func handleSettings() {
	_ = browser.OpenURL(generateSettigs())
}
