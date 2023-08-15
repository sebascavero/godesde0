package main

import (
	"fmt"

	"github.com/sebascavero/godesde0/ejercicios"
)

//fmt: para poder escribir en la consola

func main() {

	/*estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	//Se puede declarar y hacer la comparación en la misma linea siempre y cuando se separé por ";"
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("El sistema operativo NO es Windows, es ", os)
	} else {
		fmt.Println("El sistema operativo es Windows,", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf(os)
	}*/

	retorno1, retorno2 := ejercicios.ConvNumerico("200")
	fmt.Println(retorno1)
	fmt.Println(retorno2)
}
