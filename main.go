package main

import (
	"html/template"
	"net/http"
	"what-do-u-meme-app/app"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	// add func to handle url request
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/mock/search_by_text", app.SearchByTextMock)
	http.HandleFunc("/mock/get_trending", app.GetTrendingMock)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// listen and serve
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
