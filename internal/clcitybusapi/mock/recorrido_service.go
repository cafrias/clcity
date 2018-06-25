package mock

import (
	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/internal/clcitybusapi/client"
)

var _ clcitybusapi.RecorridoService = &RecorridoService{}

// RecorridoService mock implementation of client.RecorridoService
type RecorridoService struct {
	*client.RecorridoService
	RecorridoDeLineaSpy *Spy
}

// RecorridoDeLinea mock implementation of RecorridoDeLinea for RecorridoService.
func (s *RecorridoService) RecorridoDeLinea(l *clcitybusapi.Linea) (*clcitybusapi.Recorrido, error) {
	s.RecorridoDeLineaSpy.Invoked = true
	call := s.RecorridoDeLineaSpy.Calls
	s.RecorridoDeLineaSpy.Args = append(s.RecorridoDeLineaSpy.Args, []interface{}{
		l,
	})

	ret1, _ := s.RecorridoDeLineaSpy.Ret[call][0].(*clcitybusapi.Recorrido)
	ret2, _ := s.RecorridoDeLineaSpy.Ret[call][1].(error)

	s.RecorridoDeLineaSpy.Calls++

	return ret1, ret2
}
