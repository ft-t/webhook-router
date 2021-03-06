package configuration

import (
	"os"

	"github.com/micro/go-config"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/file"
)

type ElasticConfiguration struct {
	Url     string `json:"url"`
	Index   string `json:"index"`
	DocType string `json:"doc_type"`
}
type LoggingConfiguration struct {
	MinLevel string `json:"min_level"`
}
type GlobalConfiguration struct {
	ElasticConfiguration ElasticConfiguration `json:"elastic"`
	LoggingConfiguration LoggingConfiguration `json:"logging"`
	Port                 int                  `json:"port"`
	UiPort               int                  `json:"ui_port"`
}

var configuration *GlobalConfiguration

func GetConfiguration() GlobalConfiguration {
	if configuration != nil {
		return *configuration
	}

	temp := GlobalConfiguration{Port: 8080}
	configuration = &temp

	conf := config.NewConfig()

	var sources []source.Option

	sources = append(sources, file.WithPath("./config.json"))

	if _, err := os.Stat("./config.dev.json"); err == nil {
		sources = append(sources, file.WithPath("./config.dev.json"))
	}

	conf.Load(file.NewSource(
		sources...
	))

	conf.Get().Scan(configuration)

	return *configuration
}
