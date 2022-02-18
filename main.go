/*
DROP DOWN LIST
MAKE NAME A LINK LIKE ON LAST EXAMPLE
REMOVE THE BUTTONS AND INSTEAD HAVE THE LINK
STYLE PAGES
MAKE BACKGROUND STAY WHEN SCROLLING
ADD A TEST FILE
REMOVE UNNECESSERY THINGS FROM MAIN.GO
ADD WIKIPEDIA FOR EACH MEMBER
CREATE DOMAIN FOR SITE TO BE HOSTED
THINK ABOUT SPLITTING THE MAIN.GO FILE INTO TWO: HANDLERS.GO AND DATAGETTERS.GO?
*/
package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

// look into merging myartistfull and myartist structs
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
	// Relations      string              `json:"relations"`
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

var (
	ArtistsFull []MyArtistFull
	Artists     []MyArtist
	Dates       MyDates
	Locations   MyLocations
	Relations   MyRelations
	MemLinks    []MemberWikiLinks
)

func main() {
	// static folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/concert", concertPage)
	http.HandleFunc("/tour", tourPage)

	port := ":8080"
	fmt.Println("Server listen on port localhost", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}

func GetWikiLinks() error {

	csvFile, err := os.Open("members-wiki.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord MemberWikiLinks
	var allRecords []MemberWikiLinks

	for _, each := range csvData {
		oneRecord.Name = each[0]
		oneRecord.Link = each[1]
		allRecords = append(allRecords, oneRecord)
	}

	jsondata, err := json.Marshal(allRecords) // convert to JSON
	json.Unmarshal(jsondata, &MemLinks)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

func GetArtistsData() ([]MyArtist, error) {
	Artists := []MyArtist{}
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return Artists, errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Artists, errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Artists)
	return Artists, nil
}

func GetDatesData() (MyDates, error) {
	Dates := MyDates{}
	resp, err := http.Get(baseURL + "/dates")
	if err != nil {
		return Dates, errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Dates, errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Dates)
	return Dates, nil
}

func GetLocationsData() (MyLocations, error) {
	Locations := MyLocations{}
	resp, err := http.Get(baseURL + "/locations")
	if err != nil {
		return Locations, errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Locations, errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Locations)
	return Locations, nil
}

func GetRelationsData() (MyRelations, error) {
	Relations := MyRelations{}
	resp, err := http.Get(baseURL + "/relation")
	if err != nil {
		return Relations, errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Relations, errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Relations)
	fmt.Println(Relations.Index[0].DatesLocations)
	return Relations, nil
}

func GetData() error {
	if len(ArtistsFull) != 0 {
		return nil
	}
	Artists, err1 := GetArtistsData()
	Locations, err2 := GetLocationsData()
	Dates, err3 := GetDatesData()
	Relations, err4 := GetRelationsData()
	err5 := GetWikiLinks()
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		return errors.New("error by get data artists, locations, dates")
	}
	for i := range Artists {

		var tmpl MyArtistFull
		var addMemLinks []string
		for j := 0; j < len(Artists[i].Members); j++ {
			for m := 0; m < len(MemLinks); m++ {
				if MemLinks[m].Name == Artists[i].Members[j] {
					addMemLinks = append(addMemLinks, MemLinks[m].Link)
				}
			}
		}
		tmpl.ID = i + 1
		tmpl.Image = Artists[i].Image
		tmpl.Name = Artists[i].Name
		tmpl.Members = Artists[i].Members
		tmpl.CreationDate = Artists[i].CreationDate
		tmpl.FirstAlbum = Artists[i].FirstAlbum
		tmpl.Locations = Locations.Index[i].Locations
		tmpl.ConcertDates = Dates.Index[i].Dates
		tmpl.DatesLocations = Relations.Index[i].DatesLocations
		tmpl.WikiLink = addMemLinks
		ArtistsFull = append(ArtistsFull, tmpl)
	}
	return nil
}

func GetArtistByID(id int) (MyArtist, error) {
	for _, artist := range Artists {
		if artist.ID == id {
			return artist, nil
		}
	}
	return MyArtist{}, errors.New("not found")
}

func GetDateByID(id int) (MyDate, error) {
	for _, date := range Dates.Index {
		if date.ID == id {
			return date, nil
		}
	}
	return MyDate{}, errors.New("not found")
}

func GetLocationByID(id int) (MyLocation, error) {
	for _, location := range Locations.Index {
		if location.ID == id {
			return location, nil
		}
	}
	return MyLocation{}, errors.New("not found")
}

func GetRelationByID(id int) (MyRelation, error) {
	for _, relation := range Relations.Index {
		if relation.ID == id {
			return relation, nil
		}
	}
	return MyRelation{}, errors.New("not found")
}

func GetFullDataById(id int) (MyArtistFull, error) {
	for _, artist := range ArtistsFull {
		if artist.ID == id {
			return artist, nil
		}
	}
	return MyArtistFull{}, errors.New("not found")
}

var data []MyArtistFull

func mainPage(w http.ResponseWriter, r *http.Request) {
	err := GetData()
	if err != nil {
		errors.New("error by get data")
	}

	main := r.FormValue("main")
	search := r.FormValue("search")
	filterByCreationFrom := r.FormValue("startCD")
	filterByCreationTill := r.FormValue("endCD")
	filterByFA := r.FormValue("startFA")
	filterByFAend := r.FormValue("endFA")

	if !(search == "" && len(data) != 0) {
		data = Search(search)
	}

	if filterByCreationFrom != "" || filterByCreationTill != "" {
		if filterByCreationFrom == "" {
			filterByCreationFrom = "1900"
		}
		if filterByCreationTill == "" {
			filterByCreationTill = "2020"
		}

	}

	if filterByFA != "" || filterByFAend != "" {
		if filterByFA == "" {
			filterByFA = "1900-01-01"
		}
		if filterByFAend == "" {
			filterByFAend = "2020-03-03"
		}

	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		handle500(err, w)
		return
	}

	if main == "Main Page" {
		data = Search("a")
	}

	if err := tmpl.Execute(w, data); err != nil {
		handle500(err, w)
		return
	}
}

func concertPage(w http.ResponseWriter, r *http.Request) {
	listOfIds := r.URL.Query()["id"]
	id, err := strconv.Atoi(listOfIds[0])
	if err != nil {
		handle500(err, w)
		return
	}

	artist, err := GetFullDataById(id)
	if err != nil {
		http.Error(w, "Bad Request: 400", 400)
		return
	}

	tmpl, err := template.ParseFiles("concert.html")
	if err != nil {
		handle500(err, w)
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		handle500(err, w)
		return
	}
}

func tourPage(w http.ResponseWriter, r *http.Request) {
	listOfIds := r.URL.Query()["id"]
	id, err := strconv.Atoi(listOfIds[0])
	if err != nil {
		handle500(err, w)
		return
	}

	artist, err := GetFullDataById(id)
	if err != nil {
		http.Error(w, "Bad Request: 400", 400)
		return
	}

	tmpl, err := template.ParseFiles("tour.html")
	if err != nil {
		handle500(err, w)
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		handle500(err, w)
		return
	}
}

// this needs a better name, figure out what it's doing
// if a variable is being used in a function it needs a very
// good reason to not be in the function signature
// sometimes it might feel like a variable goes through 5 functions before being used
// that's a whole other thing but better than global scope
func ConverterStructToString(ArtistsFull []MyArtistFull) ([]string, error) {
	var data []string
	for i := 1; i <= len(ArtistsFull); i++ {
		artist, err1 := GetArtistByID(i)
		locations, err2 := GetLocationByID(i)
		date, err3 := GetDateByID(i)
		if err1 != nil || err2 != nil || err3 != nil {
			return data, errors.New("error by converter")
		}

		str := artist.Name + " "
		for _, member := range artist.Members {
			str += member + " "
		}
		str += strconv.Itoa(artist.CreationDate) + " "
		str += artist.FirstAlbum + " "
		for _, location := range locations.Locations {
			str += location + " "
		}
		for _, d := range date.Dates {
			str += d + " "
		}
		data = append(data, str)
	}
	return data, nil
}

func Search(search string) []MyArtistFull {
	if search == "" {
		return ArtistsFull
	}
	art, err := ConverterStructToString(ArtistsFull)
	if err != nil {
		errors.New("error by converter")
	}
	var search_artist []MyArtistFull

	for i, artist := range art {
		lower_band := strings.ToLower(artist)
		for i_name, l_name := range []byte(lower_band) {
			lower_search := strings.ToLower(search)
			if lower_search[0] == l_name {
				length_name := 0
				indx := i_name
				for _, l := range []byte(lower_search) {
					if l == lower_band[indx] {
						if indx+1 == len(lower_band) {
							break
						}
						indx++
						length_name++
					} else {
						break
					}
				}
				if len(search) == length_name {
					band, err := GetFullDataById(i + 1)
					if err != nil {
						fmt.Println(err)
					}
					search_artist = append(search_artist, band)
					break
				}
			}
		}

	}
	return search_artist
}

func handle500(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["500"] = "Internal Server Error"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
