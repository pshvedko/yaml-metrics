package collector

import "github.com/pshvedko/yaml-metrics/collector/yaml"

func NewCollectorYaml(name string) (*Collector, error) {
	promoter, err := yaml.NewPromoter(name)
	if err != nil {
		return nil, err
	}
	return NewCollector(promoter)
}
