package main

//fmt: para poder escribir en la consola
import (
	"fmt"

	"github.com/sebascavero/godesde0/variables"
)

func main() {

	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
