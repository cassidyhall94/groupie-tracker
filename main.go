/*
DROP DOWN LIST
MAKE NAME A LINK LIKE ON LAST EXAMPLE
REMOVE THE BUTTONS AND INSTEAD HAVE THE LINK
STYLE PAGES
MAKE BACKGROUND STAY WHEN SCROLLING
ADD A TEST FILE
CREATE DOMAIN FOR SITE TO BE HOSTED
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

type MyArtistFull struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      []string            `json:"locations"`
	ConcertDates   []string            `json:"concertDates"`
	DatesLocations map[string][]string `json:"datesLocations"`
	WikiLink       []string
}

type MyArtist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type MyLocation struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type MyRelation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type MyDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type MyDates struct {
	Index []MyDate `json:"index"`
}

type MyLocations struct {
	Index []MyLocation `json:"index"`
}

type MyRelations struct {
	Index []MyRelation `json:"index"`
}

type MemberWikiLinks struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func main() {
	// static folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/concert", concertPage)
	http.HandleFunc("/tour", tourPage)
	// http.HandleFunc("/about", aboutPage)

	port := ":8080"
	fmt.Println("Server listen on port localhost:8080")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}
