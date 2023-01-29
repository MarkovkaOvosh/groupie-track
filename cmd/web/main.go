package main

import (
	"fmt"
	"groupie-track/cmd/internal"
)

func main() {
	// r := internal.GetArtist()
	// fmt.Println(r[0])
	// l := internal.GetLocation()
	// d := internal.GetDates()
	rel := internal.GetRelations()
	fmt.Println(rel.Index[0])
}
