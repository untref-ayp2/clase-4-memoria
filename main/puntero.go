package main

import (
	"fmt"
)

func main() {
	a := 10
	showDuplicate(a)
	fmt.Println("main.a:", a)
	duplicate(&a) // Le envio la direccion
	fmt.Println("main.a:", a)
}

func showDuplicate(b int) {
	// a no existe
	// b vale a
	// Podemos decir que la funcion es inmutable
	b *= 2
	fmt.Println("showDuplicate.b:", b)
}

func duplicate(c *int) {
	// Esta funcion es mutable
	*c *= 2
	fmt.Println("duplicate.c:", c)
}
