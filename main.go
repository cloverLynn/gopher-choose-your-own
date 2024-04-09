package main

import (
	json2 "encoding/json"
	"fmt"
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

// TODO httpHandler
func httpHandler() any {
	return nil
}

// TODO Server Func
func startServer() any {
	return nil
}

func main() {
	readJSON("gopher.json")
}
