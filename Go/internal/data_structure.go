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
	// tener en cuenta que funcionan con referencia.

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

	fmt.Printf("Longitud: %v - Capacidad: %v\n", len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4)

	fmt.Printf("Longitud: %v - Capacidad: %v\n", len(slice), cap(slice))

	slice = append(slice, 5)

	fmt.Printf("Longitud: %v - Capacidad: %v\n", len(slice), cap(slice))

	// Controlar el tamaño de memoria con Make
	// lo recomendables es crear slice con longitud de 0 pero con la capacidad necesaria
	fmt.Printf("\n3. Make\n")

	sliceWithMake := make([]int, 20, 30)
	highCapacity := make([]int, 0, 2048)

	fmt.Printf("Slice: %v - Longitud: %v - Capacidad: %v\n", sliceWithMake, len(sliceWithMake), cap(sliceWithMake))
	fmt.Printf("Slice: %v - Longitud: %v - Capacidad: %v\n", highCapacity, len(highCapacity), cap(highCapacity))

	sliceWithMake = append(sliceWithMake, 10)
	highCapacity = append(highCapacity, 19)

	fmt.Printf("Slice: %v - Longitud: %v - Capacidad: %v\n", sliceWithMake, len(sliceWithMake), cap(sliceWithMake))
	fmt.Printf("Slice: %v - Longitud: %v - Capacidad: %v\n", highCapacity, len(highCapacity), cap(highCapacity))

	// Copia de porciones (slice) con (copy)
	fmt.Printf("\n4. Copy\n")

	original := []int{8, 10, 20}
	sliceCopy := make([]int, len(original))
	// sliceCopy := original // usamos la misma direccion de memoria

	fmt.Println("direccion de memoria original 0: ", original)
	fmt.Println("direccion de memoria sliceCopy 0: ", sliceCopy)

	// original = append(original, 40)

	newCopy := copy(sliceCopy, original)

	original[0] = 100

	fmt.Println("ver Slice original: ", original)
	fmt.Println("ver Slice con Make: ", sliceCopy)

	fmt.Println("Numeros copiados: ", newCopy, " Valores copiados: ", sliceCopy)

}
