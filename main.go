package main

type Chapter struct {
	name    string
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// TODO READ JSON FILE
func readJSON(s string) []any {

	return nil
}

// TODO Server Func
func startServer() any {
	return nil
}

func main() {

}
