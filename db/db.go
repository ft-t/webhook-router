package db

import (
	"context"
	"reflect"
	"time"

	"webhook-router/configuration"

	"github.com/olivere/elastic"
)

var elasticConfiguration configuration.ElasticConfiguration
var elasticClient *elastic.Client

type PathRule struct {
}

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

func GetRulesByPath(path string) []PathRule {
	exists, err := elasticClient.IndexExists(elasticConfiguration.Index).Do(context.Background())

	if err != nil {
		//TODO Handle
		return nil
	}

	if !exists {
		return nil
	}

	var result *elastic.SearchResult
	result, err = elasticClient.Search().Index(elasticConfiguration.Index).Type(elasticConfiguration.DocType).Query(elastic.NewTermQuery("path", path)).Do(context.Background())

	if err != nil {
		//TODO Handle
		return nil
	}

	var rules []PathRule
	for _, rule := range result.Each(reflect.TypeOf((*PathRule)(nil))) {
		rules = append(rules, rule.(PathRule))
	}

	return rules;
}
func GetAllRules() []PathRule {
	exists, err := elasticClient.IndexExists(elasticConfiguration.Index).Do(context.Background())

	if err != nil {
		//TODO Handle
		return nil
	}

	if !exists {
		return nil
	}

	var result *elastic.SearchResult
	result, err = elasticClient.Search().Index(elasticConfiguration.Index).Type(elasticConfiguration.DocType).Query(elastic.NewMatchAllQuery()).Do(context.Background())

	if err != nil {
		//TODO Handle
		return nil
	}

	var rules []PathRule
	for _, rule := range result.Each(reflect.TypeOf((*PathRule)(nil))) {
		rules = append(rules, rule.(PathRule))
	}

	return rules;
}
