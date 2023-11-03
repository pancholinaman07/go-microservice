package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "naman",
		Price: 1.0,
		SKU:   "abc-abc-abc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
