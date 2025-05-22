package internal

import "fmt"

func RunDataStructure() {
	fmt.Println("Estructuras de datos lineal")

	fmt.Println("1. Vectores (Arrays)")

	arr := [5]int{1, 2, 3, 4, 5} // se espcifica la cantidad de datos que almacenara
	//arr := [...]int{1,2,3} // con [...] dejamos qe el compilador infiera la cantidad de datos almacenados.

	fmt.Println("Mostrar el vector: ", arr)

	fmt.Println("Leer un valor especifico del vector: ", arr[2])

	// Matrices

	var checkerBoard [8][8]uint8

	fmt.Println("Tablero de ajedres: ", checkerBoard)

	// Porciones (Slices)
	// se decara igual a un vector pero no se espcifica la cantidad de datos que almacenara

	fmt.Printf("\n2. Slice\n")

	var slice []int // En este caso apunta a nil

	fmt.Println("Mostrar el Slice: ", slice)

	number := []int{1, 2, 3, 4}

	fmt.Println("Mostar Slice con valores: ", number)

	// Agregar elemento a un slice

	number = append(number, 5)

	fmt.Println("Mostar Slice agregado un valor nuevo: ", number)

	fmt.Printf("\n2. Medir dimesiones\n")
	// Dimenciones con Len y Cap

	fmt.Printf("Longitud: %v. Capacidad: %v.\n", len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4)

	fmt.Printf("Longitud: %v. Capacidad: %v.\n", len(slice), cap(slice))

	slice = append(slice, 5)

	fmt.Printf("Longitud: %v. Capacidad: %v.\n", len(slice), cap(slice))

}
