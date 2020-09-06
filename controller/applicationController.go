package controller

import (
	. "HotelAutomation/service"
	"fmt"
)

type ApplicationController struct {
	hotelService *HotelService
}

func NewApplicationController(service *HotelService) *ApplicationController {
	return &ApplicationController{hotelService: service}
}

func (app *ApplicationController) Runner(request CreateHotelRequest) {
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)
	fmt.Print(hotelService.PrintHotelAppliancesState())
}

func (app *ApplicationController) printHotelAppliancesState() {

}
