package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ArtistUrl string = "https://groupietrackers.herokuapp.com/api/artists"

type Artist struct {
	Id           int      `json:"id"`
	Img          string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    locat    `json:"locations"`
	ConcertDates dates    `json:"concertDates"`
	Relations    rel      `json:"relations"`
}

type locat struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type rel struct {
	Id   int
	Data map[string][]string `json:"datesLocations"`
}

func GetArtist() []Artist {
	r, err := http.Get(ArtistUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	var list []Artist

	// if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
	// 	return nil
	// }

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &list)

	return list
}
