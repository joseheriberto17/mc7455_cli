// file ListAtCommad.go

package AtCommand

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// En este archivo se definen los comandos AT que se ejecutarán y se extraeran los datos de sus respuestas.

// se genera const <nombre>Pattern para cada comando AT
// y se genera una variable <nombre>Def que es un ATCommandDef[T] con
// el comando AT, el patrón y la función de parseo correspondiente

// Getversion representa la versión de firmware del módem.
type Getversion struct{ Version string }

const getversionPattern = "AT+GMR\r\r\n%s "

var GetVersionDef = ATCommandDef[Getversion]{
	Cmd:     "AT+GMR\r",
	Pattern: getversionPattern,
	Parse: func(r string) (Getversion, error) {
		var out Getversion
		n, _ := fmt.Sscanf(r, getversionPattern,
			&out.Version)

		if condition := n == 0; condition {
			out.Version = "Null"
		}
		return out, nil
	},
}

// PcVolt representa el voltaje de alimentación del módem.
type PcVolt struct {
	State  string
	MilliV int
	ADC    int
}

const pcvoltPattern = "AT!PCVOLT?\r\r\nVolt state: %s\r\nPower supply voltage: %d mV (ADC: %d)"

var PcVoltDef = ATCommandDef[PcVolt]{
	Cmd:     "AT!PCVOLT?\r",
	Pattern: pcvoltPattern,
	Parse: func(resp string) (PcVolt, error) {
		var out PcVolt
		n, _ := fmt.Sscanf(resp, pcvoltPattern,
			&out.State, &out.MilliV, &out.ADC)

		if condition := n == 0; condition {
			out.State = "Null"
			out.MilliV = 0
			out.ADC = 0
		}
		return out, nil
	},
}

// PcTemp representa la temperatura del módem.
type PcTemp struct {
	State string
	Temp  float32
}

const pctempPattern = "AT!PCTEMP?\r\r\nTemp state: %s\r\nTemperature: %f C"

var PcTempDef = ATCommandDef[PcTemp]{
	Cmd:     "AT!PCTEMP?\r",
	Pattern: pctempPattern,
	Parse: func(resp string) (PcTemp, error) {
		var out PcTemp
		n, _ := fmt.Sscanf(resp, pctempPattern,
			&out.State, &out.Temp)

		if condition := n == 0; condition {
			out.State = "Null"
			out.Temp = 0
		}
		return out, nil
	},
}

// // ContErr representa el contador de errores del módem.
// type ContErr struct {
// 	value int
// }

// func (e ContErr) Label() string { return "Contador de errores módem (!ERR)" }
// func (e ContErr) Value() string { return fmt.Sprintf("%d", e.value) }
// func (e ContErr) OK() bool      { return e.value == 0 }

// const errPattern = "AT!ERR\r\r\n%02d   %02x %s %05d"

// var ContErrDef = ATCommandDef[ContErr]{
// 	Cmd:     "AT!ERR?\r",
// 	Pattern: errPattern,
// 	Parse: func(resp string) (ContErr, error) {
// 		var out ContErr
// 		n, _ := fmt.Sscanf(resp, errPattern,
// 			&out.value)

// 		if condition := n == 0; condition {
// 			out.value = 0
// 		}
// 		return out, nil
// 	},
// }

// CFun representa el modo de operatividad general del módem.
type CFun struct{ Mode int }

const cfunPattern = "AT+CFUN?\r\r\n+CFUN: %d"

var CFunDef = ATCommandDef[CFun]{
	Cmd:     "AT+CFUN?\r",
	Pattern: cfunPattern,
	Parse: func(r string) (CFun, error) {
		var out CFun
		n, _ := fmt.Sscanf(r, cfunPattern, &out.Mode)

		if condition := n == 0; condition {
			out.Mode = 0
		}
		return out, nil
	},
}

// Want representa la alimentación de la antena GPS del módem.
type Want struct{ value int }

const wantPattern = "AT+WANT?\r\r\n+WANT: %d"

var WantDef = ATCommandDef[Want]{
	Cmd:     "AT+WANT?\r",
	Pattern: wantPattern,
	Parse: func(r string) (Want, error) {
		var out Want
		n, _ := fmt.Sscanf(r, wantPattern, &out.value)

		if condition := n == 0; condition {
			out.value = 0
		}
		return out, nil
	},
}

// CReg representa el estado de registro del módem en la red celular (voz).
type CReg struct {
	N    int
	Stat int
}

