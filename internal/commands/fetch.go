package commands

import (
	"fmt"
	"os"

	"github.com/friasdesign/clcity/internal/clcitybusapi"
	"github.com/friasdesign/clcity/internal/clcitybusapi/client"
)

func fetch(nEmp int, dumpPath string) (*clcitybusapi.Empresa, error) {
	if _, err := os.Stat(dumpPath); os.IsNotExist(err) {
		return nil, errDumpDirDoesntExist
	}
	cli := client.NewClient(nil, dumpPath)

	// Create empresa
	fmt.Println("Creating empresa ...")
	emp := clcitybusapi.NewEmpresa(nEmp)
	fmt.Println("Creating empresa ... DONE!")

	// Fetch lineas
	fmt.Println("Fetching lineas ...")
	lin, err := cli.LineaService().LineasPorEmpresa(emp)
	if err != nil {
		return nil, err
	}
	emp.Lineas = lin
	fmt.Println("Fetching lineas ... DONE!")

	// Fetch paradas
	fmt.Println("Fetching paradas ...")
	par, err := cli.ParadaService().ParadasPorEmpresa(emp)
	if err != nil {
		return nil, err
	}
	emp.Paradas = par
	fmt.Println("Fetching paradas ... DONE!")

	return emp, nil
}
