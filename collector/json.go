package collector

import "github.com/pshvedko/yaml-metrics/collector/json"

func NewCollectorJson(name string) (*Collector, error) {
	promoter, err := json.NewPromoter(name)
	if err != nil {
		return nil, err
	}
	return NewCollector(promoter)
}
