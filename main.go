package main

// go code for Rick and Morty API example
// https://github.com/pitakill/rickandmortyapigowrapper
// https://git.learn.01founders.co/root/public/src/branch/master/subjects/groupie-tracker/audit
// https://pkg.go.dev/encoding/json#Unmarshal  -- used for the data
// https://groupietrackers.herokuapp.com/api -- JSON data for website
// https://github.com/Muslimah94/groupie-tracker -- code example on github
// https://gist.github.com/rodkranz/0a8ed14fa44b5860f6668efae02b3ea5 -- big files using marshal

import (
	"fmt"
	"log"
	"net/http"
)

// type artistData struct {
// 	ID           int                 `json:"id"`
// 	Image        string              `json:"image"`
// 	Name         string              `json:"name"`
// 	Members      []string            `json:"members"`
// 	CreationDate int                 `json:"creationDate"`
// 	FirstAlbum   string              `json:"firstAlbum"`
// 	Relation     string              `json:"relations"`
// 	Concerts     map[string][]string `json:"datesLocations"`
// }

// type relation struct {
// 	ID       int                 `json:"id"`
// 	Concerts map[string][]string `json:"datesLocations"`
// }

// func (e *artistData) Unmarshal(b []byte) error {
// 	return json.Unmarshal(b, e)
// }

func main() {
	http.HandleFunc("/", process)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	fmt.Printf("Starting server at port localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// type Groupie struct {
	// 	artists   string
	// 	locations string
	// 	dates     string
	// 	relation  string
	// }

	// var groupies []interface{}
	// err := json.Unmarshal(str, &groupies)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for _, v := range groupies {
	// 	z := v.(map[string]interface{})
	// 	for k2, v2 := range z {
	// 		fmt.Println("Key:", k2, "Value:", v2)
	// 	}
	// }
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Status not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "index.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // if there is an error it returns bad request 400
			return
		}

		// input := r.FormValue("input")
		// banner := r.FormValue("Banner")

		// response, err := AsciiArt(input, banner)
		// if err != nil {
		// 	http.Error(w, "No such file or directory: Internal Server Error 500", http.StatusInternalServerError)
		// 	return
		// }

		// _, _ = w.Write([]byte(response)) // Write returns the response with a 200 status code in the header as this is built into the Write function
	default:
		http.Error(w, "Sorry, only GET and POST methods are supported.", http.StatusUnsupportedMediaType)
	}
}
