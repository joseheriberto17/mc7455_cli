package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

func main() {

	//Lista de puertos disponibles
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Entra a puerto objetivo
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

	// 2. Limpia el buffer de entrada por si había datos anteriores
	port.ResetInputBuffer()

	// 3. Envía AT con solo <CR>
	_, _ = port.Write([]byte("AT!GSTATUS?\r"))

	// 4. Lee con un scanner por líneas hasta ver OK
	sc := bufio.NewScanner(port)
	sc.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// divide por CR/LF
		if i := strings.IndexAny(string(data), "\r\n"); i >= 0 {
			return i + 1, data[:i], nil
		}
		return
	})
	port.SetReadTimeout(2 * time.Second)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		} // ignora líneas vacías
		fmt.Printf("RX: %q\n", line)
		if line == "OK" { // imprime hasta que detecte un ok ,falta tratar error
			break
		}
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}

}
