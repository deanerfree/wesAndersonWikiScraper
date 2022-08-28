# Wes Anderson Wiki Scraper

## Purpose:
To create a simple webscraper and to learn Golang.
I enjoy Wes Anderson movies and thought this would be a great way to learn the golang.

## Tools Used:
- [Golang](https://go.dev)
- [Colly](github.com/gocolly/colly)
- convstr
- Regex

## Data:
Data from [wikipedia.org](https://en.wikipedia.org/wiki/Wes_Anderson_filmography)

## Challenges:
- Understanding syntax of Golang
- Collecting non unique data from wikipedia, using colly then collecting the Title of the films and the Year it was made including upcoming films
- Collecting the data as a long string then cleaning the string data by applying filters to remove unwanted data. This string was broken into a series of substrings and using a for loop divided the wanted data into Year and Title using a struct and appending that data into a new Slice(array).
