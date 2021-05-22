package data

import "testing"

func TestValidation(t *testing.T) {

	p := Game{
		Name:      "looll",
		Publisher: "s22",
		Rating:    -1,
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
