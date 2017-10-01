package scraper

// ScrapeMovies :nodoc:
func ScrapeMovies() {
	// cinemas, err := repository.GetAllCinemas()
	// if err != nil {
	// 	log.Error(err)
	// }

	// for _, cinema := range cinemas {
	// 	doc, err := goquery.NewDocument(cinema.URL)
	// 	if err != nil {
	// 		log.Error(err)
	// 	}

	// 	foundMovies := doc.Find("div#makan table").Eq(0)
	// 	foundMovies = foundMovies.Find("tr:not([class='title'])")
	// 	for index := 0; index < foundMovies.Length(); index++ {
	// 		movie := foundMovies.Eq(index).Find("td:nth-child(1)")
	// 		movieLink, exists := movie.Find("a:nth-child(1)").Attr("href")
	// 		if exists {
	// 			movieDoc, err := goquery.NewDocument(movieLink)
	// 			if err != nil {
	// 				log.Error(err)
	// 			}

	// 			column := movieDoc.Find("div.col-m_392").
	// 			title := column.Find("h2:nth-child(1)").Text()

	// 		}
	// 	}

	// }

}
