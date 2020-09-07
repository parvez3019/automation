package main

import (
	. "HotelAutomation/controller"
	. "HotelAutomation/service"
)

func main() {
	hotelService := NewHotelService()
	powerControllerService := NewPowerControllerService(hotelService)
	app := NewPowerAutomationController(hotelService, powerControllerService, NewMotionController())
	hotelConfig := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	app.Init(hotelConfig)
}
