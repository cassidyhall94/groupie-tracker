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

func TestGetArtistByID(t *testing.T) {
	testCases := []struct{
		input  struct{
			ID int
			MyArtists []MyArtist
		}
		want 	MyArtist
		wantErr bool
	}{
		{
			input: {
				ID: 1,
				MyArtists: {
					{},
				},
			},		
			want: {				
					id:    1,
					image: "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
					name:  "Queen",
					members: {
						"Freddie Mercury",
						"Brian May",
						"John Daecon",
						"Roger Meddows-Taylor",
						"Mike Grose",
						"Barry Mitchell",
						"Doug Fogie",
					},
					creationDate: 1970,
					firstAlbum:   "14-12-1973",
					locations: 
					// "https://groupietrackers.herokuapp.com/api/locations/1",
					{
						"north_carolina-usa",
						"georgia-usa",
						"los_angeles-usa",
						"saitama-japan",
						"osaka-japan",
						"nagoya-japan",
						"penrose-new_zealand",
						"dunedin-new_zealand",
					},
					concertDates:
					// "https://groupietrackers.herokuapp.com/api/dates/1",
					{
						0: "23-08-2019",
						1: "22-08-2019",
						2: "20-08-2019",
						3: "26-01-2020",
						4: "28-01-2020",
						5: "30-01-2019",
						6: "07-02-2020",
						7: "10-02-2020",
					},
					relations:
					// "https://groupietrackers.herokuapp.com/api/relation/1",
					{
						"dunedin-new_zealand": "10-02-2020",
						"georgia-usa":         "22-08-2019",
						"los_angeles_usa":     "20-08-2019",
						"nagoya-japan":        "30-01-2019",
						"north_carolina-usa":  "23-08-2019",
						"osaka-japan":         "28-01-2020",
						"penrose-new_zealand": "07-02-2020",
						"saitama-japan":       "26-01-2020",
					},
				},
			},
			wantErr: false,
		},
	}
