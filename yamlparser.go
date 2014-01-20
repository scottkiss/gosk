package gosk

import (
	"github.com/scottkiss/go-gypsy/yaml"
	"log"
	"os"
	"strings"
)

type YamlParser struct{}

var YAML_FILES = [3]string{"config.yml", "pages.yml", "nav.yml"}

func (yp *YamlParser) parse(root string) map[string]interface{} {
	var yamlFilesConfig = make(map[string]interface{})

	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	for _, yamlFile := range YAML_FILES {
		path := root + yamlFile
		if !isExists(path) {
			log.Panic(path + " file not found!")
			os.Exit(1)
		}

		config, err := yaml.ReadFile(path)
		if err != nil {
			log.Panic(err)
			os.Exit(1)
		}

		yamlFilesConfig[yamlFile] = config
	}

	return yamlFilesConfig
}
