package internal

import (
	"fmt"
)

// Definir una estructura
type cuboide struct {
	width  float64
	height float64
	deep   float64
}

// Incrustado de estructura
type product struct {
	name  string
	price int
}

// go hace una simulacion de herencia que usan otros lenguajes
type phone struct {
	product
	inches  int
	battery int
}

// Opciones funcionales
type student struct {
	name     string
	birth    string
	discount bool
}

type option func(*student)

func StructRun() {

	fmt.Println("============= Data Struct =============")
	fmt.Println()

	// instanciar una variable de tipo estrucutra
	// var c1 cuboide // forma 1
	c1 := cuboide{} // forma 2

	c1.width = 9
	c1.height = 5

	fmt.Printf("El ancho: %f y el alto: %f del cuboide \n", c1.width, c1.height)
	fmt.Println()
	// instanciar una variable de tipo estrucutra con valores

	// c2 := cuboide{
	// 	width: 15,
	// 	height: 12,
	// 	deep: 20,
	// }
	c2 := cuboide{width: 15, height: 12, deep: 20}

	// formas de mostrar los Struct
	fmt.Printf("forma 1: %v\n", c2)
	fmt.Printf("forma 2: %+v\n", c2)
	fmt.Printf("forma 3: %#v\n", c2)

	fmt.Println()
	fmt.Println("----------------- Punteros a Struct ---------------------")
	fmt.Println()

	p := &cuboide{width: 3, height: 6, deep: 4}
	v := *p

	fmt.Printf("Mostrando desde un puntero: %+v\n", v)

	fmt.Println()
	fmt.Println("----------------- Receptores de funcion y creacion de metodos ---------------------")
	fmt.Println()

	newCubooide := cuboide{
		width:  8,
		height: 6,
		deep:   4,
	}

	fmt.Println(newCubooide.valueCuboide())
	fmt.Println("El volumen del cuboide es de: ", newCubooide.volume())
	fmt.Println("Se redimenciono el cuboide:")
	newCubooide.redesign(10, 8, 6)
	fmt.Println(newCubooide.valueCuboide())
	fmt.Println("El volumen del cuboide es de: ", newCubooide.volume())

	fmt.Println()
	fmt.Println("----------------- Inscrustado de estructra ---------------------")
	fmt.Println()

	product1 := product{
		name:  "Reproductor MP3",
		price: 179,
	}

	fmt.Printf("Producto 1: %+v\n", product1)
	product1.discountProduct(0.15)
	fmt.Printf("Producto %s tiene un valor de %d con el descuento\n", product1.name, product1.price)

	phone1 := phone{
		product: product{
			name:  "Xiaomi",
			price: 800,
		},
		inches:  6,
		battery: 2400,
	}

	fmt.Printf("Telefono 1: %+v\n", phone1)
	phone1.discountProduct(0.10)
	fmt.Printf("Telefono %s tiene un valor de %d con el descuento\n", phone1.name, phone1.price)

	fmt.Println()
	fmt.Println("----------------- Opciones funcionales como alternativa a constructores ---------------------")
	fmt.Println()

	student1 := newStudent()
	fmt.Printf("Estudiante 1: %+v\n", student1)
	student2 := newStudent(name("Gerzon"), birth("2005/10/05"), discount())
	fmt.Printf("Estudiante 2: %+v\n", student2)

}

// Metodos para la estructura cuboide
func (c cuboide) valueCuboide() string {
	return fmt.Sprintf("%v x %v x %v", c.width, c.height, c.deep)
}

func (c cuboide) volume() float64 {
	return c.width * c.height * c.deep
}

func (c *cuboide) redesign(width, height, deep float64) {
	c.width = width
	c.height = height
	c.deep = deep
}

func (p *product) discountProduct(discount float32) {
	p.price = int(float32(p.price) * (1 - discount))
}

// Opciones funcionales
func name(name string) option {
	return func(s *student) {
		s.name = name
	}
}
func birth(birth string) option {
	return func(s *student) {
		s.birth = birth
	}
}
func discount() option {
	return func(s *student) {
		s.discount = true
	}
}
func newStudent(options ...option) student {
	stud := student{}

	for _, option := range options {
		option(&stud)
	}
	return stud
}
