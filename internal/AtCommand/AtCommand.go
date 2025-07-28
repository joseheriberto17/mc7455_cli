// file AtCommad.go
package AtCommand

import "fmt"

// ATCommandDef define un comando AT y su patrón de respuesta
// T es el tipo de dato que se espera extraer de la respuesta del comando AT
// Por ejemplo, puede ser un string, int, struct, etc.
// Este patrón se usa para escanear la respuesta del comando AT y extraer lasvariables necesarias.
type ATCommandDef[T any] struct {
	Cmd     string
	Pattern string // Plantilla para fmt.Sscanf
	Parse   func(resp string) (T, error)
}

// Run ejecuta el comando AT y devuelve el resultado
func (def *ATCommandDef[T]) Extract(resp string) (T, error) {
	var out T
	if def.Parse != nil {
		return def.Parse(resp)
	}
	lenght, err := fmt.Sscanf(resp, def.Pattern, &out)

	if condition := lenght < 1; condition {
		return out, fmt.Errorf("respuesta no coincide con el patrón: %s", def.Pattern) // Error si no se pudo extraer
	}

	return out, err
}
