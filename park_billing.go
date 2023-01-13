package main

type VehicleType int

const (
	Car VehicleType = iota
	Motorcycle
)

func CalculateParkingBill(vehicleType VehicleType, duration int) float32 {
	var bill int

	if vehicleType == Motorcycle {
		bill = 3000
		if duration >= 24 {
			bill += (duration / 24) * 20000
		}
		bill += (duration - 1) * 2000
	}

	if vehicleType == Car {
		bill = 7000
		if duration >= 24 {
			bill += (duration / 24) * 50000
		}
		bill += (duration - 1) * 5000
	}
	return float32(bill)
}
