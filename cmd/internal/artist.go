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
	//
	Data map[string][]string `json:"datesLocations"`
}

type indexRelations struct {
	Index []Relations `json:"index"`
}

type Relations struct {
	Data map[string][]string `json:"datesLocations"`
}

type allInfo struct {
	artists []Artist
	rel     indexRelations
}
