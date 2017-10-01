package scraper

import (
	"github.com/antchfx/xquery/html"
	"theaterzero.com/backend/lib/log"
	"theaterzero.com/backend/lib/persistence/model"
	"theaterzero.com/backend/lib/persistence/repository"
	"theaterzero.com/backend/scraper/cineplex"
)

// ScrapeCinemas :nodoc:
func ScrapeCinemas() {
	doc, err := htmlquery.LoadURL(cineplex.CinemaListURL())
	if err != nil {
		log.Error(err)
	}

	nodes := htmlquery.Find(doc, cineplex.CinemaElement())

	cs := make([]*model.Cinema, len(nodes))
	for i, el := range nodes {
		c := ExtractCinema(el)
		cs[i] = c
	}

	err = repository.BulkSaveCinema(cs)
	if err != nil {
		log.Errorf("Fetching cinemas failed. %s\n", err.Error())
	}
}
