package config

import (
	"time"

	"github.com/spf13/viper"
	"theaterzero.com/backend/lib/log"
)

// ElasticSearchURL :nodoc:
func ElasticSearchURL() string {
	return viper.GetString("elasticsearch_url")
}

// ElasticSearchTimeout :nodoc:
func ElasticSearchTimeout() time.Duration {
	return time.Duration(viper.GetInt("elasticsearch_timeout")) * time.Second
}

// ScraperIndexURL :nodoc:
func ScraperIndexURL(path string) string {
	return ElasticSearchURL() + "/" + viper.GetString("scraper_index_name") + "/" + path
}

// RedisURL :nodoc:
func RedisURL() string {
	return viper.GetString("redis_url")
}

// RedisPoolCount :nodoc:
func RedisPoolCount() string {
	return viper.GetString("redis_pool_count")
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.SetDefault("elasticsearch_url", "http://localhost:9200")
	viper.SetDefault("elasticsearch_timeout", 5)

	viper.SetDefault("redis_url", "redis://localhost:6379")
	viper.SetDefault("redis_pool_count", 10)

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Failed on loading configs")
	}
}
