package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Site struct {
	URL   string
	pages []*Page
}

type Page struct {
	URL    string
	movies []*Movie
	query  string
}

type Movie struct {
	Title     string
	URL       string
	Thumbnail string
}

func (s *Site) Run() {
	for _, page := range s.pages {
		page.Run()
	}
}

func (p *Page) Run() {
	doc, err := goquery.NewDocument(p.URL)

	if err != nil {
		return
	}

	doc.Find(p.query).Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h4").Text())
		url, _ := s.Find("h4 a").Attr("href")
		thumbnail, _ := s.Find("image img").Attr("src")

		movie := &Movie{title, url, thumbnail}
		p.movies = append(p.movies, movie)
	})

	for _, m := range p.movies {
		m.Run()
	}
}

func (m *Movie) Run() {
	fmt.Println(m)
}
