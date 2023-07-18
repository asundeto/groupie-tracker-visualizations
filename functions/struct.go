package yinyang

import "net/http"

var artistsUrl string = "https://groupietrackers.herokuapp.com/api/artists"

var client *http.Client

type Artist struct {
	Id             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	Locations      string   `json:"locations"`
	ConcertDates   string   `json:"ConcertDates"`
	Relations      string   `json:"relations"`
	DatesLocations interface{}
	DatesLocat     string
}

type datesLocations struct {
	Id             int `json:"id"`
	DatesLocations interface{}
}

type Relat struct {
	IndexRels []IndexRelat `json:"index"`
}

type IndexRelat struct {
	Id       int                 `json:"id"`
	DateLocs map[string][]string `json:"datesLocations"`
}
