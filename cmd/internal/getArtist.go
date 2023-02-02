package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ArtistUrl string = "https://groupietrackers.herokuapp.com/api"

var (
	List    []Artist
	ListRel indexRelations
)

func GetArtists() error {
	r, err := http.Get(ArtistUrl + "/artists")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&List)
}

func GetRelations() error {
	r, err := http.Get(ArtistUrl + "/relation")
	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&ListRel)
}

func GetAllInfo() error {
	if err := GetArtists(); err != nil {
		return err
	}
	if err := GetRelations(); err != nil {
		return err
	}
	return nil
}
