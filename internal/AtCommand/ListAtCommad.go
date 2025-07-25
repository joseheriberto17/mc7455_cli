// file ListAtCommad.go

package AtCommand

import (
	"fmt"
	"strings"
)

// Getversion representa la versión de firmware del módem.
type Getversion struct{ Version string }

func (g Getversion) Label() string { return "Versión de firmware (esperado SWI9X30C_02.38.00.00)" }
func (g Getversion) Value() string { return g.Version }
func (g Getversion) OK() bool      { return g.Version == "SWI9X30C_02.38.00.00" }

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

func (v PcVolt) Label() string { return "Voltaje de alimentación (3.3-4.1 V)" }
func (v PcVolt) Value() string { return fmt.Sprintf("%.2f V", float64(v.MilliV)/1000) }
func (v PcVolt) OK() bool      { f := float64(v.MilliV) / 1000; return f >= 3.3 && f <= 4.1 }

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

func (t PcTemp) Label() string { return "Temperatura del módem (-20 a 85 °C)" }
func (t PcTemp) Value() string { return fmt.Sprintf("%d °C", int(t.Temp)) }
func (t PcTemp) OK() bool      { return t.Temp >= -20 && t.Temp <= 85 }

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

// ContErr representa el contador de errores del módem.
type ContErr struct {
	value int
}

func (e ContErr) Label() string { return "Contador de errores módem (!ERR)" }
func (e ContErr) Value() string { return fmt.Sprintf("%d", e.value) }
func (e ContErr) OK() bool      { return e.value == 0 }

const errPattern = "AT!ERR\r\r\n%02d   %02x %s %05d"

var ContErrDef = ATCommandDef[ContErr]{
	Cmd:     "AT!ERR?\r",
	Pattern: errPattern,
	Parse: func(resp string) (ContErr, error) {
		var out ContErr
		n, _ := fmt.Sscanf(resp, errPattern,
			&out.value)

		if condition := n == 0; condition {
			out.value = 0
		}
		return out, nil
	},
}

// CFun representa el modo de operatividad general del módem.
type CFun struct{ Mode int }

func (c CFun) Label() string { return "Operatividad general del modem (AT+CFUN)" }
func (c CFun) Value() string { return fmt.Sprintf("%d", c.Mode) }
func (c CFun) OK() bool      { return c.Mode != 0 }

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

func (w Want) Label() string { return "Alimentación antena GPS (AT+WANT)" }
func (w Want) Value() string { return fmt.Sprintf("%d", w.value) }
func (w Want) OK() bool      { return w.value == 1 }

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

func (c CReg) Label() string { return "Registrado CS (voz) (AT+CREG)" }
func (c CReg) Value() string { return fmt.Sprintf("%d", c.Stat) }
func (c CReg) OK() bool      { return c.Stat == 1 }

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

func (c CGReg) Label() string { return "Registrado PS (datos) (AT+CGREG)" }
func (c CGReg) Value() string { return fmt.Sprintf("%d", c.Stat) }
func (c CGReg) OK() bool      { return c.Stat == 1 }

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

func (g GetBand) Label() string { return "Banda activa (AT!GETBAND)" }
func (g GetBand) Value() string { return g.Band }
func (g GetBand) OK() bool      { return !strings.Contains(g.Band, "No") }

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
	PDPType int
	APN     string
	// IP      string
	// P1, P2  int
	// P3, P4  int
}

func (c CGDCont) Label() string { return "PDP context / APN (AT+CGDCONT)" }
func (c CGDCont) Value() string {
	if strings.Contains(c.APN, "NEBULAENGINEERING.tigo.com") {
		return c.APN

	}
	return "No APN"
}
func (c CGDCont) OK() bool { return strings.Contains(c.APN, "NEBULAENGINEERING.tigo.com") }

const cgdcontPattern = "AT+CGDCONT?\r\r\n+CGDCONT: %d,\"IP\",\"%s" //\",\"%*s\",%*d,%*d,%*d,%*d"

