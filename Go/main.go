package main

import (
	"fmt"
)

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

	fmt.Println("\n---------- FUNCION RETORNA MULTIPLES VALORES ------------------")

	max, min := maxMin(8, 14)

	fmt.Printf("El valor mayor es: %d y el menor es: %d\n", max, min)
	fmt.Println("\n---------- FUNCION RETORNA MULTIPLES VALORES NO USAR UN VALOR  ------------------")

	m, _ := maxMinValue(6, 10)

	fmt.Printf("El valor mayor es: %d\n", m)

	// Literales de Funciones

	fmt.Println("\n =============== Literales de una funcion ================== ")

	var operator func(int, int) int

	operator = res

	fmt.Println("Resultado de la resta: ", operator(10, 5))

	operator = multiply

	fmt.Println("Resultado de la resta: ", operator(10, 5))

	fmt.Println("\n =============== Funciones con numero de parametro indefinidos ================== ")

	resultSumaMultiParams := sumaInfiniteParams(1, 4, 7)

	fmt.Println("El valor de la suma con multiple parametros es de: ", resultSumaMultiParams)

	fmt.Println("\n =============== Operador Difusor ================== ")

	numberSlice := []int{2, 5, 6, 3, 8}
	numberSlice2 := []int{6, 7, 89}

	resultSumaSlice := sumaInfiniteParams(numberSlice...)

	fmt.Println("El valor de la suma con multiple paramentro pasando un slice: ", resultSumaSlice)

	numberSlice = append(numberSlice, numberSlice2...) // para poder agregar un slice a otro slice usamos difusor (...)

	fmt.Println("Resultado de agregar un slice sobre otro: ", numberSlice)
}

// Crear funciones base
func sum() int {
	return 5 + 10
}

// funcion con parametros
func res(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

// funcion que retorna multiples valores
func maxMin(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

// funcion que retorna multiples valores con nombre definido
func maxMinValue(a, b int) (max int, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}

func sumaInfiniteParams(numbers ...int) int {
	result := 0

	for _, value := range numbers {
		result += value
	}

	return result
}
