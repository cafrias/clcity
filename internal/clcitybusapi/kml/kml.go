package kml

import (
	"fmt"
	"os"

	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/twpayne/go-kml"
)

// Generate generates a KML file for all 'Parada' and 'Linea.Recorrido' of given 'Empresa'.
func Generate(e *clcitybusapi.Empresa, path string) error {
	// Check Empresa meets requirements
	if len(e.Paradas) == 0 {
		fmt.Printf("WARNING: Empresa id '%v', doesn't have any Parada.\n", e.Codigo)
	}
	if len(e.Lineas) == 0 {
		return ErrNoLineas
	}

	var linKML []kml.Element
	for _, lin := range e.Lineas {
		var coord []kml.Coordinate
		for _, point := range lin.Recorrido.Puntos {
			c := kml.Coordinate{
				Lat: point.Lat,
				Lon: point.Lon,
			}
			coord = append(coord, c)
		}
		folder := kml.Folder(
			kml.Name(lin.Descripcion),
			kml.Placemark(
				kml.Name(lin.Descripcion),
				kml.Style(
					kml.LineStyle(
						kml.Color(lin.Color),
						kml.Width(2.5),
					),
				),
				kml.LineString(
					kml.Tessellate(true),
					kml.Coordinates(coord...),
				),
			),
		)

		linKML = append(linKML, folder)
	}

	parKML := []kml.Element{
		kml.Name("Paradas"),
	}
	for _, par := range e.Paradas {
		place := kml.Placemark(
			kml.Name(par.Codigo),
			kml.Point(
				kml.Coordinates(
					kml.Coordinate{
						Lon: par.Punto.Lon,
						Lat: par.Punto.Lat,
					},
				),
			),
		)
		parKML = append(parKML, place)
	}
	parFolderKML := kml.Folder(parKML...)

	docChildren := []kml.Element{
		kml.Name(e.Nombre),
		kml.Description(""),
	}
	docChildren = append(docChildren, linKML...)
	docChildren = append(docChildren, parFolderKML)

	baseKML := kml.KML(kml.Document(docChildren...))

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	err = baseKML.WriteIndent(f, "", "  ")

	return err
}
