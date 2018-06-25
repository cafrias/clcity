package fixtures

import (
	"testing"

	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/pkg/geo"
)

// TestKMLGenerate provides fixture data for test `TestKML_Generate`
func TestKMLGenerate(t *testing.T) *clcitybusapi.Empresa {
	fEmp := &clcitybusapi.Empresa{
		Codigo: 1,
		Lang:   "ESP",
		Nombre: "CityBus",
	}
	fLin := []*clcitybusapi.Linea{
		&clcitybusapi.Linea{
			Codigo:        1,
			Descripcion:   "Linea A",
			CodigoEntidad: 355,
			Empresa:       fEmp,
			Recorrido: &clcitybusapi.Recorrido{
				Puntos: []geo.Point{
					geo.Point{
						Lon: -67.6978612,
						Lat: -53.7828671,
					},
					geo.Point{
						Lon: -67.7069592,
						Lat: -53.7891804,
					},
					geo.Point{
						Lon: -67.7098131,
						Lat: -53.7877859,
					},
					geo.Point{
						Lon: -67.6999211,
						Lat: -53.7809526,
					},
					geo.Point{
						Lon: -67.6971745,
						Lat: -53.7823473,
					},
					geo.Point{
						Lon: -67.6976037,
						Lat: -53.7827149,
					},
				},
			},
		},
		&clcitybusapi.Linea{
			Codigo:        2,
			Descripcion:   "Linea B",
			CodigoEntidad: 356,
			Empresa:       fEmp,
			Recorrido: &clcitybusapi.Recorrido{
				Puntos: []geo.Point{
					geo.Point{
						Lon: -67.6980329,
						Lat: -53.7829812,
					},
					geo.Point{
						Lon: -67.7018738,
						Lat: -53.7856435,
					},
					geo.Point{
						Lon: -67.6988697,
						Lat: -53.7871901,
					},
					geo.Point{
						Lon: -67.6974535,
						Lat: -53.7879,
					},
					geo.Point{
						Lon: -67.6950932,
						Lat: -53.7862647,
					},
					geo.Point{
						Lon: -67.6978183,
						Lat: -53.7849463,
					},
					geo.Point{
						Lon: -67.6966166,
						Lat: -53.7840588,
					},
					geo.Point{
						Lon: -67.697947,
						Lat: -53.7832855,
					},
					geo.Point{
						Lon: -67.6970243,
						Lat: -53.7827783,
					},
					geo.Point{
						Lon: -67.6975393,
						Lat: -53.7826135,
					},
					geo.Point{
						Lon: -67.6978612,
						Lat: -53.7828671,
					},
				},
			},
		},
	}
	fPar := []*clcitybusapi.Parada{
		&clcitybusapi.Parada{
			Codigo: "RG001",
			Punto: geo.Point{
				Lon: -67.6980329,
				Lat: -53.7829812,
			},
		},
		&clcitybusapi.Parada{
			Codigo: "RG002",
			Punto: geo.Point{
				Lon: -67.7068949,
				Lat: -53.7857196,
			},
		},
		&clcitybusapi.Parada{
			Codigo: "RG003",
			Punto: geo.Point{
				Lon: -67.6970243,
				Lat: -53.7853012,
			},
		},
	}

	fEmp.Lineas = fLin
	fEmp.Paradas = fPar

	return fEmp
}