const cgregPattern = "AT+CGREG?\r\r\n+CGREG: %d,%d"

var CGRegDef = ATCommandDef[CGReg]{
	Cmd:     "AT+CGREG?\r",
	Pattern: cgregPattern,
	Parse: func(r string) (CGReg, error) {
		var out CGReg
		n, _ := fmt.Sscanf(r, cgregPattern, &out.N, &out.Stat)

		if condition := n == 0; condition {
			out.N = 0
			out.Stat = 0
		}

		return out, nil
	},
}

// CGReg representa el estado de registro del módem en la red celular (datos).
type CGReg struct {
	N    int
	Stat int
}

const cregPattern = "AT+CREG?\r\r\n+CREG: %d,%d"

var CRegDef = ATCommandDef[CReg]{
	Cmd:     "AT+CREG?\r",
	Pattern: cregPattern,
	Parse: func(r string) (CReg, error) {
		var out CReg
		n, _ := fmt.Sscanf(r, cregPattern, &out.N, &out.Stat)

		if condition := n == 0; condition {
			out.N = 0
			out.Stat = 0
		}

		return out, nil
	},
}

// GetBand representa la banda activa del módem.
type GetBand struct{ Band string }

const getbandPattern = "AT!GETBAND?\r\r\n!GETBAND: %s" // descripción de banda activa  :contentReference[oaicite:2]{index=2}

var GetBandDef = ATCommandDef[GetBand]{
	Cmd:     "AT!GETBAND?\r",
	Pattern: getbandPattern,
	Parse: func(r string) (GetBand, error) {
		var out GetBand
		n, _ := fmt.Sscanf(r, getbandPattern, &out.Band)

		if condition := n == 0; condition {
			out.Band = "Null"
		}

		return out, nil
	},
}

// CGDCont representa un contexto PDP (Packet Data Protocol) del módem.
type CGDCont struct {
	APN string
}

var CGDContDef = ATCommandDef[CGDCont]{
	Cmd: "AT+CGDCONT?\r",
	Parse: func(resp string) (CGDCont, error) {
		var out CGDCont
		// Captura el 3er campo entre comillas (el APN) de líneas +CGDCONT:
		var reCGDCONT_APN = regexp.MustCompile(`\+CGDCONT:\s*\d+,"[^"]*","([^"]+)"`)

		if m := reCGDCONT_APN.FindStringSubmatch(resp); m != nil {
			out.APN = m[1]
		}
		return out, nil
	},
}

// CGDCont representa un contexto PDP (Packet Data Protocol) del módem.
type GPSAtInfo struct {
	InView int
	AvgSNR float64
	Count  int
}

const gpssatinfoPattern = "AT!GPSSATINFO?\r\r\nSatellites in view:  %d"

var GPSAtInfoDef = ATCommandDef[GPSAtInfo]{
	Cmd:     "AT!GPSSATINFO?\r",
	Pattern: gpssatinfoPattern,
	Parse: func(resp string) (GPSAtInfo, error) {
		var out GPSAtInfo
		var (
			reSatCount = regexp.MustCompile(`Satellites in view:\s*(\d+)`)
			reSNR      = regexp.MustCompile(`SNR:\s*(\d+)`)
		)
		//
		if m := reSatCount.FindStringSubmatch(resp); m != nil {
			out.InView, _ = strconv.Atoi(m[1])
		}

		snrs := reSNR.FindAllStringSubmatch(resp, -1)
		var sum int
		for _, m := range snrs {
			n, err := strconv.Atoi(m[1])
			if err != nil {
				continue
			}
			sum += n
			out.Count++
		}
		if out.Count > 0 {
			out.AvgSNR = float64(sum) / float64(out.Count)
		}
		return out, nil
	},
}

// GStatus representa el estado del módem según el comando AT!GSTATUS.
type GStatus struct {
	Time   int
	Temp   int
	Creset int
	Mode   string
	// 	State   string
	// 	PSState string
	// 	LTEBand string
}

const gstatusPattern = "AT!GSTATUS?\r\r\n!GSTATUS: \r\nCurrent Time:  %d\t\tTemperature: %d\r\nReset Counter: %d\t\tMode:        %s         " //\r\nSystem mode:   %s        \tPS state:    %s     \r\nLTE band:      %s     \t\tLTE bw:      %d MHz  \r\n" //LTE Rx chan:   %d\t\tLTE Tx chan: %d\r\nLTE CA state:  %s\r\nEMM state:     %s     \t%s \r\nRRC state:     %s  \r\nIMS reg state: %s  \t\t\r\n\r\nPCC RxM RSSI:  %d\t\tRSRP (dBm):  %d\r\nPCC RxD RSSI:  %d\t\tRSRP (dBm):  %d\r\nTx Power:      %s\t\tTAC:         %s\r\nRSRQ (dB):     %f\t\tCell ID:     %s\r\nSINR (dB):      %d\r\n\r\n\r\nOK\r\n"

