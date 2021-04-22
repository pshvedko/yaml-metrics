package collector

import (
	"github.com/pshvedko/yaml-metrics/collector/json"
	"reflect"
	"testing"
)

const jsonFile = "../metrics.json"

func TestNewCollectorJson(t *testing.T) {
	p, err := json.NewPromoter(jsonFile)
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
			args:    args{name: jsonFile},
			want:    &Collector{m: map[string]Metric{}, p: p},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCollectorJson(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCollectorJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollectorJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
