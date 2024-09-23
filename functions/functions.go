package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relations struct {
	Index []SubRel `json:"index"`
}

type SubRel struct {
	Id            int                 `json:"id"`
	DateLocations map[string][]string `json:"datesLocations"`
}

type Locations struct {
	Index []SubLocal `json:"index"`
}

type SubLocal struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

type Dates struct {
	Index []SubDate `json:"index"`
}
type SubDate struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Result struct {
	Geometry Geometry `json:"geometry"`
}

type Response struct {
	Results []Result `json:"results"`
}

var (
	artists []Artist
	rel     Relations
	locals  Locations
	dat     Dates
	Cord    Response
)

func GetArtistData() {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error in getting the data from the artist link:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading the body:", err)
		return
	}

	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Println("Error with JSON unmarshal:", err)
		return
	}
}

func GetRelationData() {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error in getting the data from the artist link:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading the body:", err)
		return
	}

	err = json.Unmarshal(body, &rel)
	if err != nil {
		fmt.Println("Error with JSON unmarshal:", err)
		return
	}
}

func GetLocationData() {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("Error in getting the data from the artist link:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading the body:", err)
		return
	}

	err = json.Unmarshal(body, &locals)
	if err != nil {
		fmt.Println("Error with JSON unmarshal:", err)
		return
	}
}

func GetDatesData() {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("Error in getting the data from the artist link:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading the body:", err)
		return
	}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		fmt.Println("Error with JSON unmarshal:", err)
		return
	}
}

func GetCord(localisation string) {

	res, err := http.Get(fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=AIzaSyDx9yeMTOVMGX9o4z_xaHrTq87QbQ3bvD0", localisation))
	if err != nil {
		fmt.Println("Error in getting the data from the artist link:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading the body:", err)
		return
	}

	err = json.Unmarshal(body, &Cord)
	if err != nil {
		fmt.Println("Error with JSON unmarshal:", err)
		return
	}

}

func FitchAllData() {
	GetArtistData()
	GetRelationData()
	GetLocationData()
	GetDatesData()
}
