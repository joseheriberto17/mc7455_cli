package main

import (
	"fmt"
	"time"

	"github.com/nebulaengineering/mc7455_cli/internal/AtCommand"
	"github.com/nebulaengineering/mc7455_cli/internal/PortCustom"
	"github.com/nebulaengineering/mc7455_cli/internal/ui"
)

//	var commands = []string{
//		"AT!PCTEMP?\r",
//		"AT+CFUN?\r",
//		"AT+WANT?\r",
//		// "AT+WANT=1\r",
//		"AT!GSTATUS?\r",
//		"AT+CREG?\r",
//		"AT+CGREG?\r",
//		"AT!GETBAND?\r",
//		"AT+CGDCONT?\r",
//		"AT!GPSSTATUS?\r",
//		"AT!CUSTOM?\r",
//		"AT!GPSSATINFO?\r",
//		"AT!ERR\r",
//		// "AT!ERR=0\r",
//		// "AT!GPSEND=0,255\r",
//		// "AT!GPSFIX=1,60,10\r",
//		"AT!GPSLOC?\r",
//	}
func main() {

	// -------------------------------------------------------------------

	serialport := PortCustom.NewSerialPort("/dev/ttyMODEM", 115200)

	//abrir puerto (maneja el error internamente)
	serialport.OpenePort()

	var tabla [][]string // aquí guardaremos las filas

	for i, c := range AtCommand.SeqCmd {
		raw, _ := serialport.SendCommand([]byte(c.Cmd()), 2*time.Second)
		// fmt.Printf("Comando: %q\n", c.Cmd())
		dato, _ := c.Run(raw) // dato implementa raws
		// if err != nil {
		// 	fmt.Println("parse error:", err)
		// 	continue
		// }

		// ---- construir la fila de 3 columnas ------------
		fila := []string{
			dato.Label(),
			dato.Value(),
			fmt.Sprint(dato.OK()),
		}
		// ---- añadir la fila separadores a la tabla -------------------
		if i == 4 || i == 8 {
			tabla = append(tabla, []string{"-------------------------", "-------", "------"}) // fila separadora
		}

		tabla = append(tabla, fila)
	}

	// -------------------------------------------------------

	tbl := ui.Frontend{
		Title: "Diagnóstico MC7455 - visión general",
		Head:  []string{"Chequeo", "Valor", "OK?"},
	}
	tbl.RenderTitle()
	tbl.Render(tabla)

}