var CGDContDef = ATCommandDef[CGDCont]{
	Cmd:     "AT+CGDCONT?\r",
	Pattern: cgdcontPattern,
	Parse: func(r string) (CGDCont, error) {
		var out CGDCont
		n, _ := fmt.Sscanf(r, cgdcontPattern,
			&out.PDPType, &out.APN)
		// ,&out.IP,&out.P1, &out.P2, &out.P3, &out.P4)

		if condition := n == 0; condition {
			out.PDPType = 0
			out.APN = "Null"
			// out.IP = "Null"
			// out.P1 = 0
			// out.P2 = 0
			// out.P3 = 0
			// out.P4 = 0
		}

		return out, nil
	},
}

// CGDCont representa un contexto PDP (Packet Data Protocol) del módem.
type GPSAtInfo struct {
	Nsat int
}

func (c GPSAtInfo) Label() string { return "Numero de satelites (> 6)" }
func (c GPSAtInfo) Value() string { return fmt.Sprintf("%d", c.Nsat) }
func (c GPSAtInfo) OK() bool      { return c.Nsat > 6 }

const gpssatinfoPattern = "AT!GPSSATINFO?\r\r\nSatellites in view:  %d"

var GPSAtInfoDef = ATCommandDef[GPSAtInfo]{
	Cmd:     "AT!GPSSATINFO?\r",
	Pattern: gpssatinfoPattern,
	Parse: func(r string) (GPSAtInfo, error) {
		var out GPSAtInfo
		n, _ := fmt.Sscanf(r, gpssatinfoPattern, &out.Nsat)
		if condition := n == 0; condition {
			out.Nsat = 0
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

func (g GStatus) Label() string { return "Estado del módem en la RED (AT!GSTATUS)" }
func (g GStatus) Value() string { return g.Mode }
func (g GStatus) OK() bool      { return g.Mode == "ONLINE" }

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

func (g GPSStatus) Label() string { return "Estado del GPS (AT!GPSSTATUS)" }
func (g GPSStatus) Value() string { return fmt.Sprint(g.Status) }
func (g GPSStatus) OK() bool      { return g.Status != -1 }

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

type raws interface {
	Label() string
	Value() string
	OK() bool
}

// ---------------------------------------------------------------------------------------

// Ejecutor genérico de un comando AT:
type ATExec interface {
	Cmd() string                  // qué enviar al módem
	Run(raw string) (raws, error) // cómo convertir la respuesta
}

// envuelve un *ATCommandDef[T] y lo hace cumplir ATExec
type ExecWrap[T raws] struct {
	def *ATCommandDef[T]
}

func (w ExecWrap[T]) Cmd() string { return w.def.Cmd }

func (w ExecWrap[T]) Run(raw string) (raws, error) {
	v, err := w.def.Parse(raw) // v es tipo T
	if err != nil {
		return nil, err
	}
	return any(v).(raws), nil // garantiza que T implementa raws
}

var SeqCmd = []ATExec{
	ExecWrap[Getversion]{&GetVersionDef},
	ExecWrap[PcVolt]{&PcVoltDef},
	ExecWrap[PcTemp]{&PcTempDef},
	ExecWrap[CFun]{&CFunDef},
	ExecWrap[GStatus]{&GStatusDef},
	// ExecWrap[ContErr]{&ContErrDef},
	ExecWrap[CGReg]{&CGRegDef},
	ExecWrap[CReg]{&CRegDef},
	ExecWrap[GetBand]{&GetBandDef},

	ExecWrap[Want]{&WantDef},
	ExecWrap[CGDCont]{&CGDContDef},
	// ExecWrap[GPSStatus]{&GPSStatusDef},
	ExecWrap[GPSAtInfo]{&GPSAtInfoDef},
}

// fecha-hora y estado  :contentReference[oaicite:3]{index=3}
// const gpslocPattern = "AT!GPSLOC?\r\r\nLat: %f\r\nLon: %f\r\nTime: %s"            // posición última  :contentReference[oaicite:5]{index=5}

// Comandos que sólo devuelven “OK”
// AT+WANT=1, AT!ERR=0, AT!GPSEND=0,255, AT!GPSFIX=1,60,10 → basta con verificar “OK”.

// const pcInfoPattern = "AT!PCINFO?\r\r\nState: %s\r\nLPM voters - Temp:%d, Volt:%d, User:%d, W_DISABLE: %d, IMSWITCH: %d, BIOS:%d, LWM2M:%d, OMADM:%d, FOTA:%d\r\nLPM persistence - %s"
