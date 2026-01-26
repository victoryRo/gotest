package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expected: %q but got: %q", expected, repeated)
	}
}

// BenchmarkRepeat prueba la eficiencia del codigo
// go test -bench=.
// Nota: De forma predeterminada,
// los Benchmark se ejecutan secuencialmente.
// Solo se cronometra el cuerpo del bucle;
// excluye automáticamente el código de configuración
// y limpieza del cronometraje de la prueba de rendimiento.
// Un Benchmark típico se estructura así:
func BenchmarkRepeat(b *testing.B) {
	//... setup ...
	for b.Loop() {
		Repeat("a", 10) //... code to measure ...
	}
	//... cleanup ...
}

// go test -bench=. -benchmem
// La bandera   -benchmem informa sobre las asignaciones de memoria
// B/op:        el número de bytes asignados por iteración
// allocs/op:   el número de asignaciones de memoria por iteración

func ExampleRepeat() {
	res := Repeat("c", 10)
	fmt.Println(res)
	// Output: cccccccccc
}
