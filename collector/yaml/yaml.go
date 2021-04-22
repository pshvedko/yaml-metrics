package yaml

import (
	"gopkg.in/yaml.v2"
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
	err = yaml.NewDecoder(f).Decode(&m)
	if err != nil {
		return nil, err
	}
	return &Promoter{m: m}, nil
}
