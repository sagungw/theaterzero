package cineplex

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
	"theaterzero.com/backend/lib/log"
)

const xpathFile = "./scraper/cineplex/xpaths.yaml"

// Web :nodoc:
var web cineplex

type cineplex struct {
	CinemaListURL string `yaml:"cinema_list_url"`
	CinemaElement string `yaml:"cinema_element"`
}

func init() {
	b, err := ioutil.ReadFile(xpathFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(b, &web)
	if err != nil {
		log.Fatal(err)
	}
}

// CinemaListURL :nodoc:
func CinemaListURL() string {
	return web.CinemaListURL
}

// CinemaElement :nodoc:
func CinemaElement() string {
	return web.CinemaElement
}
