package main

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
