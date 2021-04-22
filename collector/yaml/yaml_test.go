package yaml

import (
	"reflect"
	"testing"
)

const yamlFile = "../../metrics.yaml"

func TestNewPromoter(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *Promoter
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Ok",
			args: args{name: yamlFile},
			want: &Promoter{m: map[string]interface{}{
				"currencies": []interface{}{
					map[interface{}]interface{}{"name": "usd", "value": 70.5},
					map[interface{}]interface{}{"name": "eur", "value": 80},
					map[interface{}]interface{}{"name": "rur", "value": 1},
				}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPromoter(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPromoter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPromoter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromoter_Promote(t *testing.T) {
	type fields struct {
		m map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
		{
			name: "Ok",
			fields: fields{m: map[string]interface{}{
				"currencies": []interface{}{
					map[interface{}]interface{}{"name": "usd", "value": 70.5},
					map[interface{}]interface{}{"name": "eur", "value": 80},
					map[interface{}]interface{}{"name": "rur", "value": 1},
				}}},
			want: map[string]interface{}{
				"currencies": []interface{}{
					map[interface{}]interface{}{"name": "usd", "value": 70.5},
					map[interface{}]interface{}{"name": "eur", "value": 80},
					map[interface{}]interface{}{"name": "rur", "value": 1},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Promoter{
				m: tt.fields.m,
			}
			if got := p.Promote(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Promote() = %v, want %v", got, tt.want)
			}
		})
	}
}
