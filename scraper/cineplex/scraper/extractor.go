package scraper

import (
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
	"theaterzero.com/backend/lib/persistence/model"
)

// ExtractCinema :nodoc:
func ExtractCinema(node *html.Node) *model.Cinema {
	node.RemoveChild(htmlquery.FindOne(node, "/span"))

	cinemaName := htmlquery.InnerText(node)
	url := htmlquery.SelectAttr(node, "href")

	return &model.Cinema{
		ID:   model.GenerateID(),
		Name: cinemaName,
		URL:  url,
	}
}
