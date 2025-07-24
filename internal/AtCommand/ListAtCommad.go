// file ListAtCommad.go

package AtCommand

import "fmt"

type PcVolt struct {
	State  string
	MilliV int
	ADC    int
}

type PcTemp struct {
	State string
	Temp  float32
}

type CFun struct{ Mode int }

type Want struct{ Value int }

type CReg struct {
	N    int
	Stat int
}

type CGReg struct {
	N    int
	Stat int
}

// type GStatus struct {
// 	N    int
// 	Stat int
// }

type GetBand struct{ Band string }
type Getversion struct{ Version string }

type CGDCont struct {
	Cid     int
	PDPType string
	APN     string
	IP      string
	P1, P2  int
	P3, P4  int
}

const getversionPattern = "AT+GMR\r\r\n%s "

// Mediciones de alimentación y temperatura
const pcvoltPattern = "AT!PCVOLT?\r\r\nVolt state: %s\r\nPower supply voltage: %d mV (ADC: %d)"
const pctempPattern = "AT!PCTEMP?\r\r\nTemp state: %s\r\nTemperature: %04.2f C"

// Estado / modo de la radio
const cfunPattern = "AT+CFUN?\r\r\n+CFUN: %d"
const wantPattern = "AT+WANT?\r\r\n+WANT: %d"
const cregPattern = "AT+CREG?\r\r\n+CREG: %d,%d"
const cgregPattern = "AT+CGREG?\r\r\n+CGREG: %d,%d"

// const gstatusPattern = "AT!GSTATUS?\r\r\n!GSTATUS:\r\nCurrent Time: %d\t\tTemperature: %d" // cabecera, luego vienen más campos  :contentReference[oaicite:1]{index=1}
const getbandPattern = "AT!GETBAND?\r\r\n!GETBAND: %s" // descripción de banda activa  :contentReference[oaicite:2]{index=2}

// PDP context / APN
const cgdcontPattern = "AT+CGDCONT?\r\r\n+CGDCONT: %d,\"%[^\"]\",\"%[^\"]\",\"%[^\"]\",%d,%d,%d,%d"

// GPS – estado, sesión y satélites
// const gpsstatusPattern = "AT!GPSSTATUS?\r\r\n%d %d %d %d %s Last Fix Status = %s" // fecha-hora y estado  :contentReference[oaicite:3]{index=3}
// const gpssatinfoPattern = "AT!GPSSATINFO?\r\r\n* SV: %d  ELEV:%d  AZI:%d  SNR:%d" // línea típica por satélite  :contentReference[oaicite:4]{index=4}
// const gpslocPattern = "AT!GPSLOC?\r\r\nLat: %f\r\nLon: %f\r\nTime: %s"            // posición última  :contentReference[oaicite:5]{index=5}

// Diagnóstico
// const errPattern = "AT!ERR\r\r\n%02d   %02x %s %05d" // una línea del listado de errores  :contentReference[oaicite:6]{index=6}

// Comandos que sólo devuelven “OK”
// AT+WANT=1, AT!ERR=0, AT!GPSEND=0,255, AT!GPSFIX=1,60,10 → basta con verificar “OK”.

// const pcInfoPattern = "AT!PCINFO?\r\r\nState: %s\r\nLPM voters - Temp:%d, Volt:%d, User:%d, W_DISABLE: %d, IMSWITCH: %d, BIOS:%d, LWM2M:%d, OMADM:%d, FOTA:%d\r\nLPM persistence - %s"

var PcVoltDef = ATCommandDef[PcVolt]{
	Cmd:     "AT!PCVOLT?\r",
	Pattern: pcvoltPattern,
	Parse: func(resp string) (PcVolt, int, error) {
		var out PcVolt
		lenght, err := fmt.Sscanf(resp, pcvoltPattern,
			&out.State, &out.MilliV, &out.ADC)
		return out, lenght, err
	},
}

var PcTempDef = ATCommandDef[PcTemp]{
	Cmd:     "AT!PCTEMP?\r",
	Pattern: pctempPattern,
	Parse: func(resp string) (PcTemp, int, error) {
		var out PcTemp
		n, err := fmt.Sscanf(resp, pctempPattern,
			&out.State, &out.Temp)
		return out, n, err
	},
}

var CFunDef = ATCommandDef[CFun]{
	Cmd:     "AT+CFUN?\r",
	Pattern: cfunPattern,
	Parse: func(r string) (CFun, int, error) {
		var out CFun
		n, err := fmt.Sscanf(r, cfunPattern, &out.Mode)
		return out, n, err
	},
}

var WantDef = ATCommandDef[Want]{
	Cmd:     "AT+WANT?\r",
	Pattern: wantPattern,
	Parse: func(r string) (Want, int, error) {
		var out Want
		n, err := fmt.Sscanf(r, wantPattern, &out.Value)
		return out, n, err
	},
}

var CRegDef = ATCommandDef[CReg]{
	Cmd:     "AT+CREG?\r",
	Pattern: cregPattern,
	Parse: func(r string) (CReg, int, error) {
		var out CReg
		n, err := fmt.Sscanf(r, cregPattern, &out.N, &out.Stat)
		return out, n, err
	},
}

var CGRegDef = ATCommandDef[CGReg]{
	Cmd:     "AT+CGREG?\r",
	Pattern: cgregPattern,
	Parse: func(r string) (CGReg, int, error) {
		var out CGReg
		n, err := fmt.Sscanf(r, cgregPattern, &out.N, &out.Stat)
		return out, n, err
	},
}

var GetBandDef = ATCommandDef[GetBand]{
	Cmd:     "AT!GETBAND?\r",
	Pattern: getbandPattern,
	Parse: func(r string) (GetBand, int, error) {
		var out GetBand
		n, err := fmt.Sscanf(r, getbandPattern, &out.Band)
		return out, n, err
	},
}

var CGDContDef = ATCommandDef[CGDCont]{
	Cmd:     "AT+CGDCONT?\r",
	Pattern: cgdcontPattern,
	Parse: func(r string) (CGDCont, int, error) {
		var out CGDCont
		n, err := fmt.Sscanf(r, cgdcontPattern,
			&out.Cid, &out.PDPType, &out.APN, &out.IP,
			&out.P1, &out.P2, &out.P3, &out.P4)
		return out, n, err
	},
}

var GetVersionDef = ATCommandDef[Getversion]{
	Cmd:     "AT+GMR\r",
	Pattern: getversionPattern,
	Parse: func(r string) (Getversion, int, error) {
		var out Getversion
		n, err := fmt.Sscanf(r, getversionPattern,
			&out.Version)
		return out, n, err
	},
}
