package internal

type Artist struct {
	Id           int      `json:"id"`
	Img          string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type indexLocation struct {
	Index []location `json:"index"`
}

type location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type indexDates struct {
	Index []date `json:"index"`
}

type date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type indexRelations struct {
	Index []relations `json:"index"`
}

type relations struct {
	Data map[string][]string `json:"datesLocations"`
}
