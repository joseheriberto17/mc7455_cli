package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.bug.st/serial"
)

func main() {

	// Abrir el puerto objetivo (ajustar según tu sistema)
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("/dev/ttyUSB2", mode)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Limpiar el buffer antes de comenzar
	port.ResetInputBuffer()
	port.SetReadTimeout(2 * time.Second)

	// Crear lector de consola
	consoleReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Herramienta de Diagnostico y configuracion de GPS")
		fmt.Println("\nAcciones especiales")
		fmt.Println("\t1. Habilitar permisos para comando restringidos")
		fmt.Println("\t2. Reiniciar el modem")

		fmt.Println("\nEstado del módem y alimentacion:")
		fmt.Println("\t3. Verificar si el modem esta Operativo (General)")
		fmt.Println("\t4. Verificar voltaje de alimentacion del modem")
		fmt.Println("\t5. Verificar temperatura de modem")

		fmt.Println("\nAlimentación de la antena:")
		fmt.Println("\t6. Habilitar la alimentacion de modem hacia la antena GPS")
		fmt.Println("\t7. Verificar la alimentacion de modem hacia la antena GPS")

		fmt.Println("\nInformacion de GPS")
		fmt.Println("\t8. Verificar si el GPS del modem esta Operativo")
		fmt.Println("\t9. Verificar que Capacidades del modem para el GPS estan habilitados")
		fmt.Println("\t10. Verificar numero de Satélites y su señal SNR")

		fmt.Println("\nRegistros de error")
		fmt.Println("\t11. ver logs de los errores registrados")
		fmt.Println("\t12. Borrar logs de los errores registrados")

		fmt.Println("\nIniciar fix compatible y monitorizar")
		fmt.Println("\t13. Iniciar búsqueda GPS")
		fmt.Println("\t14. Determina posicion final del GPS")

		fmt.Println("\t15. Salir del programa")
		fmt.Print(">> ")

		input, _ := consoleReader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "15" {
			fmt.Println("Saliendo del programa.")
			break
		}

		// Mapa de comandos AT
		atCommands := map[string]string{
			"1":  "AT!ENTERCND=\"A710\"\r",
			"2":  "AT!RESET\r",
			"3":  "AT!PCINFO?\r",
			"4":  "AT!PCVOLT?\r",
			"5":  "AT!PCTEMP?\r",
			"6":  "AT+WANT=1\r",
			"7":  "AT+WANT?\r",
			"8":  "AT!GPSSTATUS?\r",
			"9":  "AT!CUSTOM?\r",
			"10": "AT!GPSSATINFO?\r",
			"11": "AT!ERR\r",
			"12": "AT!ERR=0\r",
			"13": "AT!GPSEND=0,255\r",
			"14": "AT!GPSFIX=1,60,10\r",
			"15": "AT!GPSLOC?\r",
		}
		cmd, ok := atCommands[input]

		// atiende una entrada no valida
		if !ok {
			fmt.Println("Opción no válida.")
			continue
		}

		// Limpiar buffer antes de enviar
		port.ResetInputBuffer()

		// Enviar comando
		_, err := port.Write([]byte(cmd))
		if err != nil {
			log.Println("Error al escribir en el puerto:", err)
			continue
		}

		// Leer respuesta
		scanner := bufio.NewScanner(port)
		scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if i := strings.IndexAny(string(data), "\r\n"); i >= 0 {
				return i + 1, data[:i], nil
			}
			return 0, nil, nil
		})

		// Antes de leer espera 2 segundos
		fmt.Println("Respuesta del dispositivo:")
		timeout := time.After(2 * time.Second)

	ReadingLoop:
		for {
			select {
			case <-timeout:
				break ReadingLoop
			default:
				if scanner.Scan() {
					line := strings.TrimSpace(scanner.Text())
					if line != "" {
						fmt.Println(line)
						if line == "OK" {
							break ReadingLoop
						}
					}
				}
			}
		}
		// atiende el error de lectura
		if scanner.Err() != nil {
			log.Println("Error leyendo desde el puerto:", scanner.Err())
		}

		fmt.Print("Presione ENTER para continuar: ")
		consoleReader.ReadString('\n')
	}
}
