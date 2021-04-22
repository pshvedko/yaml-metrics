package collector

import (
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var ErrNilPromoter = errors.New("promoter is nil")

type Metric struct {
	t prometheus.ValueType
	v string
	k []string
	d *prometheus.Desc
}

type Promoter interface {
	Promote() map[string]interface{}
}

type Collector struct {
	m map[string]Metric
	p Promoter
}

func (c Collector) Describe(descriptors chan<- *prometheus.Desc) {
	for _, m := range c.m {
		descriptors <- m.d
	}
}

func (c Collector) Collect(metrics chan<- prometheus.Metric) {
	var a []string
	for k, v := range c.p.Promote() {
		if m, ok := c.m[k]; ok {
			switch v := v.(type) {
			case []interface{}:
				for _, v := range v {
					switch v := v.(type) {
					case map[interface{}]interface{}:
						for _, k := range m.k {
							if v, ok := v[k]; ok {
								switch v := v.(type) {
								case bool, string, int, float64, int64:
									a = append(a, fmt.Sprint(v))
									continue
								default:
									// FIXME
								}
							}
							a = append(a, "") // FIXME
						}
						if v, ok := v[m.v]; ok {
							switch v := v.(type) {
							case int:
								metrics <- prometheus.MustNewConstMetric(m.d, m.t, float64(v), a...)
							case int64:
								metrics <- prometheus.MustNewConstMetric(m.d, m.t, float64(v), a...)
							case float64:
								metrics <- prometheus.MustNewConstMetric(m.d, m.t, v, a...)
							default:
								// FIXME
							}
						}
					default:
						// FIXME
					}
					a = a[:0]
				}
			case int:
				metrics <- prometheus.MustNewConstMetric(m.d, m.t, float64(v))
			case int64:
				metrics <- prometheus.MustNewConstMetric(m.d, m.t, float64(v))
			case float64:
				metrics <- prometheus.MustNewConstMetric(m.d, m.t, v)
			default:
				// FIXME
			}
		}
	}
}

func (c Collector) Map(name, help, valueKey string, valueType prometheus.ValueType, keys []string, labels prometheus.Labels) {
	c.m[name] = Metric{
		t: valueType,
		v: valueKey,
		k: keys,
		d: prometheus.NewDesc(name, help, keys, labels),
	}
}

func NewCollector(promoter Promoter) (*Collector, error) {
	if promoter != nil {
		return &Collector{m: map[string]Metric{}, p: promoter}, nil
	}
	return nil, ErrNilPromoter
}
