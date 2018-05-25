package db

import (
	"webhook-router/configuration"
	"github.com/olivere/elastic"
	"time"
)

var elasticConfiguration configuration.ElasticConfiguration
var elasticClient *elastic.Client

func getClient() *elastic.Client {
	if elasticClient != nil {
		return elasticClient
	}

	client, err := elastic.NewClient(
		elastic.SetURL(elasticConfiguration.Url),
		elastic.SetHealthcheckInterval(10*time.Second),
	)

	if err != nil {
		// TODO Handle
	}

	elasticClient = client
	return elasticClient
}

func InitDb(configuration configuration.ElasticConfiguration) {
	elasticConfiguration = configuration
}

func GetRulesByPath(path string) {

}
func GetAllRules() {

}
