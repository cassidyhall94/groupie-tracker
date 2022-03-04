package main

import (
	"reflect"
	"testing"
)

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
							"Freddie Mercury Brian May John Daecon Roger Meddows-Taylor Mike Grose Barry Mitchell Doug Fogie",
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
					"Freddie Mercury Brian May John Daecon Roger Meddows-Taylor Mike Grose Barry Mitchell Doug Fogie",
				},
				CreationDate: 1970,
				FirstAlbum:   "14-12-1973",
				Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",

				ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",

				Relations: "https://groupietrackers.herokuapp.com/api/relation/1",
			},
			wantErr: false,
		},
		{
			in: input{
				ID: 55,
				MyArtists: []MyArtist{
					{
						ID:    55,
						Image: "https://groupietrackers.herokuapp.com/api/images/mychemicalromance.jpeg",
						Name:  "My Chemical Romance",
						Members: []string{
							"Gerard Way Mikey Way Ray Toro Frank Iero",
						},
						CreationDate: 2001,
						FirstAlbum:   "23-07-2002",
						Locations:    "https://groupietrackers.herokuapp.com/api/locations/55",

						ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/55",

						Relations: "https://groupietrackers.herokuapp.com/api/relation/55",
					},
				},
			},
			want: MyArtist{
				ID:    0,
				Image: "",
				Name:  "",
				Members: []string{
					"",
				},
				CreationDate: 0,
				FirstAlbum:   "",
				Locations:    "",

				ConcertDates: "",

				Relations: "",
			},
			wantErr: true,
		},
	}

	// loop over testcases, run getartistbyID function inside loop: compare result to want.
	for _, tc := range testCases {
		result, err := GetArtistByID(tc.in.ID, tc.in.MyArtists)
		if err != nil {
			t.Logf("result failed error:%+v", err)
		}
		if !reflect.DeepEqual(result, tc.want) && tc.wantErr == false {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
		if reflect.DeepEqual(result, tc.want) && tc.wantErr == true {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
	}
}

type inputdate struct {
	ID    int
	Dates MyDates
}

func TestGetDateByID(t *testing.T) {
	testCases := []struct {
		in      inputdate
		want    MyDate
		wantErr bool
	}{
		{
			in: inputdate{
				ID: 1,
				Dates: MyDates{
					Index: []MyDate{
						{
							ID: 1,
							Dates: []string{
								"*23-08-2019 *22-08-2019 *20-08-2019 *26-01-2020 *28-01-2020 *30-01-2019 *07-02-2020 *10-02-2020",
							},
						},
					},
				},
			},
			want: MyDate{
				ID: 1,
				Dates: []string{
					"*23-08-2019 *22-08-2019 *20-08-2019 *26-01-2020 *28-01-2020 *30-01-2019 *07-02-2020 *10-02-2020",
				},
			},
			wantErr: false,
		},
		{
			in: inputdate{
				ID: 55,
				Dates: MyDates{
					Index: []MyDate{
						{
							ID: 1,
							Dates: []string{
								"*30-01-2020 *28-01-2020 *30-09-2022 *07-02-2020 *11-02-2020",
							},
						},
					},
				},
			},
			want: MyDate{
				ID: 0,
				Dates: []string{
					"",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		result, err := GetDateByID(tc.in.ID, tc.in.Dates)
		if err != nil {
			t.Logf("result failed error:%+v", err)
		}
		if !reflect.DeepEqual(result, tc.want) && tc.wantErr == false {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
		if reflect.DeepEqual(result, tc.want) && tc.wantErr == true {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
	}
}

type inputlocations struct {
	ID        int
	Locations MyLocations
}

func TestGetLocationsByID(t *testing.T) {
	testCases := []struct {
		in      inputlocations
		want    MyLocation
		wantErr bool
	}{
		{
			in: inputlocations{
				ID: 1,
				Locations: MyLocations{
					Index: []MyLocation{
						{
							ID: 1,
							Locations: []string{
								"north_carolina-usa georgia-usa los_angeles-usa saitama-japan osaka-japan nagoya-japan penrose-new_zealand dunedin-new_zealand",
							},
							Dates: "https://groupietrackers.herokuapp.com/api/dates/1",
						},
					},
				},
			},
			want: MyLocation{
				ID: 1,
				Locations: []string{
					"north_carolina-usa georgia-usa los_angeles-usa saitama-japan osaka-japan nagoya-japan penrose-new_zealand dunedin-new_zealand",
				},
				Dates: "https://groupietrackers.herokuapp.com/api/dates/1",
			},
			wantErr: false,
		},
		{
			in: inputlocations{
				ID: 55,
				Locations: MyLocations{
					Index: []MyLocation{
						{
							ID: 55,
							Locations: []string{
								"new_jersey-usa new_york-usa los_angeles-usa saitama-japan london-united_kingdom nagoya-japan penrose-new_zealand dunedin-new_zealand",
							},
							Dates: "https://groupietrackers.herokuapp.com/api/dates/55",
						},
					},
				},
			},
			want: MyLocation{
				ID: 0,
				Locations: []string{
					"",
				},
				Dates: "",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		result, err := GetLocationByID(tc.in.ID, tc.in.Locations)
		if err != nil {
			t.Logf("result failed error:%+v", err)
		}
		if !reflect.DeepEqual(result, tc.want) && tc.wantErr == false {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
		if reflect.DeepEqual(result, tc.want) && tc.wantErr == true {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
	}
}

type inputrelations struct {
	ID        int
	Relations MyRelations
}

func TestGetRelationByID(t *testing.T) {
	testCases := []struct {
		in      inputrelations
		want    MyRelation
		wantErr bool
	}{
		{
			in: inputrelations{
				ID: 1,
				Relations: MyRelations{
					Index: []MyRelation{
						{
							ID: 1,
							DatesLocations: map[string][]string{
								"dunedin-new_zealand": {"10-02-2020"}, "georgia-usa": {"22-08-2019"}, "los_angeles-usa": {"20-08-2019"}, "nagoya-japan": {"30-01-2019"}, "north_carolina-usa": {"23-08-2019"}, "osaka-japan": {"28-01-2020"}, "penrose-new_zealand": {"07-02-2020"}, "saitama-japan": {"26-01-2020"},
							},
						},
					},
				},
			},
			want: MyRelation{
				ID: 1,
				DatesLocations: map[string][]string{
					"dunedin-new_zealand": {"10-02-2020"}, "georgia-usa": {"22-08-2019"}, "los_angeles-usa": {"20-08-2019"}, "nagoya-japan": {"30-01-2019"}, "north_carolina-usa": {"23-08-2019"}, "osaka-japan": {"28-01-2020"}, "penrose-new_zealand": {"07-02-2020"}, "saitama-japan": {"26-01-2020"},
				},
			},
			wantErr: false,
		},
		{
			in: inputrelations{
				ID: 55,
				Relations: MyRelations{
					Index: []MyRelation{
						{
							ID: 55,
							DatesLocations: map[string][]string{
								"dunedin-new_zealand": {"10-02-2020"}, "georgia-usa": {"22-08-2019"}, "los_angeles-usa": {"20-08-2019"}, "nagoya-japan": {"30-01-2019"}, "north_carolina-usa": {"23-08-2019"}, "osaka-japan": {"28-01-2020"}, "penrose-new_zealand": {"07-02-2020"}, "saitama-japan": {"26-01-2020"},
							},
						},
					},
				},
			},
			want: MyRelation{
				ID: 0,
				DatesLocations: map[string][]string{
					"": {""},
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		result, err := GetRelationByID(tc.in.ID, tc.in.Relations)
		if err != nil {
			t.Logf("result failed error:%+v", err)
		}
		if !reflect.DeepEqual(result, tc.want) && tc.wantErr == false {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
		if reflect.DeepEqual(result, tc.want) && tc.wantErr == true {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
	}
}

type inputfulldata struct {
	ID          int
	ArtistsFull []MyArtistFull
}

func TestGetFullDataByID(t *testing.T) {
	testCases := []struct {
		in      inputfulldata
		want    MyArtistFull
		wantErr bool
	}{
		{
			in: inputfulldata{
				ID: 1,
				ArtistsFull: []MyArtistFull{
					{
						ID:    1,
						Image: "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
						Name:  "Queen",
						Members: []string{
							"Freddie Mercury Brian May John Daecon Roger Meddows-Taylor Mike Grose Barry Mitchell Doug Fogie",
						},
						CreationDate: 1970,
						FirstAlbum:   "14-12-1973",
						Locations:    []string{"north_carolina-usa georgia-usa los_angeles-usa saitama-japan osaka-japan nagoya-japan penrose-new_zealand dunedin-new_zealand"},
						ConcertDates: []string{
							"*23-08-2019 *22-08-2019 *20-08-2019 *26-01-2020 *28-01-2020 *30-01-2019 *07-02-2020 *10-02-2020",
						},
						DatesLocations: map[string][]string{
							"dunedin-new_zealand": {"10-02-2020"}, "georgia-usa": {"22-08-2019"}, "los_angeles-usa": {"20-08-2019"}, "nagoya-japan": {"30-01-2019"}, "north_carolina-usa": {"23-08-2019"}, "osaka-japan": {"28-01-2020"}, "penrose-new_zealand": {"07-02-2020"}, "saitama-japan": {"26-01-2020"},
						},
						WikiLink: []string{
							"https://en.wikipedia.org/wiki/Freddie_Mercury https://en.wikipedia.org/wiki/Brian_May https://en.wikipedia.org/wiki/John_Deacon https://en.wikipedia.org/wiki/Roger_Taylor_(Queen_drummer) https://www.wikidata.org/wiki/Q6847086 https://en.wikipedia.org/wiki/Queen_(band) https://en.wikipedia.org/wiki/Queen_(band)",
						},
					},
				},
			},
			want: MyArtistFull{
				ID:    1,
				Image: "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
				Name:  "Queen",
				Members: []string{
					"Freddie Mercury Brian May John Daecon Roger Meddows-Taylor Mike Grose Barry Mitchell Doug Fogie",
				},
				CreationDate: 1970,
				FirstAlbum:   "14-12-1973",
				Locations:    []string{"north_carolina-usa georgia-usa los_angeles-usa saitama-japan osaka-japan nagoya-japan penrose-new_zealand dunedin-new_zealand"},
				ConcertDates: []string{
					"*23-08-2019 *22-08-2019 *20-08-2019 *26-01-2020 *28-01-2020 *30-01-2019 *07-02-2020 *10-02-2020",
				},
				DatesLocations: map[string][]string{
					"dunedin-new_zealand": {"10-02-2020"}, "georgia-usa": {"22-08-2019"}, "los_angeles-usa": {"20-08-2019"}, "nagoya-japan": {"30-01-2019"}, "north_carolina-usa": {"23-08-2019"}, "osaka-japan": {"28-01-2020"}, "penrose-new_zealand": {"07-02-2020"}, "saitama-japan": {"26-01-2020"},
				},
				WikiLink: []string{
					"https://en.wikipedia.org/wiki/Freddie_Mercury https://en.wikipedia.org/wiki/Brian_May https://en.wikipedia.org/wiki/John_Deacon https://en.wikipedia.org/wiki/Roger_Taylor_(Queen_drummer) https://www.wikidata.org/wiki/Q6847086 https://en.wikipedia.org/wiki/Queen_(band) https://en.wikipedia.org/wiki/Queen_(band)",
				},
			},
			wantErr: false,
		},
		{
			in: inputfulldata{
				ID: 55,
				ArtistsFull: []MyArtistFull{
					{
						ID:    55,
						Image: "https://groupietrackers.herokuapp.com/api/images/mychemicalromance.jpeg",
						Name:  "My Chemical Romance",
						Members: []string{
							"Gerard Way Mikey Way Ray Toro Frank Iero",
						},
						CreationDate: 2001,
						FirstAlbum:   "23-07-2002",
						Locations:    []string{"new_jersey-usa new_york-usa los_angeles-usa saitama-japan london-united_kingdom nagoya-japan penrose-new_zealand dunedin-new_zealand"},
						ConcertDates: []string{
							"*30-01-2020 *28-01-2020 *30-09-2022 *07-02-2020 *11-02-2020",
						},
						DatesLocations: map[string][]string{
							"dunedin-new_zealand": {"10-02-2020"}, "georgia-usa": {"22-08-2019"}, "los_angeles-usa": {"20-08-2019"}, "nagoya-japan": {"30-01-2019"}, "north_carolina-usa": {"23-08-2019"}, "osaka-japan": {"28-01-2020"}, "penrose-new_zealand": {"07-02-2020"}, "saitama-japan": {"26-01-2020"},
						},
						WikiLink: []string{
							"https://en.wikipedia.org/wiki/Gerard_Way https://en.wikipedia.org/wiki/Ray_Toro https://en.wikipedia.org/wiki/Mikey_Way https://en.wikipedia.org/wiki/Frank_Iero",
						},
					},
				},
			},
			want: MyArtistFull{
				ID:    0,
				Image: "",
				Name:  "",
				Members: []string{
					"",
				},
				CreationDate: 0,
				FirstAlbum:   "",
				Locations:    []string{""},
				ConcertDates: []string{
					"",
				},
				DatesLocations: map[string][]string{
					"": {""},
				},
				WikiLink: []string{
					"",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		result, err := GetFullDataByID(tc.in.ID, tc.in.ArtistsFull)
		if err != nil {
			t.Logf("result failed error:%+v", err)
		}
		if !reflect.DeepEqual(result, tc.want) && tc.wantErr == false {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
		if reflect.DeepEqual(result, tc.want) && tc.wantErr == true {
			t.Fatalf("result got %+v, expected %+v", result, tc.want)
		}
	}
}

func TestGetData(t *testing.T) {
	MyArtistFull, MyArtist, MyLocations, MyDates, MyRelations, MemberWikiLinks, err := GetData()
	if err != nil {
		t.Fatalf("get data function err:%+v", err)
	}
	if len(MyArtistFull) < 1 {
		t.Fatalf("GetData function: MyArtistFull failed due to length < 1")
	}
	if len(MyArtist) < 1 {
		t.Fatalf("GetData function: MyArtist failed due to length < 1")
	}
	if len(MyLocations.Index) < 1 {
		t.Fatalf("GetData function: MyLocations failed due to length < 1")
	}
	if len(MyDates.Index) < 1 {
		t.Fatalf("GetData function: MyDates failed due to length < 1")
	}
	if len(MyRelations.Index) < 1 {
		t.Fatalf("GetData function: MyRelations failed due to length < 1")
	}
	if len(MemberWikiLinks) < 1 {
		t.Fatalf("GetData function: MemberWikiLinks failed due to length < 1")
	}
}
