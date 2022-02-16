/*
INTERNAL SERVER ERROR - 500
CHECK WHY FILTERS DON'T WORK
MAKE SEARCH WORK (COULD LIST THE BANDS BUT MAYBE FILTER BETTER?)
MAKE NAME A LINK LIKE ON LAST EXAMPLE
REMOVE THE BUTTONS AND INSTEAD HAVE THE LINK
STYLE PAGES
HAVE FEW BANDS PER PAGE
MAKE BACKGROUND STAY WHEN SCROLLING
ADD A TEST FILE
REMOVE UNNECESSERY THINGS FROM MAIN.GO
ADD WIKIPEDIA FOR EACH MEMBER
*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

var (
	ArtistsFull []MyArtistFull
	Artists     []MyArtist
	Dates       MyDates
	Locations   MyLocations
	Relations   MyRelations
)

func GetArtistsData() error {
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Artists)
	return nil
}

func GetDatesData() error {
	resp, err := http.Get(baseURL + "/dates")
	if err != nil {
		return errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Dates)
	return nil
}

func GetLocationsData() error {
	resp, err := http.Get(baseURL + "/locations")
	if err != nil {
		return errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Locations)
	return nil
}

func GetRelationsData() error {
	resp, err := http.Get(baseURL + "/relation")
	if err != nil {
		return errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Relations)
	return nil
}

func GetData() error {
	if len(ArtistsFull) != 0 {
		return nil
	}
	err1 := GetArtistsData()
	err2 := GetLocationsData()
	err3 := GetDatesData()
	err4 := GetRelationsData()
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return errors.New("error by get data artists, locations, dates")
	}
	for i := range Artists {
		var tmpl MyArtistFull
		tmpl.ID = i + 1
		tmpl.Image = Artists[i].Image
		tmpl.Name = Artists[i].Name
		tmpl.Members = Artists[i].Members
		tmpl.CreationDate = Artists[i].CreationDate
		tmpl.FirstAlbum = Artists[i].FirstAlbum
		tmpl.Locations = Locations.Index[i].Locations
		tmpl.ConcertDates = Dates.Index[i].Dates
		tmpl.DatesLocations = Relations.Index[i].DatesLocations
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

	var mem1, mem2, mem3, mem4, mem5, mem6, mem7, mem8 int
	mem1, err1 := strconv.Atoi(r.FormValue("mem1"))
	if err1 != nil {
		mem1 = 0
	}
	mem2, err2 := strconv.Atoi(r.FormValue("mem2"))
	if err2 != nil {
		mem2 = 0
	}
	mem3, err3 := strconv.Atoi(r.FormValue("mem3"))
	if err3 != nil {
		mem3 = 0
	}
	mem4, err4 := strconv.Atoi(r.FormValue("mem4"))
	if err4 != nil {
		mem4 = 0
	}
	mem5, err5 := strconv.Atoi(r.FormValue("mem5"))
	if err5 != nil {
		mem5 = 0
	}
	mem6, err6 := strconv.Atoi(r.FormValue("mem6"))
	if err6 != nil {
		mem6 = 0
	}
	mem7, err7 := strconv.Atoi(r.FormValue("mem7"))
	if err7 != nil {
		mem7 = 0
	}
	mem8, err8 := strconv.Atoi(r.FormValue("mem8"))
	if err8 != nil {
		mem8 = 0
	}
	mem := []int{mem1, mem2, mem3, mem4, mem5, mem6, mem7, mem8}
	sum := 0
	for _, n := range mem {
		sum += n
	}
	println("startCD:", filterByCreationFrom)
	fmt.Println("mem:", mem)

	filterByFA := r.FormValue("startFA")
	filterByFAend := r.FormValue("endFA")
	fmt.Println("filterFA:", filterByFA)
	fmt.Println("filterFAend:", filterByFAend)

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
		http.Error(w, err.Error(), 400)
		return
	}

	if main == "Main Page" {
		data = Search("a")
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func concertPage(w http.ResponseWriter, r *http.Request) {
	listOfIds := r.URL.Query()["id"]
	id, _ := strconv.Atoi(listOfIds[0])

	artist, _ := GetFullDataById(id)

	for key, value := range artist.DatesLocations {
		fmt.Print(key + "  - ")
		for _, e := range value {
			println(e)
		}
	}

	tmpl, err := template.ParseFiles("concert.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func tourPage(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("tour")
	id, _ := strconv.Atoi(idStr)
	artist, _ := GetFullDataById(id)

	for key, value := range artist.DatesLocations {
		fmt.Print(key + "  - ")
		for _, e := range value {
			println(e)
		}
	}

	tmpl, err := template.ParseFiles("tour.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func ConverterStructToString() ([]string, error) {
	var data []string
	for i := 1; i <= len(Artists); i++ {
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
	println("Convert to str Done!")
	return data, nil
}

func Search(search string) []MyArtistFull {
	if search == "" {
		return ArtistsFull
	}
	art, err := ConverterStructToString()
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
					band, _ := GetFullDataById(i + 1)
					search_artist = append(search_artist, band)
					break
				}
			}
		}

	}
	println("Search str Done!")
	return search_artist
}

func main() {
	// static folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/concert", concertPage)
	http.HandleFunc("/tour", tourPage)

	port := ":8080"
	println("Server listen on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Listen and Server", err)
	}
}
