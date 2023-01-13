package main_test

import (
	. " shared-projects/exercise-go-fundamentals"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateParkingBill(t *testing.T) {
	type args struct {
		vehicleType  VehicleType
		parkDuration int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "should return 3000 when motorcycle parked for 1hr",
			args: args{
				vehicleType:  Motorcycle,
				parkDuration: 1,
			},
			want: 3000,
		},
		{
			name: "should return 5000 when motorcycle parked for 2hr",
			args: args{
				vehicleType:  Motorcycle,
				parkDuration: 2,
			},
			want: 5000,
		},
		{
			name: "should return 71000 when motorcycle parked for 25hr",
			args: args{
				vehicleType:  Motorcycle,
				parkDuration: 25,
			},
			want: 71000,
		},
		{
			name: "should return 7000 when car parked for 1hr",
			args: args{
				vehicleType:  Car,
				parkDuration: 1,
			},
			want: 7000,
		},
		{
			name: "should return 12000 when car parked for 2hr",
			args: args{
				vehicleType:  Car,
				parkDuration: 2,
			},
			want: 12000,
		},
		{
			name: "should return 177000 when car parked for 25hr",
			args: args{
				vehicleType:  Car,
				parkDuration: 25,
			},
			want: 177000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CalculateParkingBill(tt.args.vehicleType, tt.args.parkDuration))
		})
	}
}
