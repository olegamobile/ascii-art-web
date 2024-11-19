package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var pageTemplate *template.Template

type webPage struct {
	Input  string
	Font   string
	Output string
}

func main() {
	pageTemplate, _ = pageTemplate.ParseGlob("templates/*.html")
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/ascii-art", mainPageHandler)

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	var mainPage webPage
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		ErrorHandler(404, w)
		return
	}
	if r.Method == "GET" && r.URL.Path == "/" {
		mainPage.Input = ""
		mainPage.Font = "standard.txt"
	} else if r.Method == "POST" && r.URL.Path == "/ascii-art" {
		mainPage.Input = r.FormValue("input")
		mainPage.Font = r.FormValue("banner")
		if mainPage.Input == "" || mainPage.Font == "" {
			ErrorHandler(400, w)
			return
		}

		switch mainPage.Font {
		case "standard":
			mainPage.Font = "standard.txt"
		case "shadow":
			mainPage.Font = "shadow.txt"
		case "thinkertoy":
			mainPage.Font = "thinkertoy.txt"
		default:
			mainPage.Font = "standard.txt"
		}
		asciiArtGenerate(&mainPage)
		if strings.Contains(mainPage.Output, "Error") {
			ErrorHandler(500, w)
			return
		}
	} else {
		ErrorHandler(405, w)
		return
	}

	pageTemplate.ExecuteTemplate(w, "index.html", mainPage)
}

func asciiArtGenerate(mainpage *webPage) {
	mainpage.Input = strings.ReplaceAll(mainpage.Input, "\r\n", "\\n")
	if !isTextPrintable(mainpage.Input) {
		mainpage.Output = "Text is not printable. Please enter valid ASCII characters."
	} else {
		mainpage.Output = convertToAsciiArt(mainpage.Input, mainpage.Font)
	}
	mainpage.Input = strings.ReplaceAll(mainpage.Input, "\\n", "\r\n")
}

func ErrorHandler(err int, w http.ResponseWriter) {
	errorMessage := ""

	w.WriteHeader(err)

	switch err {
	case 400:
		errorMessage = "400 bad request"
	case 404:
		errorMessage = "404 not found"
	case 405:
		errorMessage = "405 method not allowed"
	default:
		errorMessage = "500 internal server error"
	}
	pageTemplate.ExecuteTemplate(w, "error.html", errorMessage)
}
