package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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
	rrBodystring := (strings.Count(rr.Body.String(), "/concert?id="))
	if rrBodystring < 2 {
		t.Fatal("Not enough entries in the page")
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestTourPageHandler(t *testing.T) {
	idsToTest := map[int]int{
		0: http.StatusBadRequest,
		1: http.StatusOK,
	}
	for id, wantStatus := range idsToTest {
		req, err := http.NewRequest("GET", fmt.Sprintf("/tour?id=%d", id), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tourPage)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != wantStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}

func TestConcertPageHandler(t *testing.T) {
	idsToTest := map[int]int{
		0: http.StatusBadRequest,
		1: http.StatusOK,
	}
	for id, wantStatus := range idsToTest {
		req, err := http.NewRequest("GET", fmt.Sprintf("/concert?id=%d", id), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(concertPage)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != wantStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}
