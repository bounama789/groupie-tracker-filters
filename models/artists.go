package models

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Relations    string   `json:"relations"`
}

type ArtistResponse struct {
	Id           int       `json:"id"`
	ImageURL     string    `json:"image"`
	Name         string    `json:"name"`
	Members      []string  `json:"members"`
	CreationDate int       `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	Locations    Location  `json:"locations"`
	ConcertDates  Dates     `json:"concertDates"`
	Relations    Relations `json:"relations"`
	TotalConcerts int `json:"totalConcerts"`
}
