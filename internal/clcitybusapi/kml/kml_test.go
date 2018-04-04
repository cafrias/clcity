package kml_test

import (
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/kml"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

func TestKML_Generate_ErrNoLineas(t *testing.T) {
	emp := &clcitybusapi.Empresa{
		Paradas: []*clcitybusapi.Parada{
			&clcitybusapi.Parada{},
		},
	}

	err := kml.Generate(emp, "")
	if err != kml.ErrNoLineas {
		t.Fatalf("Unexpected error: %v\n", err)
	}
}

func TestKML_Generate_ErrNoParadas(t *testing.T) {
	emp := &clcitybusapi.Empresa{}

	err := kml.Generate(emp, "")
	if err != kml.ErrNoParadas {
		t.Fatalf("Unexpected error: %v\n", err)
	}
}
