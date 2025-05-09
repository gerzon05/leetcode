package main

import "fmt"

func main() {

	fmt.Println("Hola Mundo desde Go")

	// Tipos de datos en Go:

	// - Números enteros con signo (int):
	//     int8, int16, int32, int64

	// - Números enteros sin signo (uint):
	//     uint8, uint16, uint32, uint64
	//     *El tamaño de 'uint' sin sufijo depende de la plataforma (32 o 64 bits)*

	// - Cadenas de texto:
	//     string (codificación UTF-8)

	// - Caracteres Unicode:
	//     rune (alias de int32)

	// - Números de punto flotante:
	//     float32, float64

	// - Números complejos:
	//     complex64, complex128 (contienen parte real e imaginaria) ej: 2.23 + 3i

	// Definir Variables: (var,const) <nomber> <tipo> = valor || <nombre> := valor

	var age int = 19
	year := 2025 // recomendable usar mientras no se necesite que sea global

	fmt.Println("La edad es de: ", age)
	fmt.Println("Estamos en el año: ", year)

	// varibales sin asignar valor:
	// int = 0, string = "", bool = false

	var name string

	name = "Gerzon"

	fmt.Println("Te llamas: ", name)

	// cambiar el tipo de la variblable

	var seconds float32 = 30.45
	var hours int
	hours = int(30)

	fmt.Println("Segundos: ", seconds)
	fmt.Println("Horas: ", hours)

	// Constantes
	const (
		pi       = 3.1416
		Typefont = "Times New Roman"
		Negrita  = true
	)

	fmt.Println("valor de Pi: ", pi)
	fmt.Println("Nombre de la Fuente: ", Typefont)
	fmt.Println("La fuente tiene Negrita: ", Negrita)
}
