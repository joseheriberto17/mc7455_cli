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
	port, err := serial.Open("/dev/ttyMODEM", mode)
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
		fmt.Println("\t1.  Habilitar permisos para comando restringidos")
		fmt.Println("\t2.  Realizar un Reinicio completo del modem")
		fmt.Println("\t3.  Realizar un Reinicio parcial (solo software) del modem")

		fmt.Println("\nEstado General del módem y alimentacion:")
		fmt.Println("\t4.  Verificar si el modem esta Operativo (General)")
		fmt.Println("\t5.  Verificar voltaje de alimentacion del modem (3,3-4,1 V es OK)")
		fmt.Println("\t6.  Verificar temperatura de modem")
		fmt.Println("\t7.  Verificar nivel fucional de modem en general.")

		fmt.Println("\nAlimentación de la antena GPS:")
		fmt.Println("\t8.  Verificar la alimentacion de 3.3V del modem hacia la antena GPS")
		fmt.Println("\t9.  Habilitar la alimentacion de 3.3V del modem hacia la antena GPS")

		fmt.Println("\nInformacion de la  RED MOVIL:")
		fmt.Println("\t10.  Verificar el estado general en la RED MOVIL")
		fmt.Println("\t11. Verificar estado de registro en red de datos (PS)") // cgreg
		fmt.Println("\t12. Verificar estado de registro en red clásica (CS).") // cgreg
		fmt.Println("\t13. Muestra la banda o agregación LTE/NR en servicio")
		fmt.Println("\t14. Lista perfiles PDP (CID, tipo IP, APN)")

		fmt.Println("\nInformacion de GPS")
		fmt.Println("\t15. Verificar si el GPS del modem esta Operativo")
		fmt.Println("\t16. Verificar que Capacidades del modem para el GPS estan habilitados")
		fmt.Println("\t17. Verificar numero de Satélites y su señal SNR")

		fmt.Println("\nRegistros de errores")
		fmt.Println("\t18. ver logs de los errores registrados")
		fmt.Println("\t19. Borrar logs de los errores registrados")

		fmt.Println("\nIniciar fix compatible y monitorizar")
		fmt.Println("\t20. Iniciar búsqueda GPS (paso 1) ")
		fmt.Println("\t21. Iniciar búsqueda GPS (paso 2) ")
		fmt.Println("\t22. Determina posicion final del GPS")

		fmt.Println("\t23. Salir del programa")
		fmt.Print(">> ")

		input, _ := consoleReader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "23" {
			fmt.Println("Saliendo del programa.")
			break
		}

		// Mapa de comandos AT
		atCommands := map[string]string{
			// Acciones especiales
			"1": "AT!ENTERCND=\"A710\"\r",
			"2": "AT!RESET\r",
			"3": "AT+CFUN=1,1\r",
			// Estado General del módem y alimentacion
			"4": "AT!PCINFO?\r",
			"5": "AT!PCVOLT?\r",
			"6": "AT!PCTEMP?\r",
			"7": "AT+CFUN?\r",
			// Alimentación de la antena GPS
			"8": "AT+WANT?\r",
			"9": "AT+WANT=1\r",
			// Informacion de la  RED MOVIL:
			"10": "AT!GSTATUS?\r",
			"11": "AT+CREG?\r",
			"12": "AT+CGREG?\r",
			"13": "AT!GETBAND?\r",
			"14": "AT+CGDCONT?\r",
			//informacion de GPS
			"15": "AT!GPSSTATUS?\r",
			"16": "AT!CUSTOM?\r",
			"17": "AT!GPSSATINFO?\r",
			// Registros de errores
			"18": "AT!ERR\r",
			"19": "AT!ERR=0\r",
			// Iniciar fix compatible y monitorizar
			"20": "AT!GPSEND=0,255\r",
			"21": "AT!GPSFIX=1,60,10\r",
			"22": "AT!GPSLOC?\r",
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
