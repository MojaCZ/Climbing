package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	// "fmt"
	"os"
)

var (
	IndexTpl *template.Template
	ArticleTpl *template.Template
)

func init() {
	IndexTpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
	ArticleTpl = template.Must(template.ParseGlob("./templates/Article/*.gohtml"))
}

func main() {

	http.Handle("/Climbing/files/", http.StripPrefix("/Climbing/files", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/Climbing/Article/", article)
	http.HandleFunc("/", index)
	// http.HandleFunc("/Climbing/Article/", http.StripPrefix("/Climbing/Article", article))
	http.HandleFunc("/Climbing/favicon.ico", ico)

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatalln(err)
	}
}


func ico(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static/img/note.png")
}

func index(w http.ResponseWriter, req *http.Request) {

	err := IndexTpl.ExecuteTemplate(w, "header.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = IndexTpl.ExecuteTemplate(w, "menu.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = IndexTpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = IndexTpl.ExecuteTemplate(w, "footer.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func article(w http.ResponseWriter, req *http.Request) {
	articleHTML := path.Base(req.URL.Path) + ".gohtml"

	// return error if article doesn't exists
	if _, err := os.Stat("./templates/Article/" + articleHTML); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("THIS ARTICLE DOESN'T EXISTS"))
		return
	}

	err := IndexTpl.ExecuteTemplate(w, "header.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = IndexTpl.ExecuteTemplate(w, "menu.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = ArticleTpl.ExecuteTemplate(w, articleHTML, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = IndexTpl.ExecuteTemplate(w, "footer.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
