package main

import	"fmt"
import	"net/http"
import	"html/template"
import "os"
import "io/ioutil"
//import	"path/filepath"

func main() {
	f, err := os.Open("views/tes.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	result, _ := ioutil.ReadAll(f)
	fmt.Println(string(result))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        //filepaths := filepath.Join(filepath.Dir("views"), "index.html")
        //paths := filepath.Join("views", filepath.Clean(r.URL.Path))
        tmpl := template.Must(template.ParseFiles("views/index.html",))

        var data = map[string]interface{}{
        	"title": "learning go go lang lang",
        	"name": "Superman",
        }

        err := tmpl.ExecuteTemplate(w, "index", data)
        if err != nil{
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    http.HandleFunc("/teslagi", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("views/about.html",))

        var dat = map[string]interface{}{
        	"title": "kenapa kok gak bisa",
        	"name": "Batman",
        }

        err := tmpl.ExecuteTemplate(w, "about", dat)
        if err != nil{
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })


    http.HandleFunc("/cesar", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "ini respon cesar")
    })

    http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("public"))))

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}