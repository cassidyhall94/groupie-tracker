package main

import "testing"

func TestGetWikiLinks(t *testing.T) {
	memlinks, err := GetWikiLinks()
	if err != nil {
		t.Fatalf("memlinks err:%+v", err)
	}
	if len(memlinks) < 1 {
		t.Fatalf("memlinks failed due to length < 1")
	}
}

func TestGetArtistsData(t *testing.T) {
	artists, err := GetArtistsData()
	if err != nil {
		t.Fatalf("artists err:%+v", err)
	}
	if len(artists) < 1 {
		t.Fatalf("artists failed due to length < 1")
	}
}

func TestGetDatesData(t *testing.T) {
	dates, err := GetDatesData()
	if err != nil {
		t.Fatalf("dates err:%+v", err)
	}
	if len(dates.Index) < 1 {
		t.Fatalf("dates failed due to length < 1")
	}
}

func TestGetLocationsData(t *testing.T) {
	locations, err := GetLocationsData()
	if err != nil {
		t.Fatalf("locations err:%+v", err)
	}
	if len(locations.Index) < 1 {
		t.Fatalf("locations failed due to length < 1")
	}
}

func TestGetRelationsData(t *testing.T) {
	relations, err := GetRelationsData()
	if err != nil {
		t.Fatalf("relations err:%+v", err)
	}
	if len(relations.Index) < 1 {
		t.Fatalf("relations failed due to length < 1")
	}
}

type input struct {
	ID        int
	MyArtists []MyArtist
}

func TestGetArtistByID(t *testing.T) {
	testCases := []struct {
		in      input
		want    MyArtist
		wantErr bool
	}{
		{
			in: input{
				ID: 1,
				MyArtists: []MyArtist{
					{
						ID:    1,
						Image: "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
						Name:  "Queen",
						Members: []string{
							"Freddie Mercury",
							"Brian May",
							"John Daecon",
							"Roger Meddows-Taylor",
							"Mike Grose",
							"Barry Mitchell",
							"Doug Fogie",
						},
						CreationDate: 1970,
						FirstAlbum:   "14-12-1973",
						Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",

						ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",

						Relations: "https://groupietrackers.herokuapp.com/api/relation/1",
					},
				},
			},
			want: MyArtist{
				ID:    1,
				Image: "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
				Name:  "Queen",
				Members: []string{
					"Freddie Mercury",
					"Brian May",
					"John Daecon",
					"Roger Meddows-Taylor",
					"Mike Grose",
					"Barry Mitchell",
					"Doug Fogie",
				},
				CreationDate: 1970,
				FirstAlbum:   "14-12-1973",
				Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",

				ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",

				Relations: "https://groupietrackers.herokuapp.com/api/relation/1",
			},
			wantErr: false,
		},
	}
	// loop over testcases, run getartistbyID function inside loop: compare result to want. 
}

func TestGetData(t *testing.T) {
	//check no errors, and that the returns have something in them. check length
}
