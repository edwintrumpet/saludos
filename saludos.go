package saludos

import "fmt"

func Hola(name string) string {
	return fmt.Sprintf("¡hola %s! ¿cómo estás?", name)
}
