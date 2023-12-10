package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

type staticHandler struct {
	staticPath string
	indexPage  string
}

func (h staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	log.Println(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func handlerGetHelloWorld(wr http.ResponseWriter,
	req *http.Request) {
	fmt.Fprintf(wr, "Hello, World\n")
	log.Println(req.Method) // request method
	log.Println(req.URL)    // request URL
	log.Println(req.Header) // request headers
	log.Println(req.Body)   // request body)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	result := "Login "
	r.ParseForm()

	if validateUser(r.FormValue("username"), r.FormValue("password")) {
		result = result + "successfull"
	} else {
		result = result + "unsuccessful"
	}

	t, err := template.ParseFiles("static/tmpl/msg.html")

	if err != nil {
		fmt.Fprintf(w, "error processing")
		return
	}

	tpl := template.Must(t, err)

	tpl.Execute(w, result)
}

func validateUser(username string, password string) bool {
	return (username == "admin") && (password == "admin")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", postHandler).Methods("POST")

	spa := staticHandler{staticPath: "static", indexPage: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
