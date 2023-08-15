package ejercicios

import "strconv"

func ConvNumerico(texto string) (int, string) {

	// La funcion Atoi regresa 2 variables (valor, error), para omitir el error, se coloca una "_"
	//retorno1, _ := strconv.Atoi(texto)

	retorno1, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}

	if retorno1 > 100 {
		return retorno1, "Es mayor a 100"
	} else {
		return retorno1, "Es menor a 100"
	}
}
