package json

import (
	"encoding/json"
	"os"
)

type Promoter struct {
	m map[string]interface{}
}

func (p Promoter) Promote() map[string]interface{} {
	return p.m
}

func NewPromoter(name string) (*Promoter, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var m map[string]interface{}
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		return nil, err
	}
	for k, v := range m {
		switch v := v.(type) {
		case map[string]interface{}:
			m[k] = hashmap(v)
		case []interface{}:
			m[k] = array(v)
		}
	}
	return &Promoter{m: m}, nil
}

func array(a []interface{}) []interface{} {
	for i, v := range a {
		switch v := v.(type) {
		case map[string]interface{}:
			a[i] = hashmap(v)
		case []interface{}:
			a[i] = array(v)
		}
	}
	return a
}

func hashmap(m map[string]interface{}) map[interface{}]interface{} {
	h := map[interface{}]interface{}{}
	for k, v := range m {
		switch v := v.(type) {
		case map[string]interface{}:
			h[k] = hashmap(v)
		case []interface{}:
			h[k] = array(v)
		default:
			h[k] = v
		}
	}
	return h
}
