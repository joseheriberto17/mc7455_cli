package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/nebulaengineering/mc7455_cli/internal/AtCommand"
	"github.com/nebulaengineering/mc7455_cli/internal/PortCustom"
	"github.com/nebulaengineering/mc7455_cli/internal/ui"
)

var reHDOP = regexp.MustCompile(`^\$(?:GP|GN)GGA,([^,]*,){7}([^,]*)`)

func main() {

	// -------------------------------------------------------------------
	fmt.Print("\033[H\033[2J")
	serialport := PortCustom.NewSerialPort("/dev/ttyMODEM", 115200)
	gpsport := PortCustom.NewSerialPort("/dev/ttyGPS", 9600)

	gpsport.OpenePort() //abrir puerto GPS (maneja el error internamente)
	GpsInfo, _ := gpsport.ReadPrefix(2 * time.Second)
	gpsport.ClosePort() //cerrar puerto GPS

	nmea := GpsInfo
	// fmt.Println(nmea)

	if nmea == "" {
		fmt.Println("No se recibió información del GPS")
		return
	}
	var hdop float64 = 0

	m := reHDOP.FindStringSubmatch(nmea)
	if m != nil {
		hdopStr := m[2]
		hdop, _ = strconv.ParseFloat(hdopStr, 64)
	}

	// test
	// fmt.Print(hdop)

	//abrir puerto (maneja el error internamente)
	serialport.OpenePort()

	var tabla [][]string // aquí guardaremos las filas

	for i, c := range AtCommand.SeqCmd {
		raw, _ := serialport.SendCommand([]byte(c.Cmd()), 2*time.Second)
		// fmt.Printf("Comando: %q\n", c.Cmd())
		dato, _ := c.Run(raw) // dato implementa raws

		// ---- construir la fila de 3 columnas ------------
		fila := []string{
			dato[0].Label,
			dato[0].Value,
			fmt.Sprint(dato[0].OK),
		}
		// ---- añadir la fila separadores a la tabla -------------------
		if i == 5 || i == 12 {
			tabla = append(tabla, []string{"-------------------------", "-------", "------"}) // fila separadora
		}

		tabla = append(tabla, fila)
	}
	// valor HDOP

	tabla = append(tabla, []string{"Valor HDOP (< 1.5)", fmt.Sprintf("%01.1f", hdop), strconv.FormatBool(hdop < 1.5 && hdop != 0.0)}) // fila separadora

	// -------------------------------------------------------

	tbl := ui.Frontend{
		Title: "Diagnóstico MC7455 - visión general",
		Head:  []string{"Chequeo", "Valor", "OK?"},
	}
	tbl.RenderTitle()
	tbl.Render(tabla)

}
