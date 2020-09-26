package controller

import (
	mocks "HotelAutomation/_mocks"
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	. "HotelAutomation/service"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestShouldCreateHotelAndRegisterDevicesInPowerController(t *testing.T) {
	mockHotelService, mockPCService, mController, paController := setupMockServices()
	request := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}

	mockHotelService.On("CreateHotel", request).Return()
	mockPCService.On("RegisterDevices").Return()

	mockHotelService.CreateHotel(request)
	paController.Init(10*time.Second, make(chan bool))

	assert.Equal(t, 1, len(mockHotelService.Calls))
	assert.Equal(t, 1, len(mockPCService.Calls))
	assert.Len(t, mController.subscribers, 1)
}

func TestShouldToggleNotCheckThresholdPowerInCaseOfErrorInUpdate(t *testing.T) {
	mockHotelService, mockPowerControllerService, motionController, paController := setupMockServices()
	location := CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}

	toggleLightBulbReq := ToggleApplianceRequest{AppType: "Light", TurnOn: true, Location: location}

	mockPowerControllerService.On("Update", toggleLightBulbReq).Return(errors.New("NotFound"))

	err := paController.Update(MovementDetectedEvent{Movement: true, Location: location})

	assert.EqualError(t, err, "NotFound")
	assert.Equal(t, 0, len(mockHotelService.Calls))
	assert.Equal(t, 1, len(mockPowerControllerService.Calls))
	assert.Len(t, motionController.subscribers, 1)
}

func TestShouldToggleOnLightAtLocationWhenMovementDetectedAndKeepACOnIfPowerConsumptionLessThanThreshold(t *testing.T) {
	mockHotelService, mockPowerControllerService, motionController, paController := setupMockServices()
	location := CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}

	toggleLightBulbReq := ToggleApplianceRequest{AppType: "Light", TurnOn: true, Location: location}

	mockPowerControllerService.On("Update", toggleLightBulbReq).Return(nil)
	mockPowerControllerService.On("ToggleApplianceToReverseState", 1, SUB, AC, true).Return()
	mockPowerControllerService.On("TotalPowerConsumptionAtFloor", 1).Return(10)
	mockHotelService.On("GetNumberOfCorridors", 1, MAIN).Return(1)
	mockHotelService.On("GetNumberOfCorridors", 1, SUB).Return(1)

	err := paController.Update(MovementDetectedEvent{Movement: true, Location: location})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(mockHotelService.Calls))
	assert.Equal(t, 3, len(mockPowerControllerService.Calls))
	assert.Len(t, motionController.subscribers, 1)
}

func TestShouldToggleOnLightAtLocationWhenMovementDetectedAndTurnSubCorridorACIfPowerConsumptionMoreThanThreshold(t *testing.T) {
	mockHotelService, mockPowerControllerService, motionController, paController := setupMockServices()
	location := CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}

	toggleLightBulbReq := ToggleApplianceRequest{AppType: "Light", TurnOn: true, Location: location}

	mockPowerControllerService.On("Update", toggleLightBulbReq).Return(nil)
	mockPowerControllerService.On("ToggleApplianceToReverseState", 1, SUB, AC, true).Return()
	mockPowerControllerService.On("ToggleApplianceToReverseState", 1, SUB, AC, false).Return()
	mockPowerControllerService.On("TotalPowerConsumptionAtFloor", 1).Return(50)
	mockHotelService.On("GetNumberOfCorridors", 1, MAIN).Return(1)
	mockHotelService.On("GetNumberOfCorridors", 1, SUB).Return(1)

	err := paController.Update(MovementDetectedEvent{Movement: true, Location: location})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(mockHotelService.Calls))
	assert.Equal(t, 4, len(mockPowerControllerService.Calls))
	assert.Len(t, motionController.subscribers, 1)
}

func TestShouldTurnAcBackOnIfPowerConsumptionGoesLesserThanThreshold(t *testing.T) {
	mockHotelService, mockPowerControllerService, motionController, paController := setupMockServices()
	location := CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}

	toggleLightBulbReq := ToggleApplianceRequest{AppType: "Light", TurnOn: false, Location: location}

	mockPowerControllerService.On("Update", toggleLightBulbReq).Return(nil)
	mockPowerControllerService.On("ToggleApplianceToReverseState", 1, SUB, AC, true).Return()
	mockPowerControllerService.On("TotalPowerConsumptionAtFloor", 1).Return(10)
	mockHotelService.On("GetNumberOfCorridors", 1, MAIN).Return(1)
	mockHotelService.On("GetNumberOfCorridors", 1, SUB).Return(1)

	err := paController.Update(MovementDetectedEvent{Movement: false, Location: location})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(mockHotelService.Calls))
	assert.Equal(t, 3, len(mockPowerControllerService.Calls))
	assert.Len(t, motionController.subscribers, 1)
}

func setupMockServices() (*mocks.HotelServiceI, *mocks.PowerControllerServiceI, *MotionController, *PowerAutomationController) {
	hotelService := &mocks.HotelServiceI{}
	powerControllerService := &mocks.PowerControllerServiceI{}
	paController := NewPowerAutomationController(hotelService, powerControllerService)
	motionController := NewMotionController()
	motionController.AddSubscriber(paController)
	return hotelService, powerControllerService, motionController, paController
}
