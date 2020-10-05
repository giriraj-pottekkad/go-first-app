package structs

import (
	"fmt"
)

//Vehicle exported
type Vehicle struct {
	regNo       int
	vehicleType string
}

//Car exported
type Car struct {
	Vehicle
	model             string
	hasFourWheelDrive bool
	isFullOption      bool
	engine            string
}

//DisplayVehicles exported
func DisplayVehicles() Car {
	car := Car{}
	car.regNo = 8775675
	car.vehicleType = "sedan"
	car.model = "Toyota Etios"
	car.hasFourWheelDrive = false
	car.isFullOption = true

	if car.isFullOption && car.hasFourWheelDrive {
		fmt.Println("This is a 4 wheel drive full option car")
	} else if car.isFullOption && !car.hasFourWheelDrive {
		fmt.Println("This is a full option car")
	}

	return car
}
