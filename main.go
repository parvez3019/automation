package main

import (
	. "HotelAutomation/IO"
	. "HotelAutomation/runner"
)

func main() {
	consoleReader := NewConsoleReader()
	NewApplication(consoleReader).Run()
}

//Movement in Floor 1, Sub corridor 2
//No movement in Floor 1, Sub corridor 2 for a minute