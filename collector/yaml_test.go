package collector

import (
	"github.com/pshvedko/yaml-metrics/collector/yaml"
	"reflect"
	"testing"
)

func TestNewCollectorYaml(t *testing.T) {
	p, err := yaml.NewPromoter(yamlFile)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *Collector
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Ok",
			args:    args{name: yamlFile},
			want:    &Collector{m: map[string]Metric{}, p: p},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCollectorYaml(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCollectorYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollectorYaml() got = %v, want %v", got, tt.want)
			}
		})
	}
}
