package main

import "testing"

func TestGetArtistsData(t *testing.T) {

	artists, err := GetArtistsData()
	if err != nil {
		t.Fatalf("artists err:%+v", err)
	}
	if len(artists) < 1 {
		t.Fatalf("artists failed due to length < 1")
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
				ID: 0,
				MyArtists: {
					{},
				},
			},
			want: {
				ID: ,
				Image: ,
			},
			wantErr: ,
		},
	}
}
