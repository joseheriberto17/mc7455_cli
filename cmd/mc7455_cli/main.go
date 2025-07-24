package main

import (
	"github.com/nebulaengineering/mc7455_cli/internal/AtCommand"
	"github.com/nebulaengineering/mc7455_cli/internal/ui"
)

// var commands = []string{
// 	"AT!PCTEMP?\r",
// 	"AT+CFUN?\r",
// 	"AT+WANT?\r",
// 	// "AT+WANT=1\r",
// 	"AT!GSTATUS?\r",
// 	"AT+CREG?\r",
// 	"AT+CGREG?\r",
// 	"AT!GETBAND?\r",
// 	"AT+CGDCONT?\r",
// 	"AT!GPSSTATUS?\r",
// 	"AT!CUSTOM?\r",
// 	"AT!GPSSATINFO?\r",
// 	"AT!ERR\r",
// 	// "AT!ERR=0\r",
// 	// "AT!GPSEND=0,255\r",
// 	// "AT!GPSFIX=1,60,10\r",
// 	"AT!GPSLOC?\r",
// }

var SeqCmd = []string{
	AtCommand.GetVersionDef.Cmd, // "AT!PCVOLT?\r"
	AtCommand.PcVoltDef.Cmd,     // "AT!PCVOLT?\r"
	AtCommand.PcTempDef.Cmd,     // "AT!PCTEMP?\r"
	AtCommand.CFunDef.Cmd,       // "AT+CFUN?\r"
	AtCommand.WantDef.Cmd,       // "AT+WANT?\r"
	AtCommand.CRegDef.Cmd,       // "AT+CREG?\r"
	AtCommand.CGRegDef.Cmd,      // "AT+CGREG?\r"
	AtCommand.GetBandDef.Cmd,    // "AT!GETBAND?\r"
	AtCommand.CGDContDef.Cmd,    // "AT+CGDCONT?\r"
	// Añade aquí los que falten…
}

// ---------------- Información del MÓDEM ----------------
var rowsModem = [][]string{
	{"Versión de firmware (esperado SWI9X30C_02.38.00.00)", "02.36.03.00", "✔"},
	{"Voltaje de alimentación (3.3-4.1 V)", "3.71 V", "✔"},
	{"Temperatura del módem (-20…85 °C)", "34 °C", "✔"},
	{"Contador de errores módem (!ERR)", "0", "✔"},
	{"-----------------------------", "-------", "-------"},
	{"Registrado CS (voz)", "Sí", "✔"},
	{"Registrado PS (datos)", "Sí", "✔"},
	{"Tecnología / Banda", "LTE B7", "✔"},
	{"Agregación de portadoras (CA)", "B7+B28", "✔"},
	{"RSRP (≥ -110 dBm)", "-95 dBm", "✔"},
	{"SINR (≥ 5 dB)", "14.3 dB", "✔"},
	{"RSSI (≥ -85 dBm)", "-69 dBm", "✔"},
	{"-----------------------------", "-------", "-------"},
	{"Alimentación antena GPS (+WANT)", "ON", "✔"},
	{"GPS habilitado", "Sí", "✔"},
	{"Satélites en vista (≥ 4)", "6", "✔"},
	{"Sats con SNR > 25 dB (≥ 4)", "5", "✔"},
	{"Tipo de fix (2D/3D)", "3D", "✔"},
	{"DOP horizontal (≤ 2.5)", "1.2", "✔"},
}

func main() {

	tbl := ui.Frontend{
		Title: "Diagnóstico MC7455 - visión general",
		Head:  []string{"Chequeo", "Valor", "OK?"},
	}
	tbl.RenderTitle()
	tbl.Render(rowsModem)

	// -------------------------------------------------------------------

	// serialport := PortCustom.NewSerialPort("/dev/ttyMODEM", 115200)

	// //abrir puerto (maneja el error internamente)
	// serialport.OpenePort()

	// // // for  i  := 0; i < len(commands); i++ {
	// data, _ := serialport.SendCommand([]byte(SeqCmd[0]), 2*time.Second)
	// fmt.Print("\nsalida: ")
	// fmt.Printf("%q", data)

	// getversion, lencht, _ := AtCommand.GetVersionDef.Extract(data)

	// fmt.Printf("len: %d status %s", lencht, getversion.Version)

	// }
	// -------------------------------------------------------
}
