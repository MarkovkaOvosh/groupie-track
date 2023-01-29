package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ArtistUrl string = "https://groupietrackers.herokuapp.com/api"

func GetArtist() []Artist {
	r, err := http.Get(ArtistUrl + "/artists")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	var list []Artist

	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		return nil
	}

	return list
}

func GetLocation() indexLocation {
	var list indexLocation

	r, err := http.Get(ArtistUrl + "/locations")
	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		fmt.Println(err)
	}

	return list
}

func GetDates() indexDates {
	var list indexDates

	r, err := http.Get(ArtistUrl + "/dates")
	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		fmt.Println(err)
	}

	return list
}

func GetRelations() indexRelations {
	var list indexRelations

	r, err := http.Get(ArtistUrl + "/relation")
	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		fmt.Println(err)
	}

	return list
}
