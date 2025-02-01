package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// prevent "/" catching all the URL
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := "Hello world :)"
	w.Write([]byte(resp))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	resp := "Snippet Display..."
	w.Write([]byte(resp))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// only allow POST method
	if r.Method != "POST" {
		// allowed 
		// Changing the response header map after a call to 
		// w.WriteHeader() or w.Write() will have no effect on the headers that the user receives
		w.Header().Set("Allow", "POST")

		// response
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := "Create snippet here..."
	w.Write([]byte(resp))
}

func snippetTest(w http.ResponseWriter, r *http.Request) {
	// purpose: try resquesting with /snippet instead of /snippet/
	// error 301 permanene request

	resp := "Snippet!!!"
	w.Write([]byte(resp))
}

func main() {
	// web server
	mux := http.NewServeMux()
	//subtree path
	mux.HandleFunc("/", home) // gonna catch other reg (e.g., /foo)
	mux.HandleFunc("/snippet/", snippetTest)

	// fixed path
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// listen
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
