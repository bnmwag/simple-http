package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type Data struct {
	RequestCount  int
	Uptime        int
	RequestPath   string
	ServerPort    string
	ServerVersion string
	Subdomain     string
}

func main() {
	serverStartTime := time.Now()

	data := Data{
		RequestCount:  0,
		Uptime:        0,
		RequestPath:   "",
		ServerPort:    "8080",
		ServerVersion: "1.0.0",
		Subdomain:     "",
	}

	indexTemplate, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			http.NotFound(w, r)
			return
		}

		data.RequestCount++
		data.RequestPath = r.RequestURI // TODO: sanitize path
		data.Uptime = int(time.Since(serverStartTime).Seconds())

		if len(strings.Split(r.Host, ".")) > 2 {
			data.Subdomain = strings.Split(r.Host, ".")[0]
			println(data.Subdomain)
		}

		err := indexTemplate.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})

	// start server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
