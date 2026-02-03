package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected: %d but got: %d", expected, sum)
	}
}

// Testable Examples in Go
// https://go.dev/blog/examples

// INFO: func Example

// Las funciones de ejemplo se compilan al ejecutar pruebas.
// Dado que el compilador de Go valida estos ejemplos,
// puede estar seguro de que los ejemplos de su documentación
// siempre reflejan el comportamiento actual del código.

func ExampleAdd() {
	res := Add(2, 5)
	fmt.Println(res)
	// Output: 7
}