var GStatusDef = ATCommandDef[GStatus]{
	Cmd:     "AT!GSTATUS?\r",
	Pattern: gstatusPattern,
	Parse: func(r string) (GStatus, error) {
		var out GStatus
		n, _ := fmt.Sscanf(r, gstatusPattern, &out.Time, &out.Temp, &out.Creset,
			&out.Mode) //, //&out.State, &out.PSState, &out.LTEBand)

		if condition := n == 0; condition {
			out.Time = 0
			out.Temp = 0
			out.Creset = 0
			out.Mode = "Null"
			// out.State = "Null"
			// out.PSState = "Null"
			// out.LTEBand = "Null"
		}

		return out, nil
	},
}

// GStatus representa el estado del módem según el comando AT!GSTATUS.
type GPSStatus struct {
	Status int // ONLINE, OFFLINE, etc.

}

// -----------------------AT!GPSSTATUS?\r\r\nCurrent time: %*d %*d %*d %*d %*d:%*d:%*d\r\n\r\n%*d %*d %*d %*d %*d:%*d:%*d Last Fix Status    = NONE
const gpsstatusPattern = "AT!GPSSTATUS?\r\r\nCurrent time: %*s %d" // %*d %*d %*d:%*d:%*d\r\n\r\n%*d %*d %*d %*d %*d:%*d:%*d Last Fix Status    = %s"

var GPSStatusDef = ATCommandDef[GPSStatus]{
	Cmd:     "AT!GPSSTATUS?\r",
	Pattern: gpsstatusPattern,
	Parse: func(r string) (GPSStatus, error) {
		var out GPSStatus
		n, _ := fmt.Sscanf(r, gpsstatusPattern, &out.Status)

		if condition := n == 0; condition {
			out.Status = -1
		}

		return out, nil
	},
}

// ---------------------------------------------------------------------------------------
// Row representa una fila de la tabla de resultados
// Contiene una etiqueta, un valor y un indicador de éxito (OK)
// Esta estructura se utiliza para almacenar los resultados de los comandos AT
// y se convierte en una fila de la tabla que se muestra al usuario
type Row struct {
	Label string
	Value string
	OK    bool
}

// ATExec es una interfaz que define un comando AT y su ejecución
// Cada comando AT debe implementar esta interfaz para ser ejecutado
// y devolver un resultado en forma de fila (Row)
type ATExec interface {
	Cmd() string
	Run(raw string) ([]Row, error)
}

// AttrStep es un paso de ejecución de comando AT con atributos específicos
// T es el tipo de dato que se espera extraer de la respuesta del comando AT
// Por ejemplo, puede ser un string, int, struct, etc.
// Este patrón se usa para escanear la respuesta del comando AT y extraer las variables necesarias.
type AttrStep[T any] struct {
	Def   *ATCommandDef[T]
	Label string
	Value func(T) string
	OK    func(T) bool
}

func (s AttrStep[T]) Cmd() string { return s.Def.Cmd }

func (s AttrStep[T]) Run(raw string) ([]Row, error) {
	v, err := s.Def.Parse(raw)
	if err != nil {
		return nil, err
	}
	r := singleRow{s.Label, s.Value(v), s.OK(v)}
	return []Row{r.toRow()}, nil
}

// convierte singleRow a Row
type singleRow struct {
	label string
	value string
	ok    bool
}

// toRow convierte singleRow a Row
// Esta función es necesaria para que singleRow implemente la interfaz Row
func (r singleRow) toRow() Row { return Row{r.label, r.value, r.ok} }

