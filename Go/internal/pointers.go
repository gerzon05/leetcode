package pointers

import (
	"fmt"
)

func Run() {
	fmt.Println("Aprenderemos sobre los punteros")

	i := 10 // guardas el valor de la variables
	a := &i // guardar la direccion de memoria de la vatiable i
	p := *a // guardar el valor a donde apunta esa direccion de memoria a

	fmt.Println("Direccion de memoria &: ", a)
	fmt.Println("Apuntado a la direccion de memoria *: ", p)

	*a = 20 // cambiar el valor de la direccion de memoria que se esta referenciando

	fmt.Println("Apuntado a la direccion de memoria desde a *: ", *a)
	fmt.Println("Apuntado a la direccion de memoria desde p: ", p)
	fmt.Println("Apuntado a la direccion de memoria desde i: ", i)

}
