package transloc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// VehiclesURL is the transloc api for getting vehicles in an agency
const VehiclesURL = "http://feeds.transloc.com/3/vehicle_statuses?agencies=%d"

// Vehicles describes a group of vehicles
type Vehicles struct {
	Agency   int
	Auccess  bool      `json:"success"`
	Vehicles []Vehicle `json:"vehicles"`
}

// NewVehicles describes a groups of vehicles
func NewVehicles(agency int) *Vehicles {
	return &Vehicles{
		Agency: agency,
	}
}

func (v *Vehicles) generateVehiclesURL() string {
	return fmt.Sprintf(VehiclesURL, v.Agency)
}

// Refresh populates the list of vehicles
func (v *Vehicles) Refresh() error {
	resp, err := http.Get(v.generateVehiclesURL())
	if err != nil {
		return errors.Wrap(err, "error refreshing vehicles")
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}

// GetRoute gets a list of vehicles with the specific route ID
func (v *Vehicles) GetRoute(routeID int) []Vehicle {
	vehicles := make([]Vehicle, 0)

	for _, vehicle := range v.Vehicles {
		if vehicle.RouteID == routeID {
			vehicles = append(vehicles, vehicle)
		}
	}

	return vehicles
}
