package internal

import (
	"fmt"
)

// Definir interfaces -> publica: PascalCase ; privada: camelCase
type greeter interface {
	greet() string
}

// Manejo Seguro De Tipo De Dato
type dog struct{}
type cat struct{}
type frog struct{}
type kangaroo struct{}

func RunInterface() {
	fmt.Println("============= INTERFACE =============")
	fmt.Println()

	var sld greeter
	sld = dog{}
	fmt.Println("Saludar desde la interface greeter(dog): ", sld.greet())
	sld = cat{}
	fmt.Println("Saludar desde la interface greeter(Cat): ", sld.greet())

	fmt.Println()
	fmt.Println("----------------- Manejo Seguro De Tipo De Datos ---------------------")
	fmt.Println()

	animals := []interface{}{123, kangaroo{}, "Un helicoptero", dog{}, frog{}}

	for _, animal := range animals {
		fmt.Printf("%#v", animal)
		if d, ok := animal.(dog); ok {
			fmt.Print(": ", d.barks())
		}
		if d, ok := animal.(frog); ok {
			fmt.Print(": ", d.croaks())
		}
		if d, ok := animal.(kangaroo); ok {
			fmt.Print(": ", d.jumps())
		}
		fmt.Println()
	}

	// usando interface con puntero
	var value interface{} = &dog{}

	if d, ok := value.(dog); ok {
		fmt.Println("No mostara nada", d.barks())
	}
	if d, ok := value.(*dog); ok {
		fmt.Println("Muestra el valor: ", d.barks())
	}

	// Validacion mas segura
	for _, animal := range animals {
		switch x := animal.(type) {
		case dog:
			fmt.Println("El animal es un perro: ", x.barks())
		case frog:
			fmt.Println("El animal es una rana: ", x.croaks())
		case kangaroo:
			fmt.Println("El animal es un perro: ", x.jumps())
		default:
			fmt.Println("el valor a continuacion no es valido: ", x)
		}
	}

}

func (d dog) greet() string {
	return "!Guau¡"
}
func (d cat) greet() string {
	return "!Miau¡"
}

// Manejo Seguro De Tipo De Dato
func (d *dog) barks() string {
	return "!Guau¡"
}
func (d *frog) croaks() string {
	return "!Croack¡"
}
func (d *kangaroo) jumps() string {
	return "!Boing¡"
}
