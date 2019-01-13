package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/pkg/browser"
	"strconv"
)

var (
	generalPage,
	settings,
	exit *systray.MenuItem

	buttons = make(map[int][]timerButton, 0)
)

func initTray() {
	appData := LoadAppData()
	/* Now data load from DB
	systray.SetTitle(getText("Time Calculator"))
	systray.SetTooltip(getText("Time Calculator"))
	icon := getIcon("assets/icon.ico")
	if icon == nil {
		fmt.Println("onReady.getIcon is nil")
		return
	}
	systray.SetIcon(icon)
	*/
	systray.SetTitle(appData.Title)
	systray.SetTooltip(appData.Tooltip)
	if appData.Icon == nil {
		icon := getIcon("assets/icon.ico")
		if icon == nil {
			fmt.Println("onReady.getIcon is nil")
			return
		}
		systray.SetIcon(icon)
	} else {
		systray.SetIcon(appData.Icon)
	}

	initTrayMenu()
}

func initTrayMenu() {

	timers := LoadTimers()
	if timers != nil {
		for _, v := range timers {
			timerStartName := getText("Start") + " " + v.Name
			timerStart := systray.AddMenuItem(timerStartName, timerStartName)
			timerEndName := getText("End") + " " + v.Name
			timerEnd := systray.AddMenuItem(timerEndName, timerEndName)
			timerEnd.Hide()
			index := len(buttons)
			buttons[index] = []timerButton{{v.Id, timerStart}, {v.Id, timerEnd}}

			go func() {
				for {
					select {
					case <-timerStart.ClickedCh:
						handleMenuItemClickTimerStart(index)
					case <-timerEnd.ClickedCh:
						handleMenuItemClickTimerStop(index)
					}
				}
			}()
		}
	}
	/*
		myselfStart = systray.AddMenuItem(getText("Start myself timer"), getText("Myself timer StartState"))
		myselfStart.Hide()

		myselfStop = systray.AddMenuItem(getText("Stop myself timer"), getText("Myself timer stop"))
		myselfStop.Hide()
	*/

	systray.AddSeparator()

	generalPage = systray.AddMenuItem(getText("General page"), getText("General page"))
	settings = systray.AddMenuItem(getText("Settings"), getText("Settings"))
	exit = systray.AddMenuItem(getText("Exit"), getText("Exit"))
}

func trayMenuHandler() {
	for {
		select {
		case <-generalPage.ClickedCh:
			_ = browser.OpenURL(generateMenu())
		case <-settings.ClickedCh:
			handleSettings()
		case <-exit.ClickedCh:
			systray.Quit()
		}
	}
}

func trayUpdate() {

}

func handleMenuItemClickTimerStart(index int) {
	handleTimerStart(index)
	sendMessage(Message{-1, "StartTimer", strconv.Itoa(index + 1)}, Users)
}

func handleMenuItemClickTimerStop(index int) {
	sendMessage(Message{-1, "StopTimer", strconv.Itoa(index + 1)}, Users)
	handleTimerStop(index)
}

func handleTimerStart(index int) {
	buttons[index][1].button.Show()
	buttons[index][0].button.Hide()
	handleEvent(buttons[index][0].timerId, StartState)
}

func handleTimerStop(index int) {
	buttons[index][1].button.Hide()
	buttons[index][0].button.Show()
	handleEvent(buttons[index][0].timerId, StopState)
}

func handleSettings() {
	_ = browser.OpenURL(generateSettigs())
}
