package models


type Relations struct {
	Id int    `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}