// Lista de comandos AT a ejecutar
// Cada comando es un paso con su etiqueta, valor y condición de éxito
// Se ejecutan en el orden en que están definidos
// Se pueden agregar más pasos según sea necesario
var SeqCmd = []ATExec{

	AttrStep[Getversion]{
		Def:   &GetVersionDef,
		Label: "Versión de firmware (esperado SWI9X30C_02.38.00.00)",
		Value: func(v Getversion) string { return v.Version },
		OK:    func(v Getversion) bool { return v.Version == "SWI9X30C_02.38.00.00" },
	},
	AttrStep[PcVolt]{
		Def:   &PcVoltDef,
		Label: "Estado de voltaje",
		Value: func(v PcVolt) string { return v.State },
		OK:    func(v PcVolt) bool { return v.State == "Normal" },
	},
	AttrStep[PcVolt]{
		Def:   &PcVoltDef,
		Label: "Voltaje de alimentación (3.3-4.1 V)",
		Value: func(v PcVolt) string { return fmt.Sprintf("%.2f V", float64(v.MilliV)/1000) },
		OK:    func(v PcVolt) bool { f := float64(v.MilliV) / 1000; return f >= 3.3 && f <= 4.1 },
	},
	AttrStep[PcTemp]{
		Def:   &PcTempDef,
		Label: "Temperatura del módem (-20 a 85 °C)",
		Value: func(v PcTemp) string { return fmt.Sprintf("%d °C", int(v.Temp)) },
		OK:    func(v PcTemp) bool { return v.Temp >= -20 && v.Temp <= 85 },
	},
	AttrStep[CFun]{
		Def:   &CFunDef,
		Label: "Modo de operatividad general del módem (AT+CFUN)",
		Value: func(v CFun) string { return fmt.Sprintf("%d", v.Mode) },
		OK:    func(v CFun) bool { return v.Mode != 0 },
	},
	AttrStep[CGReg]{
		Def:   &CGRegDef,
		Label: "Registrado en la red celular (voz) (AT+CREG)",
		Value: func(v CGReg) string { return fmt.Sprintf("%d", v.Stat) },
		OK:    func(v CGReg) bool { return v.Stat == 1 },
	},
	AttrStep[CReg]{
		Def:   &CRegDef,
		Label: "Registrado en la red celular (datos) (AT+CGREG)",
		Value: func(v CReg) string { return fmt.Sprintf("%d", v.Stat) },
		OK:    func(v CReg) bool { return v.Stat == 1 },
	},

	AttrStep[GetBand]{
		Def:   &GetBandDef,
		Label: "Banda activa (AT!GETBAND)",
		Value: func(v GetBand) string { return v.Band },
		OK:    func(v GetBand) bool { return !strings.Contains(v.Band, "No") },
	},
	AttrStep[CGDCont]{
		Def:   &CGDContDef,
		Label: "APN (AT+CGDCONT)",
		Value: func(v CGDCont) string {
			if strings.Contains(v.APN, "NEBULAENGINEERING.tigo.com") {
				return v.APN
			}
			return "No APN"
		},
		OK: func(v CGDCont) bool { return strings.Contains(v.APN, "NEBULAENGINEERING.tigo.com") },
	},
	AttrStep[GPSAtInfo]{
		Def:   &GPSAtInfoDef,
		Label: "Número de satélites visibles (> 6)",
		Value: func(v GPSAtInfo) string { return fmt.Sprintf("%d", v.InView) },
		OK:    func(v GPSAtInfo) bool { return v.InView > 6 },
	},
	AttrStep[GPSAtInfo]{
		Def:   &GPSAtInfoDef,
		Label: "SNR promedio de satélites visibles (> 20 dB)",
		Value: func(v GPSAtInfo) string { return fmt.Sprintf("%.2f dB", v.AvgSNR) },
		OK:    func(v GPSAtInfo) bool { return v.AvgSNR > 20 },
	},
	AttrStep[Want]{
		Def:   &WantDef,
		Label: "Alimentación de la antena GPS (AT+WANT)",
		Value: func(v Want) string { return fmt.Sprintf("%d", v.value) },
		OK:    func(v Want) bool { return v.value == 1 },
	},

	// …aquí agregas más atributos de otros structs, en el orden que quieras…
}

// fecha-hora y estado  :contentReference[oaicite:3]{index=3}
// const gpslocPattern = "AT!GPSLOC?\r\r\nLat: %f\r\nLon: %f\r\nTime: %s"            // posición última  :contentReference[oaicite:5]{index=5}

// Comandos que sólo devuelven “OK”
// AT+WANT=1, AT!ERR=0, AT!GPSEND=0,255, AT!GPSFIX=1,60,10 → basta con verificar “OK”.

// const pcInfoPattern = "AT!PCINFO?\r\r\nState: %s\r\nLPM voters - Temp:%d, Volt:%d, User:%d, W_DISABLE: %d, IMSWITCH: %d, BIOS:%d, LWM2M:%d, OMADM:%d, FOTA:%d\r\nLPM persistence - %s"
