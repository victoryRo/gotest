package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	// este es un subtest dentro del test principal
	// que nos permite crear diferentes pruebas para un mismo test.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Maria", "")
		want := "Hello Maria"

		assertCorrectMsg(t, got, want)
	})

	t.Run("say Hello world if name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"

		assertCorrectMsg(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Maria", "Spanish")
		want := "Hola Maria"

		assertCorrectMsg(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Pedro", "French")
		want := "Bonjour Pedro"

		assertCorrectMsg(t, got, want)
	})

	t.Run("in italian", func(t *testing.T) {
		got := Hello("Isabella", "Italian")
		want := "Ciao Isabella"

		assertCorrectMsg(t, got, want)
	})
}

// esta es una funcion auxiliar para el test
// donde podemos refactorizar para evitar codigo repetido
// es una buena idea aceptar a, testing.TB por que
// es una interfaz que satisfacen ambos *testing.T y *testing.B
func assertCorrectMsg(t testing.TB, got, want string) {
	// con este metodo indicamos que esta funcion es auxiliar
	// tambien permite que el error de la prueba apunte
	// a la prueba misma y no a esta funcion auxiliar
	t.Helper()

	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}
