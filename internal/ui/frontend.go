package ui

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
)

// Frontend encapsula los datos necesarios para imprimir la tabla de diagnóstico
// en consola. El método Render genera una salida ASCII uniforme utilizando
// tablewriter.
//
//  📋 Ejemplo de uso desde main:
//
//      tbl := ui.Frontend{
//          Title: "Diagnóstico MC7455 – visión general",
//          Head:  []string{"Chequeo", "Valor", "OK?"},
//          Rows:  ui.BuildRows(rep),
//      }
//      tbl.Render()
//
// De esta forma `main.go` solo necesita proveer el título, encabezados y las filas
// construidas — toda la lógica para ordenar, formatear y pintar la tabla queda
// contenida en el paquete ui.
//
// Cualquier cambio en la estética o la agrupación de datos puede realizarse aquí
// sin tocar el resto de la aplicación.

type Frontend struct {
	Title string
	Head  []string
}

func (f Frontend) RenderTitle() {
	fmt.Println("────────────────────────────────────────────────────────────")
	fmt.Println(f.Title)
	fmt.Println("────────────────────────────────────────────────────────────")
}

// Render imprime la tabla en la salida estándar con un título centrado y un marco
// uniforme.
func (f Frontend) Render(Rows [][]string) {

	tbl := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewBlueprint()), // salida ASCII
	)
	tbl.Header(f.Head)
	tbl.Bulk(Rows)
	tbl.Render()
	fmt.Println("")
}
