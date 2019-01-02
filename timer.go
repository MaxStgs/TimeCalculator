package main

func handleMyselfStart() {
	myselfStart.Hide()
	myselfStop.Show()
}

func handleMyselfStop() {
	myselfStart.Show()
	myselfStop.Hide()
}

func handleWorkStart() {
	workStart.Hide()
	workStop.Show()
}

func handleWorkStop() {
	workStart.Show()
	workStop.Hide()
}
