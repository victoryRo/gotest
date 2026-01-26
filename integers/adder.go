// Package integers for test
package integers

// go doc integers
// podemos ver la descripcion del package
// go doc integers Add
// podemos ver la documentacion de la funcion

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

// pkgsite: un servidor de documentación
// pkgsite El programa extrae y genera documentación para proyectos Go.
// Ejemplo de uso:

// $ go install golang.org/x/pkgsite/cmd/pkgsite@latest
// $ cd myproject
// $ pkgsite -open .
