package db

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"webhook-router/configuration"
	"webhook-router/structs"

	"github.com/olivere/elastic"
)

var config *configuration.GlobalConfiguration
var elasticClient *elastic.Client

func getClient() *elastic.Client {
	if elasticClient != nil {
		return elasticClient
	}

	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticConfiguration.Url),
		elastic.SetHealthcheckInterval(10*time.Second),
	)

	if err != nil {
		// TODO Handle
	}

	elasticClient = client
	return elasticClient
}

func InitDb(configuration *configuration.GlobalConfiguration) {
	config = configuration
}

func GetRulesByPath(path string) []structs.PathRule {
	exists, err := getClient().IndexExists(config.ElasticConfiguration.Index).Do(context.Background())

	if err != nil {
		//TODO Handle
		return nil
	}

	if !exists {
		return nil
	}

	var result *elastic.SearchResult
	result, err = getClient().Search().Index(config.ElasticConfiguration.Index).Type(config.ElasticConfiguration.DocType).Query(elastic.NewTermQuery("path", path)).Do(context.Background())

	if err != nil {
		//TODO Handle
		return nil
	}

	var rules []structs.PathRule
	for _, hit := range result.Hits.Hits {
		var rule structs.PathRule
		err := json.Unmarshal(*hit.Source, &rule)

		if err != nil {
			// TODO
			continue
		}

		rules = append(rules, rule)
	}

	return rules;
}

func GetAllRules() []structs.PathRule {
	exists, err := getClient().IndexExists(config.ElasticConfiguration.Index).Do(context.Background())

	if err != nil {
		//TODO Handle
		return nil
	}

	if !exists {
		return nil
	}

	var result *elastic.SearchResult
	result, err = getClient().Search().Index(config.ElasticConfiguration.Index).Type(config.ElasticConfiguration.DocType).Query(elastic.NewMatchAllQuery()).Do(context.Background())

	if err != nil {
		//TODO Handle
		return nil
	}

	var rules []structs.PathRule
	for _, rule := range result.Each(reflect.TypeOf((*structs.PathRule)(nil))) {
		rules = append(rules, rule.(structs.PathRule))
	}

	return rules;
}
