package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetTourData() ([]TourData, error) {
	TourThings := []TourData{}
	Relations, err1 := GetRelationsData()

	if err1 != nil {
		log.Fatal(err1)
	}

	f, err := os.Open("web/tour_data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var oneRecord TourData
	var allRecords []TourData

	for _, each := range csvData {
		oneRecord.ArtistID, _ = strconv.Atoi(each[0])
		oneRecord.RelationID = each[1]
		oneRecord.City = each[2]
		oneRecord.Country = each[3]
		oneRecord.TourDates = Relations.Index[oneRecord.ArtistID-1].DatesLocations[each[1]]
		allRecords = append(allRecords, oneRecord)
	}

	jsondata, err := json.Marshal(allRecords) // convert to JSON
	json.Unmarshal(jsondata, &TourThings)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return TourThings, nil
}

func GetWikiLinks() ([]MemberWikiLinks, error) {
	MemLinks := []MemberWikiLinks{}
	csvFile, err := os.Open("web/static/members-wiki.txt")
	if err != nil {
		fmt.Printf("cvsFile error: %+v", err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("csvData error: %+v", err)
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
		fmt.Printf("jsondata error: %+v", err)
		os.Exit(1)
	}
	return MemLinks, nil
}

func GetArtistsData() ([]MyArtist, error) {
	Artists := []MyArtist{}
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return Artists, fmt.Errorf("error by get Artists: %w", err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Artists, fmt.Errorf("error by ReadAll Artists: %w", err)
	}
	json.Unmarshal(bytes, &Artists)
	return Artists, nil
}

func GetDatesData() (MyDates, error) {
	Dates := MyDates{}
	resp, err := http.Get(baseURL + "/dates")
	if err != nil {
		return Dates, fmt.Errorf("error by get Dates: %w", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Dates, fmt.Errorf("error by ReadAll Dates: %w", err)
	}
	json.Unmarshal(bytes, &Dates)
	return Dates, nil
}

func GetLocationsData() (MyLocations, error) {
	Locations := MyLocations{}
	resp, err := http.Get(baseURL + "/locations")
	if err != nil {
		return Locations, fmt.Errorf("error by get Locations: %w", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Locations, fmt.Errorf("error by ReadAll Locations: %w", err)
	}
	json.Unmarshal(bytes, &Locations)
	return Locations, nil
}

func GetRelationsData() (MyRelations, error) {
	Relations := MyRelations{}
	resp, err := http.Get(baseURL + "/relation")
	if err != nil {
		return Relations, fmt.Errorf("error by get Relations: %w", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Relations, fmt.Errorf("error by ReadAll Relations: %w", err)
	}
	json.Unmarshal(bytes, &Relations)
	return Relations, nil
}

func GetData() ([]MyArtistFull, []MyArtist, MyLocations, MyDates, MyRelations, []MemberWikiLinks, []TourData, error) {
	Artists, err1 := GetArtistsData()
	Locations, err2 := GetLocationsData()
	Dates, err3 := GetDatesData()
	Relations, err4 := GetRelationsData()
	MemLinks, err5 := GetWikiLinks()
	TourThings, err6 := GetTourData()
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		return []MyArtistFull{}, []MyArtist{}, MyLocations{}, MyDates{}, MyRelations{}, []MemberWikiLinks{}, []TourData{}, fmt.Errorf("error from get data artists: %v, locations: %v, dates: %v, relations: %v, memlinks: %v, or TourThings: %v", err1, err2, err3, err4, err5, err6)
	}
	ret := []MyArtistFull{}
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

		var addCity []string
		for m := 0; m < len(TourThings); m++ {
			if TourThings[m].ArtistID == Artists[i].ID {
				addCity = append(addCity, TourThings[m].City)
			}
		}

		var addCountry []string
		for m := 0; m < len(TourThings); m++ {
			if TourThings[m].ArtistID == Artists[i].ID {
				addCountry = append(addCountry, TourThings[m].Country)
			}
		}

		var addDates [][]string
		for m := 0; m < len(TourThings); m++ {
			if TourThings[m].ArtistID == Artists[i].ID {
				addDates = append(addDates, TourThings[m].TourDates)
			}
		}

		var addDatesString []string
		for _, date := range addDates {
			s := strings.Join(date, " | ")
			addDatesString = append(addDatesString, s)
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
		tmpl.TourCity = addCity
		tmpl.TourCountry = addCountry
		tmpl.TourDates = addDates
		tmpl.TourDateString = addDatesString
		ret = append(ret, tmpl)
	}
	return ret, Artists, Locations, Dates, Relations, MemLinks, TourThings, nil
}

func GetArtistByID(id int, Artists []MyArtist) (MyArtist, error) {
	for _, artist := range Artists {
		if artist.ID == id {
			return artist, nil
		}
	}
	return MyArtist{}, errors.New("artist not found")
}

func GetDateByID(id int, Dates MyDates) (MyDate, error) {
	for _, date := range Dates.Index {
		if date.ID == id {
			return date, nil
		}
	}
	return MyDate{}, errors.New("date not found")
}

func GetLocationByID(id int, Locations MyLocations) (MyLocation, error) {
	for _, location := range Locations.Index {
		if location.ID == id {
			return location, nil
		}
	}
	return MyLocation{}, errors.New("location not found")
}

func GetRelationByID(id int, Relations MyRelations) (MyRelation, error) {
	for _, relation := range Relations.Index {
		if relation.ID == id {
			return relation, nil
		}
	}
	return MyRelation{}, errors.New("relation not found")
}

func GetFullDataByID(id int, ArtistsFull []MyArtistFull) (MyArtistFull, error) {
	for _, artist := range ArtistsFull {
		if artist.ID == id {
			return artist, nil
		}
	}
	return MyArtistFull{}, errors.New("fulldata not found")
}
