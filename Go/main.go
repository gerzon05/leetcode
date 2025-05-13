package main

import "fmt"

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

var (
	age          int = 19
	name         string
	seconds      float32 = 30.45
	hours        int
	yourAge      int
	valueProgram rune
	close        rune
)

// Constantes
const (
	pi       = 3.1416
	Typefont = "Times New Roman"
	Negrita  = true
)

func main() {

	fmt.Println("Hola Mundo desde Go")

	year := 2025 // recomendable usar mientras no se necesite que sea global

	fmt.Println("La edad es de: ", age)
	fmt.Println("Estamos en el año: ", year)

	// varibales sin asignar valor:
	// int = 0, string = "", bool = false
	name = "Gerzon"

	fmt.Println("Te llamas: ", name)

	// cambiar el tipo de la variblable
	hours = int(30)

	fmt.Println("Segundos: ", seconds)
	fmt.Println("Horas: ", hours)

	// Mostrar valores constantes

	fmt.Println("valor de Pi: ", pi)
	fmt.Println("Nombre de la Fuente: ", Typefont)
	fmt.Println("La fuente tiene Negrita: ", Negrita)

	// Operadores Numericos de Comparacion (==, >=, <=, >, <, !=)

	equall := age == 18
	ofAge := age >= 18

	fmt.Println("La edad es igual a 18: ", equall)
	fmt.Println("La edad es mayor o igual a 18: ", ofAge)

	// Operadores logicos (&&, ||, !)

	// pedir valores por consola
	fmt.Print("Cual es tu edad: ")
	fmt.Scan(&yourAge)
	// fmt.Scanf("%d", &yourAge) especificar el tipo que debe recibir
	fmt.Println("Tu edad es de ", yourAge, "años")

	fmt.Printf("!hola %s¡ tu edad es de %d años.\n", name, yourAge)

	// Bloques Condicionales

	if yourAge < 18 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Eres mayor de edad")
	}

	// Si la variable solo se va a uasar para la condicion es buena practica declararla dentro del if

	if evenNumber := 4; evenNumber%2 == 0 {
		fmt.Printf("El numero %d es par\n", evenNumber)
	} else {
		fmt.Printf("El numero %d es impar\n", evenNumber)
	}

	// Switch - Case

	switch year {
	case 18:
		fmt.Println("Is big.. In switch")
	default:
		fmt.Println("Is default in switch")
	}

	resultSum := sum()

	fmt.Println("El resultado de suma es: ", resultSum)

	resultRes := res(10, 5)

	fmt.Println("El resultado de la resta: ", resultRes)

	// Bucle For
	// Primera forma
	for {
		fmt.Print("Ingresa \"s\" para salir de programa o \"n\" para continuar: ")
		fmt.Scanf("%c\n", &valueProgram)
		if valueProgram == 'n' || valueProgram == 'N' {
			continue
		} else if valueProgram == 's' || valueProgram == 'S' {
			break
		}
		fmt.Println("El valor ingresado es Incorrecto")
	}
	fmt.Println("!Adios¡")

	// Usar el for como el while de otros lenguaje de programacion
	for close != 's' && close != 'S' {
		fmt.Print("Salir? (s/n): ")
		fmt.Scanf("%c\n", &close)
	}
	fmt.Println("!Gracias¡")

	// for comun de todos los lenguajes

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// Ocultacion de Variables

	a := 0
	b := 0

	if true {
		a := 0
		b = 1
		a++
		b++
	}

	fmt.Printf("Valor de a = %d, valor de b = %d\n", a, b)

}

// Crear funciones base
func sum() int {
	return 5 + 10
}

// funcion con parametros
func res(a int, b int) int {
	return a - b
}
