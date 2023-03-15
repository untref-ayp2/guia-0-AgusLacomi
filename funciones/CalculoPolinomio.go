package funciones

//By AgusLacomi
/*
	https://www.digitalocean.com/community/tutorials/an-introduction-to-the-strings-package-in-go-es
*/
func CalculoPolinomio(coeficientes ...float32) string {

	total := 0

	// el total inicial es 0
	total := 0
	// recorrer todos los numeros
	for _, numero := range numeros {
		// en cada iteración sumar al total el valor del numero
		total = numero + total
	}
	// retornar el valor total

	return total
}
