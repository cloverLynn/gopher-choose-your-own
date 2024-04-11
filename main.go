package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func readJSON(s string) Story {
	file, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}
	var json Story
	json2.Unmarshal([]byte(file), &json)
	return json
}

func httpHandler(story Story, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := strings.Split(r.URL.Path, "/")[1]
		_, exists := story[url]
		fmt.Println(story[url])
		if exists {
			http.ServeFile(w, r, "index.html")
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func startServer(story Story, mux http.Handler) {
	fmt.Println("Starting the server on :8080")
	handler := httpHandler(story, mux)
	http.ListenAndServe(":8080", handler)
}

/*func displayChapter(chapter Chapter) {
	func hello(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}
}*/

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", intro)
	mux.HandleFunc("/test", test)
	return mux
}

func intro(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Intro")
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Testing")
}

func main() {
	story := readJSON("gopher.json")
	mux := defaultMux()
	startServer(story, mux)
}
