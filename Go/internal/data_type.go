package internal

import (
	"fmt"
	"math/rand"
	"strings"
)

func DataTypeRun() {
	fmt.Println("============= Tipos de Datos =============")
	fmt.Println()

	// Crear tipos de datos personalizados

	type email string
	type age uint8 // recibe un numero entero entre 0 - 255

	// declara varibles con tipos personalizados
	var user age = 34
	var userEmail email = "gerzon@gmail.com"

	fmt.Printf("La edad es de: %d, y el correo es: %s \n", user, userEmail)
	fmt.Println()
	fmt.Println("----------------- Tipos a partir de porciones ---------------------")
	fmt.Println()

	type year int
	type births []year

	census := births{1897, 1987, 2000, 2003, 2004, 2005}

	census = append(census, 2010)

	sum := year(0)

	for _, value := range census {
		sum += value
	}

	middle := sum / year(len(census))

	fmt.Println("Promedio de los años de nacimientos: ", middle)
	fmt.Println()
	fmt.Println("----------------- Tipos a partir de Mapas ---------------------")
	fmt.Println()
	type yearmap int
	type namesBirths map[string]yearmap

	famous := namesBirths{
		"Vicent Van Gogh": 1853,
		"Elvis Persley":   1935,
		"Salvador Dali":   1904,
	}

	famous["Rick Astley"] = 1966

	for name, year := range famous {
		fmt.Printf("El famoso %s nacio en el año: %d\n", name, year)
	}

	fmt.Println()
	fmt.Println("----------------- Tipos Funcionales ---------------------")
	fmt.Println()

	counter1 := count()

	fmt.Println("Counter1: ", counter1())
	fmt.Println("Counter1: ", counter1())
	fmt.Println("Counter1: ", counter1())

	counter2 := count()

	fmt.Println("Counter2: ", counter2())
	fmt.Println("Counter2: ", counter2())
	fmt.Println("Counter1: ", counter1())
	fmt.Println("Counter2: ", counter2())

	fmt.Print()

	numberRandom := random(2)

	fmt.Println("numero random maximo 20: ", numberRandom())
	fmt.Println("numero random maximo 20: ", numberRandom())

	fmt.Print()

	for i := 0; i < 6; i++ {
		fmt.Printf("Ejecucion %d: %v\n", i, generatorAll(cero, counter1, numberRandom))
	}

	fmt.Println()
	fmt.Println("----------------- Receptores de Funciones (Metodos) ---------------------")
	fmt.Println()

	var sb strings.Builder

	sb.WriteString("Hola")
	sb.WriteString("!")
	str := sb.String()

	fmt.Println("valor del receptor", str)

	phs := []ph{
		ph(7), ph(1.2), ph(9),
	}

	for _, ph := range phs {
		fmt.Printf("Un ph == %v es %v\n", ph, ph.category())
	}

	var c counter

	c.increment()
	c.increment()
	c.increment()
	c.increment()

	fmt.Println("Valor: ", c)
	c.reset(50)
	fmt.Println("Se reinicio el contador: ", c)
	c.increment()
	c.increment()
	fmt.Println("nuevos valores tras el reinicio: ", c)

}

// ----------------- Tipos Funcionales ---------------------

type Generator func() int

func cero() int {
	return 0
}

func count() Generator {
	countNumber := 0
	return func() int {
		countNumber++
		return countNumber
	}
}

func random(seed int64) Generator {
	number := rand.NewSource(seed)
	return func() int {
		return int(number.Int63())
	}
}

func generatorAll(gens ...Generator) []int {
	nums := make([]int, 0, len(gens))
	for _, value := range gens {
		nums = append(nums, value())
	}
	return nums
}

// Receptores de Funciones (Metodos)

type ph float32

func (p ph) category() string {
	switch {
	case p < 7:
		return "àcido"
	case p > 7:
		return "Basico"
	default:
		return "neutro"
	}
}

type counter int

func (c *counter) increment() {
	*c++
}

func (c *counter) reset(newValue int) {
	*c = counter(newValue)
}
