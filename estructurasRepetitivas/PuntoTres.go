package estructurasrepetitivas

/**
 * @param numero: variable de tipo entero
 *
 * @pre: Se ingresa por parámetro un numero entero para verificar si es un número primo.
 *
 * @post: Retorna verdadero en caso de ser un número primo la variable ingresada por parámetro.
 *
 */
func PuntoTres(numero int) bool {

	if numero <= 1 {
		return false
	}

	for i := 2; i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}
