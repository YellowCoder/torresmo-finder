package main

func main() {
	site := &Site{URL: "https://www.imdb.com"}
	site.pages = append(
		site.pages,
		&Page{
			URL:   "https://www.imdb.com/movies-in-theaters/?ref_=nv_mv_inth",
			query: ".sub-list table",
		},
	)
	site.pages = append(
		site.pages,
		&Page{
			URL:   "https://www.imdb.com/movies-coming-soon/?ref_=inth_cs",
			query: ".list_item table",
		},
	)
	site.Run()
}
