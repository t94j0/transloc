package transloc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// RoutesURL is the transloc route API
const RoutesURL = "http://feeds.transloc.com/3/routes?agencies=%d"

// Routes describes all the routes for the target agency
type Routes struct {
	Agency int
	Routes map[int]Route `json:"routes"`
}

// NewRoutes creates a new routes list and populates the routes
func NewRoutes(agency int) (*Routes, error) {
	newRoute := &Routes{
		Agency: agency,
		Routes: make(map[int]Route),
	}

	if err := newRoute.refresh(); err != nil {
		return newRoute, err
	}

	return newRoute, nil
}

func (r *Routes) generateRouteURL() string {
	return fmt.Sprintf(RoutesURL, r.Agency)
}

func (r *Routes) refresh() error {
	resp, err := http.Get(r.generateRouteURL())
	if err != nil {
		return errors.Wrap(err, "error refreshing routes")
	}

	var routes struct {
		Routes []Route `json:"routes"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&routes); err != nil {
		return err
	}

	for _, route := range routes.Routes {
		r.Routes[route.ID] = route
	}

	return nil
}

// Refresh populates the routes.
func (r *Routes) Refresh() error {
	return r.refresh()
}

// Get gets a target routeID
func (r *Routes) Get(routeID int) Route {
	return r.Routes[routeID]
}
