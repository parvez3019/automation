package main

import . "HotelAutomation/IO"

func main() {
	consoleReader := NewConsoleReader()
	NewApplicationRunner().Run(consoleReader)
}
