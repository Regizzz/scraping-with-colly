package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type News struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	allNews := make([]News, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com", "www.news.ycombinator.com"),
	)

	collector.OnHTML(".athing", func(element *colly.HTMLElement) {
		newsId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}

		newsTitle := element.Text

		news := News{
			ID:    newsId,
			Title: newsTitle,
		}

		allNews = append(allNews, news)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://news.ycombinator.com")

	writeJSON(allNews)
}

func writeJSON(data []News) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable create json file")
		return
	}

	_ = ioutil.WriteFile("hackernews.json", file, 0644)
}
