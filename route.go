package transloc

import "github.com/pkg/errors"

// Route describes a route
type Route struct {
	// ID is the route ID used to identify the route
	ID int `json:"id"`
	// Agency is the ID to identify the agency using TRANSLOC
	Agency int `json:"agency_id"`
	// Name is the human-readable name for the route
	Name string `json:"long_name"`
	// Active describes if the route is active or not
	Active bool `json:"is_active"`
	// Type is the type (bus, etc)
	Type string `json:"type"`
	// Vehicles are the vehicles on that route
	Vehicles *Vehicles
}

// NewRoute creates a new route which is described by an agency and a route ID
func NewRoute(agency, id int) (*Route, error) {
	r := &Route{
		ID:     id,
		Agency: agency,
	}
	if err := r.refreshVehicles(); err != nil {
		return r, err
	}

	return r, nil
}

// RefreshVehicles populates the vehicles in the route. This must be called.
func (r *Route) RefreshVehicles() error {
	return r.refreshVehicles()
}

func (r *Route) refreshVehicles() error {
	vehicles, err := NewVehicles(r.Agency)
	if err != nil {
		return errors.Wrap(err, "error refreshing vehicles")
	}

	routeVehicles := vehicles.GetRoute(r.ID)

	r.Vehicles = &Vehicles{Vehicles: routeVehicles}

	return nil
}
