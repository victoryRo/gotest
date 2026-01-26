// Package hello test
package hello

// ### El proceso TDD y por qué son importantes los pasos ###
//
//   Escriba una prueba fallida y vea cómo falla
//   para que sepamos que hemos escrito una prueba relevante para nuestros requisitos
//   y hayamos visto que produce una descripción fácil de entender de la falla.

//   Escribir la menor cantidad de código posible para que pase
//   y así sepamos que tenemos un software que funciona.

//   Luego refactorizamos, respaldados por la seguridad de nuestras pruebas
//   para garantizar que tengamos un código bien elaborado con el que sea fácil trabajar.

const (
	// languages
	spanish = "Spanish"
	french  = "French"
	italian = "Italian"

	// greetings
	englisHelloPrefix  = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
	italianHelloPrefix = "Ciao "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

// greetingPrefix esta funcion es privada
// para que los componentes de nuesto codigo
// no quede expuesto al mundo exterior
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englisHelloPrefix
	}

	return
}

func LocalMain() {
	// fmt.Println(Hello("Victor"))
}
