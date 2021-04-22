package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/pshvedko/yaml-metrics/collector/json"
	"github.com/pshvedko/yaml-metrics/collector/yaml"
	"reflect"
	"testing"
)

const jsonFile = "../metrics.json"
const yamlFile = "../metrics.yaml"

func TestCollector_Collect(t *testing.T) {
	p, err := yaml.NewPromoter(yamlFile)
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		m map[string]Metric
		p Promoter
	}
	type args struct {
		metrics chan prometheus.Metric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "Ok",
			fields: fields{
				m: map[string]Metric{
					"currencies": {
						t: prometheus.GaugeValue,
						v: "value",
						k: []string{"name"},
						d: prometheus.NewDesc("currencies", "", []string{"name"}, prometheus.Labels{"from": "test"})},
				},
				p: p,
			},
			args: args{metrics: make(chan prometheus.Metric)},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Collector{
				m: tt.fields.m,
				p: tt.fields.p,
			}
			go func(m chan<- prometheus.Metric) {
				c.Collect(m)
				close(m)
			}(tt.args.metrics)
			for d := range tt.args.metrics {
				tt.want--
				if tt.want < 0 {
					t.Errorf("Collect() got = %#v, want nothing", d)
				}
			}
		})
	}
}

func TestCollector_Describe(t *testing.T) {
	type fields struct {
		m map[string]Metric
		p Promoter
	}
	type args struct {
		descriptors chan *prometheus.Desc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "Ok",
			fields: fields{
				m: map[string]Metric{
					"currencies": {
						t: prometheus.GaugeValue,
						v: "value",
						k: []string{"name"},
						d: prometheus.NewDesc("currencies", "", []string{"name"}, prometheus.Labels{"from": "test"})},
				},
				p: nil,
			},
			args: args{descriptors: make(chan *prometheus.Desc)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Collector{
				m: tt.fields.m,
				p: tt.fields.p,
			}
			go func(d chan<- *prometheus.Desc) {
				c.Describe(d)
				close(d)
			}(tt.args.descriptors)
			for _, v := range tt.fields.m {
				d := <-tt.args.descriptors
				if !reflect.DeepEqual(d, v.d) {
					t.Errorf("Describe() got = %#v, want %#v", d, v.d)
				}
			}
			for d := range tt.args.descriptors {
				t.Errorf("Describe() got = %#v, want nothing", d)
			}
		})
	}
}

func TestCollector_Map(t *testing.T) {
	type fields struct {
		m map[string]Metric
		p Promoter
	}
	type args struct {
		name      string
		help      string
		valueKey  string
		valueType prometheus.ValueType
		keys      []string
		labels    prometheus.Labels
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Collector
	}{
		// TODO: Add test cases.
		{
			name: "Ok",
			fields: fields{
				m: map[string]Metric{},
				p: nil,
			},
			args: args{
				name:      "currencies",
				help:      "",
				valueKey:  "value",
				valueType: prometheus.GaugeValue,
				keys:      []string{"name"},
				labels:    prometheus.Labels{"from": "test"},
			},
			want: Collector{
				m: map[string]Metric{
					"currencies": {
						t: prometheus.GaugeValue,
						v: "value",
						k: []string{"name"},
						d: prometheus.NewDesc("currencies", "", []string{"name"}, prometheus.Labels{"from": "test"})},
				},
				p: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Collector{
				m: tt.fields.m,
				p: tt.fields.p,
			}
			c.Map(tt.args.name, tt.args.help, tt.args.valueKey, tt.args.valueType, tt.args.keys, tt.args.labels)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("Map() got = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestNewCollector(t *testing.T) {
	type args struct {
		promoter Promoter
	}
	tests := []struct {
		name    string
		args    args
		want    *Collector
		wantErr bool
		err     error
	}{
		// TODO: Add test cases.
		{
			name: "Ok",
			args: args{promoter: &json.Promoter{}},
			want: &Collector{
				m: map[string]Metric{},
				p: &json.Promoter{},
			},
			wantErr: false,
		},
		{
			name:    "Fail",
			args:    args{promoter: nil},
			want:    nil,
			wantErr: true,
			err:     ErrNilPromoter,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCollector(tt.args.promoter)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewCollector() err = %v, want %v", err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollector() got = %v, want %v", got, tt.want)
			}
		})
	}
}
