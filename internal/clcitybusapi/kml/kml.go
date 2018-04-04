package kml

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

func Generate(e *clcitybusapi.Empresa, path string) error {
	// Check Empresa meets requirements
	if len(e.Paradas) == 0 {
		return ErrNoParadas
	}
	if len(e.Lineas) == 0 {
		return ErrNoLineas
	}

	return nil
}
