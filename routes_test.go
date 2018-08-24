package transloc

import (
	"fmt"
	"testing"
)

func ExampleRoutes_refresh() {
	routes, _ := NewRoutes(639)
	highpointe := routes.Get(4008304)
	fmt.Println(highpointe.Active)
}

func ExampleVehicle_vehicles() {
	// Get route information
	routes, _ := NewRoutes(639)
	highpointe := routes.Get(4008304)
	highpointe.RefreshVehicles()

	clemson := Coordinates{34.67631, -82.82393}
	if highpointe.Vehicles.Vehicles[0].Position == clemson {
		fmt.Println("It's here!")
	}
}

func TestRoutes_refresh(t *testing.T) {
	routes, err := NewRoutes(639)
	if err != nil {
		t.Error(err)
	}

	if len(routes.Routes) == 0 {
		t.Error("Displaying zero Clemson busses")
	}

	t.Log("Routes:", len(routes.Routes))
}

func TestRoutes_refreshvehicles(t *testing.T) {
	routes, err := NewRoutes(639)
	if err != nil {
		t.Error(err)
	}

	highpointe := routes.Get(4008304)
	if err := highpointe.RefreshVehicles(); err != nil {
		t.Error(err)
	}

	t.Log(highpointe.Vehicles)
}
