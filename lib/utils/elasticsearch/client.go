package elasticsearch

import (
	"net/http"

	"theaterzero.com/backend/lib/config"
	h "theaterzero.com/backend/lib/utils/http"
)

// Client :nodoc:
var Client *h.RestClient

func init() {
	Client = h.WithClient(&http.Client{Timeout: config.ElasticSearchTimeout()})
}
