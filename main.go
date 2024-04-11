package main

import (
	"bufio"
	json2 "encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
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
		if exists {
			displayChapter(story[url], w)
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

func displayChapter(chapter Chapter, w http.ResponseWriter) {
	fmt.Println("Title: " + chapter.Title)
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
	t.Execute(w, chapter)
}

func cliGame(story Story) {
	showChapter(story["intro"], story, "intro")
}

func showChapter(chapter Chapter, story Story, url string) {
	fmt.Println(chapter.Title)
	for _, p := range chapter.Paragraphs {
		fmt.Println(p)
	}
	for n, p := range chapter.Options {
		fmt.Printf("%d) ", n)
		fmt.Println(p.Text)
	}
	if url == "home" {
		fmt.Println("Thanks for playing!!")
		os.Exit(0)
	}
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	res = strings.TrimRight(res, "\n")
	res = strings.TrimSpace(res)
	i, err := strconv.Atoi(res)
	if err != nil {
		panic(err)
	}
	next := chapter.Options[i].Chapter
	newChapter := story[chapter.Options[i].Chapter]
	showChapter(newChapter, story, next)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", intro)
	return mux
}

func intro(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Intro")
}

func main() {
	cli := flag.Bool("c", false, "CLI option")
	flag.Parse()
	story := readJSON("gopher.json")
	if !*cli {
		mux := defaultMux()
		startServer(story, mux)
	} else {
		fmt.Println("Welcome to the CLI version of this game")
		cliGame(story)
	}

}
