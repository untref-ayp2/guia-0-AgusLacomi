package main

import (
	"fmt"
	"guiaCero/estructurasrepetitivas"
	"guiaCero/funciones"
)

// By AgusLacomi
func main() {

	/*Funciones*/
	funciones.PuntoUno(3.0, 2.0, 1.0)
	funciones.PuntoDos()
	minValue, maxValue := funciones.PuntoTres(5, 4, 1, 2)
	fmt.Printf("El valor mínimo es %v y el valor máximo es %v\n", minValue, maxValue)

	/*Estructuras Repetitivas*/
	numero := 5
	factorial := estructurasrepetitivas.PuntoUno(numero)
	fmt.Printf("El factorial de %v es %v\n", numero, factorial)

	a := 5
	b := 2
	producto := estructurasrepetitivas.PuntoDos(a, b)
	fmt.Println("El producto entre", a, "y", b, "es", producto)
}
