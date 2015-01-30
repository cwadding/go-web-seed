package main

import (
	"log"
	"encoding/json"
	"github.com/cwadding/go-web-seed/models"
	"github.com/gorilla/mux"
	"net/http"
);



func Index(w http.ResponseWriter, r *http.Request) {

	profile := models.Profile{"Alex", []string{"snowboarding", "programming"}}
	

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)


	// pages_url, err := App.Router.Get("page_list").URL()
	// if err != nil {
	// 	panic(err)
	// }
	// tmpl := App.Templates["index.html"]
	// ctx := make(map[string]interface{})
	// ctx["pages_url"] = pages_url
	// ctx["title"] = "Web experiment with GO"
	// err = tmpl.ExecuteTemplate(w, "base", ctx)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", Index)
    http.Handle("/", r)

    err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("An error occured when trying to server gowebexp: \n", err)
	}
}