package internal

import "fmt"

func MapssRun() {
	fmt.Println("Diccionarios (MAPAS)")

	capitals := map[string]string{
		"Colombia": "Bogota",
		"España":   "Madrid",
		"Francia":  "París",
		"Japon":    "Tokio",
	}

	fmt.Println("Valores guardados en un mapa: ", capitals)

	fmt.Printf("\n1. Agregando un valor al mapa\n\n")

	capitals["Italia"] = "Roma"

	fmt.Println("Se agrego un elemento al mapa ", capitals)

	country := "Brasil"
	capital, ok := capitals[country]

	if ok {
		fmt.Printf("La capital de %s  es: %s\n", country, capital)
	} else {
		fmt.Printf("el pais de %s no se encuentra disponible\n", country)
	}

	fmt.Printf("\n2. Eliminar un valor del mapa\n\n")

	fmt.Println("Valores del mapa actual: ", capitals)

	delete(capitals, "Francia")

	fmt.Println("Valores al eliminar un valor del mapa: ", capitals)

	fmt.Printf("\n3. Recorrer un mapa\n\n")

	for index, capital := range capitals {
		fmt.Printf("La capital de %s es %s\n", index, capital)
	}

	fmt.Printf("\n4. Contar los calores del mapa\n\n")

	fmt.Println("La cantidad de valores del mapa es de: ", len(capitals))

}
