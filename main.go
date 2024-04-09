package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
	"os"
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
	fmt.Printf("%+v\n", json)
	return json
}

func httpHandler(story Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		_, exists := story[url]
		if exists {
			//TODO implemnt serve page function
		} else {
			//TODO continue building out
		}
	}
}

func startServer(story Story) {
	fmt.Println("Starting the server on :8080")
	handler := httpHandler(story)
	http.ListenAndServe(":8080", handler)
}

func main() {
	story := readJSON("gopher.json")
	startServer(story)
}
