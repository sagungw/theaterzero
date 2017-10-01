package repository

import (
	"bytes"
	"encoding/json"

	"theaterzero.com/backend/lib/utils/elasticsearch"

	"theaterzero.com/backend/lib/config"
	"theaterzero.com/backend/lib/persistence/model"
)

const (
	allCinemasTemplateTag = "resources/json/all-cinemas.json.tmpl"
)

// GetAllCinemas :nodoc:
// func GetAllCinemas() (cinemas []*model.Cinema, err error) {
// buf, err := http.UseTemplate(allCinemasTemplateTag, struct{}{})
// if err != nil {
// 	return
// }

// resp, err := http.Post(config.ScraperIndexURL("cinema/_search"), &buf)
// if err != nil {
// 	return
// }

// result := &model.Response{}
// err = persistence.JSONToModel(resp, result)
// if err != nil {
// 	return
// }

// for _, hit := range result.Hits.Hits {
// 	source := hit.Source
// 	cinema := &model.Cinema{
// 		ID:   source["id"].(string),
// 		Name: source["name"].(string),
// 		URL:  source["url"].(string),
// 	}
// 	cinemas = append(cinemas, cinema)
// }

// return
// }

// SaveCinema :nodoc:
func SaveCinema(cinema *model.Cinema) (err error) {
	b, err := json.Marshal(cinema)
	if err != nil {
		return err
	}

	_, err = elasticsearch.Client.Put(config.ScraperIndexURL("cinema/"+cinema.ID), bytes.NewBuffer(b))
	return
}

// BulkSaveCinema :nodoc:
func BulkSaveCinema(cinemas []*model.Cinema) (err error) {
	payload := []byte{}
	for _, o := range cinemas {
		b, err := json.Marshal(o)
		if err != nil {
			return err
		}

		payload = append(payload, []byte("{\"index\":{\"_id\":"+"\""+o.ID+"\""+"}}\n")...)
		payload = append(payload, b...)
		payload = append(payload, []byte("\n")...)
	}

	_, err = elasticsearch.Client.Post(config.ScraperIndexURL("cinema/_bulk"), bytes.NewBuffer(payload))
	return err
}
