package internal

import (
	"fmt"
)

func RunPointer() {
	fmt.Println("Aprenderemos sobre los punteros")

	i := 10 // guardas el valor de la variables
	a := &i // guardar la direccion de memoria de la vatiable i
	p := *a // guardar el valor a donde apunta esa direccion de memoria a

	fmt.Println("Direccion de memoria &: ", a)               //
	fmt.Println("Apuntado a la direccion de memoria *: ", p) // 10

	*a = 20 // cambiar el valor de la direccion de memoria que se esta referenciando

	fmt.Println("Apuntado a la direccion de memoria desde a *: ", *a)
	fmt.Println("Apuntado a la direccion de memoria desde p: ", p)
	fmt.Println("Apuntado a la direccion de memoria desde i: ", i)

	// Paso por Valor - Paso por referencia
	fmt.Println("Paso por valor")

	valueIncrement := 3

	fmt.Println("Valor a Incrementar: ", valueIncrement)

	incrementValue(valueIncrement) // valor por parametro crea una copia dentro de la funcion

	fmt.Println("Valor a Incrementar: ", valueIncrement)

	fmt.Println("Paso por Referencia")

	valueReference := 3

	fmt.Println("Valor a Referenciar: ", valueReference)

	incrementValueReference(&valueReference) // Pasamos la direccion de memoria

	fmt.Println("Valor a Refenrenciar: ", valueReference)

}

func incrementValue(a int) {
	a++
}
func incrementValueReference(a *int) { // recibe la posicion de memoria sin hacer una copia del valor
	*a++ // apunta a la direccion de memoria (valor) y le aumenta 1
}
