package pipes_and_filters

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type FilterWithName struct {
	Name     string
	Function FilterWithParams
}
type FilterWithParams func(data any, params map[string]any) error

type SelectedFilterFromYaml struct {
	Name   string         `yaml:"name"`
	Params map[string]any `yaml:"params"`
}

func (p *Pipeline) LoadFiltersFromYaml(yamlPath string, availableFilters map[string]FilterWithParams) {

	// Read yaml file
	yamlFile, errReadingFile := ioutil.ReadFile(yamlPath)
	if errReadingFile != nil {
		panic(errReadingFile)
	}

	// Parse yaml file
	var selectedFilters []SelectedFilterFromYaml
	errParsingYaml := yaml.Unmarshal(yamlFile, &selectedFilters)

	if errParsingYaml != nil {
		panic(errParsingYaml)
	}

	// Insert filters in Pipe
	for _, selectedFilter := range selectedFilters {
		filterName, filterExists := availableFilters[selectedFilter.Name]
		if !filterExists {
			panic("Filter " + selectedFilter.Name + " not found")
		}
		p.Use(insertParameters(filterName, selectedFilter.Params))
	}
}

func insertParameters(missingParameterFilter FilterWithParams, params map[string]any) Filter {
	return func(data any) error {
		return missingParameterFilter(data, params)
	}
}
