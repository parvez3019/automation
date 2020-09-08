package main

import (
	. "HotelAutomation/IO"
	. "HotelAutomation/runner"
)

func main() {
	consoleReader := NewConsoleReader()
	NewApplication().Run(consoleReader)
}
