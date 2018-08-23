package main

// Coordinates describes a coordinate. The first coordinate is the latitude and the second coordinate is the longitude
type Coordinates [2]float64

// Vehicle describes a single vehicle
type Vehicle struct {
	ID        int         `json:"id"`
	Agency    int         `json:"agency_id"`
	SegmentID int         `json:"segment_id"`
	RouteID   int         `json:"route_id"`
	CallName  string      `json:"call_name"`
	Position  Coordinates `json:"position"`
	Heading   int         `json:"heading"`
	Speed     float64     `json:"speed"`
	Load      float64     `json:"load"`
}
