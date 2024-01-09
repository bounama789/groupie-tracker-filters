package models

type Location struct {
	Id        int    `json:"id"`
	Locations []string `json:"locations"`
	Dates string `json:"dates"`
}

type AllLocations struct {
	Index []Location
}
