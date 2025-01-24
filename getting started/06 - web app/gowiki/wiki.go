package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to wiki!")
}

func handlerView(w http.ResponseWriter, r *http.Request, str string) {
	fmt.Println(str)

	prefix := "/view/"
	title := r.URL.Path[len(prefix):]
	page, _ := loadPage(title)
	fmt.Fprintf(w, `
		Title: %s 
		Body: %s
	`, page.Title, page.Body)
}

// testing: scale when handler need more than response and request
func makeHandler(cb func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		str := "deez nut"
		cb(w, r, str)
	}
}

func main() {
	// test page
	// p1 := Page{Title: "test2", Body: []byte("A simply page body ")}
	// p1.save()

	// p2, _ := loadPage(p1.Title)
	// fmt.Println(string(p2.Body))

	// test http
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/view/", makeHandler(handlerView))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
