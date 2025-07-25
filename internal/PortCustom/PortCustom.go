// Nivel puerto
// gestion de puerto serial para la transferecia de datos en bajo nivel
package PortCustom

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

type SerialPort struct {
	Name string
	mode *serial.Mode
	Port serial.Port
}

// New crea un nuevo objeto SerialPort con configuración por defecto
func NewSerialPort(name string, baud int) *SerialPort {
	return &SerialPort{
		Name: name,
		mode: &serial.Mode{
			BaudRate: baud,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
		},
	}
}

// Open abre el puerto serial con la configuración dada
func (dev *SerialPort) OpenePort() {
	port, err := serial.Open(dev.Name, dev.mode)
	if err != nil {
		log.Fatalf("error al abrir el puerto %s: %v", dev.Name, err)
	}
	dev.Port = port

	// Configurar tiempo de espera
	err = dev.Port.SetReadTimeout(2 * time.Second)
	if err != nil {
		log.Fatalf("error al definir un timeout en la lectura del puerto %s: %v", dev.Name, err)
	}

	// Limpiar buffers
	err = dev.Port.ResetInputBuffer()
	if err != nil {
		log.Fatalf("error al limpiar el buffer de entrada al puerto %s: %v", dev.Name, err)
	}
	err = dev.Port.ResetOutputBuffer()
	if err != nil {
		log.Fatalf("error al limpiar el buffer de salida al puerto %s: %v", dev.Name, err)
	}

	// // test
	// fmt.Printf("puerto abierto %s\n\r", dev.Name)
}

// cierra el puerto
func (dev *SerialPort) ClosePort() error {
	if dev.Port != nil {
		return dev.Port.Close()
	}
	return nil
}

// hace un write pero sin manejo de errores
func (dev *SerialPort) Write(cmd []byte) (err error) {

	// prepara el bufer de entrada antes de imprimir
	if err = dev.Port.ResetInputBuffer(); err != nil {
		return err
	}

	if _, err = dev.Port.Write(cmd); err != nil {
		return err
	}
	// asegura que de verdad se envien los datos
	if err = dev.Port.Drain(); err != nil {
		return err
	}

	// // test
	// fmt.Printf("escritura por el puerto: %s, %s\n\r", dev.Name, cmd)
	return err
}

// hace un write pero con manejo de errores
func (dev *SerialPort) Write_Error(cmd []byte) {
	// Limpiar buffer antes de enviar
	if err := dev.Port.ResetInputBuffer(); err != nil {
		log.Fatalf("Clean buffer port fail: %s ,err: %v", dev.Name, err)
	}

	if _, err := dev.Port.Write(cmd); err != nil {
		log.Fatalf("Write port fail: %s ,err: %v", dev.Name, err)
	}

	// Esperar a que realmente salga.
	if err := dev.Port.Drain(); err != nil {
		log.Fatalf("Drain port fail: %s ,err: %v", dev.Name, err)
	}
}

// ReadUntil lee hasta que encuentre "OK", "ERROR" o timeout
func (dev *SerialPort) ReadUntil(timeout time.Duration) (string, error) {
	reader := bufio.NewReader(dev.Port)
	dev.Port.SetReadTimeout(timeout)

	var sb strings.Builder
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("error de ReadString: %v\n", err)
			break
		}
		sb.WriteString(line)

		if strings.Contains(line, "OK") || strings.Contains(line, "ERROR") {
			break
		}
	}
	// fmt.Printf("lectura en crudo: %q\n", sb.String())
	return sb.String(), nil

}

// SendCommand envía un comando y devuelve la respuesta completa
func (dev *SerialPort) SendCommand(cmd []byte, timeout time.Duration) (string, error) {
	if err := dev.Write(cmd); err != nil {
		return "", fmt.Errorf("error al enviar comando: %w", err)
	}

	// // test
	// fmt.Printf("comando enviado por el puerto: %s\n\r", dev.Name)
	return dev.ReadUntil(timeout)
}
