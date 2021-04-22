package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/pshvedko/yaml-metrics/collector/json"
	"reflect"
	"testing"
)

//func TestCollector_Collect(t *testing.T) {
//	type fields struct {
//		m map[string]Metric
//		p Promoter
//	}
//	type args struct {
//		metrics chan<- prometheus.Metric
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Collector{
//				m: tt.fields.m,
//				p: tt.fields.p,
//			}
//		})
//	}
//}
//
//func TestCollector_Describe(t *testing.T) {
//	type fields struct {
//		m map[string]Metric
//		p Promoter
//	}
//	type args struct {
//		descriptors chan<- *prometheus.Desc
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Collector{
//				m: tt.fields.m,
//				p: tt.fields.p,
//			}
//		})
//	}
//}

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
				t.Errorf("Map() got = %#v, want %#v", c, tt.want)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCollector(tt.args.promoter)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollector() got = %v, want %v", got, tt.want)
			}
		})
	}
}
