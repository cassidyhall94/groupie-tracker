package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mainPage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestTourPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/tour", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tourPage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestConcertPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/concert", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(concertPage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// func TestGroupie(t *testing.T) {
// 	testCases := []struct {
// 		banner           string
// 		input            string
// 		expectedResponse int
// 	}{
// 		{
// 			// when an invalid banner file
// 			banner:           "shadwo.txt",
// 			input:            "this is history",
// 			expectedResponse: http.StatusInternalServerError,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		form := url.Values{}
// 		form.Add("Banner", tc.banner)
// 		form.Add("input", tc.input)

// 		request := httptest.NewRequest("POST", "/ascii-art", strings.NewReader(form.Encode()))
// 		request.PostForm = form
// 		responseRecorder := httptest.NewRecorder()

// 		process(responseRecorder, request)
// 		if responseRecorder.Code != tc.expectedResponse {
// 			t.Errorf("Want status '%d', got '%d'", tc.expectedResponse, responseRecorder.Code)
// 		}
// 		// assert.Equal(t, responseRecorder.Code, tc.expectedResponse)
// 	}
// }

// func TestPostEndpoint(t *testing.T) {
// 	svr := httptest.NewServer(http.HandlerFunc(process))
// 	defer svr.Close()
// 	form := url.Values{}
// 	form.Add("Banner", "thinkertoy.txt")
// 	form.Add("input", "hello sad world")

// 	request, err := http.NewRequest("POST", svr.URL+"/ascii-art", strings.NewReader(form.Encode()))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	resp, err := http.DefaultClient.Do(request)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.StatusCode != 200 {
// 		t.Errorf("Want status '%d', got '%d'", 200, resp.StatusCode)
// 	}
// }

// func TestGetEndpoint(t *testing.T) {
// 	svr := httptest.NewServer(http.HandlerFunc(process))
// 	defer svr.Close()
// 	resp, err := http.Get(svr.URL + "/")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.StatusCode != 200 {
// 		t.Errorf("Want status '%d', got '%d'", 200, resp.StatusCode)
// 	}
// }
