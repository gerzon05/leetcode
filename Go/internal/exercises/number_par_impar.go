package exercises

import (
	"fmt"
)

func Run() {
	ParImpar(15)
	ParImpar(30)
}

func ParImpar(num int) {
	fmt.Printf("=================  Numeros Enteros Hasta %d ==========================\n", num)
	for number := 1; number <= num; number++ {
		if number%2 == 0 {
			fmt.Println("Numero Par: ", number)
		} else if number%2 == 1 {
			fmt.Println("Numero Impar: ", number)
		} else {
			fmt.Println("Este Numero no es Par ni Impar: ", number)
		}
	}
	fmt.Println("======================================================================")
}
