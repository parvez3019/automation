package controller

import (
	mock "HotelAutomation/_mocks"
	. "HotelAutomation/model"
	. "HotelAutomation/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateHotelAndRegisterDevicesInPowerController(t *testing.T) {
	mockHotelService, mockPCService, mController, paController := setupMockServices()
	request := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}

	mockHotelService.On("CreateHotel", request).Return()
	mockPCService.On("RegisterDevices").Return()

	mockHotelService.CreateHotel(request)
	paController.Init()

	assert.Equal(t, 1, len(mockHotelService.Calls))
	assert.Equal(t, 1, len(mockPCService.Calls))
	assert.Len(t, mController.subscribers, 1)
}

func TestShouldToggleOnLightAtLocationWhenMovementDetectedAndKeepACOnIfPowerConsumptionLessThanThreshold(t *testing.T) {
	mockHotelService, mockPowerControllerService, motionController, paController := setupMockServices()
	location := CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}

	toggleLightBulbReq := ToggleApplianceRequest{ApplianceType: "Light", SwitchOn: true, Location: location}
	toggleACReq := ToggleApplianceRequest{ApplianceType: "AC", SwitchOn: true, Location: location}

	mockPowerControllerService.On("Update", toggleLightBulbReq).Return(nil)
	mockPowerControllerService.On("Update", toggleACReq).Return(nil)
	mockPowerControllerService.On("TotalPowerConsumptionAtFloor", 1).Return(10)
	mockHotelService.On("GetNumberOfCorridors", MAIN).Return(1)
	mockHotelService.On("GetNumberOfCorridors", SUB).Return(1)

	paController.Update(MovementDetectedEvent{Movement: true, Location: location})

	assert.Equal(t, 2, len(mockHotelService.Calls))
	assert.Equal(t, 3, len(mockPowerControllerService.Calls))
	assert.Len(t, motionController.subscribers, 1)
}

func TestShouldToggleOnLightAtLocationWhenMovementDetectedAndTurnSubCorridorACIfPowerConsumptionMoreThanThreshold(t *testing.T) {
	mockHotelService, mockPowerControllerService, motionController, paController := setupMockServices()
	location := CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}

	toggleLightBulbReq := ToggleApplianceRequest{ApplianceType: "Light", SwitchOn: true, Location: location}
	toggleACReq := ToggleApplianceRequest{ApplianceType: "AC", SwitchOn: false, Location: location}

	mockPowerControllerService.On("Update", toggleLightBulbReq).Return(nil)
	mockPowerControllerService.On("Update", toggleACReq).Return(nil)
	mockPowerControllerService.On("TotalPowerConsumptionAtFloor", 1).Return(50)
	mockHotelService.On("GetNumberOfCorridors", MAIN).Return(1)
	mockHotelService.On("GetNumberOfCorridors", SUB).Return(1)

	paController.Update(MovementDetectedEvent{Movement: true, Location: location})

	assert.Equal(t, 2, len(mockHotelService.Calls))
	assert.Equal(t, 3, len(mockPowerControllerService.Calls))
	assert.Len(t, motionController.subscribers, 1)
}

func TestShouldTurnAcBackOnIfPowerConsumptionGoesLesserThanThreshold(t *testing.T) {
	mockHotelService, mockPowerControllerService, motionController, paController := setupMockServices()
	location := CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}

	toggleLightBulbReq := ToggleApplianceRequest{ApplianceType: "Light", SwitchOn: false, Location: location}
	toggleACReq := ToggleApplianceRequest{ApplianceType: "AC", SwitchOn: true, Location: location}

	mockPowerControllerService.On("Update", toggleLightBulbReq).Return(nil)
	mockPowerControllerService.On("Update", toggleACReq).Return(nil)
	mockPowerControllerService.On("TotalPowerConsumptionAtFloor", 1).Return(10)
	mockHotelService.On("GetNumberOfCorridors", MAIN).Return(1)
	mockHotelService.On("GetNumberOfCorridors", SUB).Return(1)

	paController.Update(MovementDetectedEvent{Movement: false, Location: location})

	assert.Equal(t, 2, len(mockHotelService.Calls))
	assert.Equal(t, 3, len(mockPowerControllerService.Calls))
	assert.Len(t, motionController.subscribers, 1)
}

func setupMockServices() (*mock.HotelServiceI, *mock.PowerControllerServiceI, *MotionController, *PowerAutomationController) {
	hotelService := &mock.HotelServiceI{}
	powerControllerService := &mock.PowerControllerServiceI{}
	paController := NewPowerAutomationController(hotelService, powerControllerService)
	motionController := NewMotionController()
	motionController.AddSubscriber(paController)
	return hotelService, powerControllerService, motionController, paController
}
