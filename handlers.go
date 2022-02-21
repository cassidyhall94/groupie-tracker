package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	main := r.FormValue("main")
	search := r.FormValue("search")
	filterByCreationFrom := r.FormValue("startCD")
	filterByCreationTill := r.FormValue("endCD")
	filterByFA := r.FormValue("startFA")
	filterByFAend := r.FormValue("endFA")

	data := []MyArtistFull{}
	if !(search == "") {
		data = Search(search)
		fmt.Println("search")
	} else {
		data = Search("a")
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
		fmt.Printf("index.html, error: %+v", err)
		handle500(err, w)
		return
	}

	if main == "Main Page" {
		data = Search("a")
		fmt.Println("main page")
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Printf("Execute(w, data) (%v) error: %+v", data, err)
		handle500(err, w)
		return
	}
}

func concertPage(w http.ResponseWriter, r *http.Request) {
	listOfIds := r.URL.Query()["id"]
	id, err := strconv.Atoi(listOfIds[0])
	if err != nil {
		fmt.Printf("Atoi(listOfIds[0])(%s) error: %+v", listOfIds, err)
		handle400(err, w)
		return
	}

	ArtistsFull, _, _, _, _, _, err := GetData()
	if err != nil || len(ArtistsFull) == 0 {
		if err == nil {
			err = errors.New("empty ArtistsFull from GetData")
		}
		fmt.Printf("GetData() error: %+v", err)
		handle500(err, w)
		return
	}
	artist, err := GetFullDataById(id, ArtistsFull)
	if err != nil {
		fmt.Printf("GetFullDataByID(%d) error: %+v", id, err)
		handle400(err, w)
		return
	}

	tmpl, err := template.ParseFiles("concert.html")
	if err != nil {
		fmt.Printf("concert.html, error: %+v", err)
		handle500(err, w)
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		fmt.Printf("Execute(w, artist) (%v) error: %+v", artist, err)
		handle500(err, w)
		return
	}
}

func tourPage(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("tour")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Atoi(idStr)(%s) error: %+v", idStr, err)
		handle400(err, w)
		return
	}

	ArtistsFull, _, _, _, _, _, err := GetData()
	if err != nil || len(ArtistsFull) == 0 {
		if err == nil {
			err = errors.New("empty ArtistsFull from GetData")
		}
		fmt.Printf("GetData() error: %+v", err)
		handle500(err, w)
		return
	}
	artist, err := GetFullDataById(id, ArtistsFull)
	if err != nil {
		fmt.Printf("GetFullDataByID(%d) error: %+v", id, err)
		handle400(err, w)
		return
	}

	tmpl, err := template.ParseFiles("tour.html")
	if err != nil {
		fmt.Printf("tour.html, error: %+v", err)
		handle500(err, w)
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		fmt.Printf("Execute(w, artist) (%v) error: %+v", artist, err)
		handle500(err, w)
		return
	}
}
