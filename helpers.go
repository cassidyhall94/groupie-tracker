package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ConverterStructToString() ([]string, []MyArtistFull, error) {
	var data []string
	ArtistsFull, Artists, Locations, Dates, _, _, _, err := GetData()
	if err != nil || len(ArtistsFull) == 0 {
		if err == nil {
			err = errors.New("empty ArtistsFull from GetData")
		}
		fmt.Printf("GetData() error: %+v", err)
		return nil, []MyArtistFull{}, err
	}
	for i := 1; i <= len(ArtistsFull); i++ {
		artist, err1 := GetArtistByID(i, Artists)
		locations, err2 := GetLocationByID(i, Locations)
		date, err3 := GetDateByID(i, Dates)
		if err1 != nil || err2 != nil || err3 != nil {
			return data, []MyArtistFull{}, errors.New("error by converter")
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
	return data, ArtistsFull, nil
}

func Search(search string) []MyArtistFull {
	ArtistsFull := []MyArtistFull{}
	if search == "" {
		return ArtistsFull
	}
	art, ArtistsFull, err := ConverterStructToString()
	if err != nil {
		fmt.Printf("ConverterStructToString: (AF: %+v)", ArtistsFull)
		log.Fatal(errors.New("error by converter"))
	}
	search_artist := []MyArtistFull{}
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
					band, err := GetFullDataByID(i+1, ArtistsFull)
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

func handle500(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["500"] = "Internal Server Error"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func handle400(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["400"] = "Bad Request"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
