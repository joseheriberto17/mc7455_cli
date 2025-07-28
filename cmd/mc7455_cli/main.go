package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	fmt.Print("\033[H\033[2J")
	serialport := PortCustom.NewSerialPort("/dev/ttyMODEM", 115200)

	//abrir puerto (maneja el error internamente)
	serialport.OpenePort()

	for {
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
			if i == 5 || i == 9 {
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

		// Crear lector de consola
		consoleReader := bufio.NewReader(os.Stdin)

		fmt.Println("\n\nAcciones del superusuario:")
		fmt.Println("\t1.  Habilitar permisos para comando restringidos")
		fmt.Println("\t2.  Realizar un Reinicio completo del modem")
		fmt.Println("\t3.  Realizar un Reinicio parcial (solo software) del modem")
		fmt.Println("\t4.  Habilitar la alimentacion de 3.3V del modem hacia la antena GPS")
		fmt.Println("\t5. Borrar logs de los errores registrados")

		fmt.Println("\nIniciar fix compatible y monitorizar")
		fmt.Println("\t6. Iniciar búsqueda GPS (paso 1) ")
		fmt.Println("\t7. Iniciar búsqueda GPS (paso 2) ")
		fmt.Println("\t8. Determina posicion final del GPS")

		fmt.Println("\t9. Salir del programa")
		fmt.Print(">> ")

		input, _ := consoleReader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "9" {
			fmt.Println("Saliendo del programa.")
			break
		}

		// Mapa de comandos AT
		atCommands := map[string]string{

			"1": "AT!ENTERCND=\"A710\"\r",
			"2": "AT!RESET\r",
			"3": "AT+CFUN=1,1\r",
			"4": "AT+WANT=1\r",
			"5": "AT!ERR=0\r",
			"6": "AT!GPSEND=0,255\r",
			"7": "AT!GPSFIX=1,60,10\r",
			"8": "AT!GPSLOC?\r",
		}
		cmd, ok := atCommands[input]

		// atiende una entrada no valida
		if !ok {
			fmt.Println("Opción no válida.")
			continue
		}

		// Enviar comando
		_, err := serialport.SendCommand([]byte(cmd), 2*time.Second)
		if err != nil {
			log.Println("Error al escribir en el puerto:", err)
			continue
		}

		fmt.Print("Presione ENTER para continuar: ")
		consoleReader.ReadString('\n')
	}

}
