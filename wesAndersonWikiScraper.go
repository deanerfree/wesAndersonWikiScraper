package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// var movies = []string{}

func main() {
	type Movies struct {
		Title string
		Year  string
	}

	var uncleanedData = []string{}
	var cleanedData []Movies

	singleSpacePattern := regexp.MustCompile(`\s+`)
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	c.OnHTML("table.wikitable tbody", func(e *colly.HTMLElement) {
		uncleanedData = append(uncleanedData, e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while scraping: %s\n", err.Error())
	})

	c.Visit("https://en.wikipedia.org/wiki/Wes_Anderson_filmography")
	// fmt.Printf("Type %T \n", uncleanedData[0])

	uncleanedData[0] = singleSpacePattern.ReplaceAllString(uncleanedData[0], " ")
	uncleanedData[0] = strings.Replace(uncleanedData[0], "Yes", "", -1)
	uncleanedData[0] = strings.Replace(uncleanedData[0], "No", "", -1)
	uncleanedData[0] = strings.Replace(uncleanedData[0], "Executive", "", -1)

	uncleanedData = strings.Split(uncleanedData[0], "    ")[1:]
	for _, item := range uncleanedData {

		if len(item) > 5 {
			if item[:3] == "TBA" {
				cleanedData = append(cleanedData, Movies{item[4:], item[:3]})
			}
			if _, err := strconv.Atoi(item[:4]); err == nil {
				cleanedData = append(cleanedData, Movies{item[5:], item[:4]})
			}
			if _, err := strconv.Atoi(item[:4]); err != nil && item[:3] != "TBA" {
				cleanedData = append(cleanedData, Movies{item, "TBA"})
			}
		}

	}

	fmt.Println(cleanedData)

}
