package transloc

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
func NewRoute(agency, id int) *Route {
	return &Route{
		ID:     id,
		Agency: agency,
	}
}

// RefreshVehicles populates the vehicles in the route. This must be called.
// TODO: Call when route is created
func (r *Route) RefreshVehicles() error {
	vehicles := NewVehicles(r.Agency)
	if err := vehicles.Refresh(); err != nil {
		return err
	}

	r.Vehicles = &Vehicles{Vehicles: vehicles.GetRoute(r.ID)}

	return nil
}
