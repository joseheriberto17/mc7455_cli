// file AtCommad.go
package AtCommand

import "fmt"

type ATCommandDef[T any] struct {
	Cmd     string                  // Lo que se envía al módem
	Pattern string                  // Plantilla para fmt.Sscanf
	Parse   func(string) (T, error) // Convierte la respuesta en un struct
}

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
