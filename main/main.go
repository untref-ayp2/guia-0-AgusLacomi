package main

import (
	"fmt"
	"guiaCero/funciones"
)

// By AgusLacomi
func main() {

	/*Funciones*/
	funciones.PuntoUno(3.0, 2.0, 1.0)
	funciones.PuntoDos()
	minValue, maxValue := funciones.PuntoTres(5, 4, 1, 2)
	fmt.Printf("El valor mínimo es %v y el valor máximo es %v\n", minValue, maxValue)
}
