package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "tea",
		Price: 1,
		SKU:   "sas-ash-dfg",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
