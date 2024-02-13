package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	var geminiData []string

	c.OnHTML("message-content", func(e *colly.HTMLElement) {
		result := e.Text
		geminiData = append(geminiData, result)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://gemini.google.com/app/e676dd297aac3034")
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(geminiData)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(jsonData)
		if err != nil {
			log.Println("Error writing JSON response:", err)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
