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

// NewRoutes creates a new routes list
func NewRoutes(agency int) *Routes {
	return &Routes{
		Agency: agency,
		Routes: make(map[int]Route),
	}
}

func (r *Routes) generateRouteURL() string {
	return fmt.Sprintf(RoutesURL, r.Agency)
}

// Refresh populates the routes.
// TODO: Add to NewRoutes
func (r *Routes) Refresh() error {
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

// Get gets a target routeID
func (r *Routes) Get(routeID int) Route {
	return r.Routes[routeID]
}
