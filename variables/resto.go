package variables

import (
	"fmt"
	"time"
)

//Las fucniones con inicial en mayuscula se ven en todos los archivos del mismo paquete y donde se importe el paquete
//Las fucniones con inicial en minuscula solo se ven en el archivo donde fueron creados

//Las varibales con inicial en mayuscula se ven en todos los archivos del mismo paquete y donde se importe el paquete
//Las variables con inicial en minuscula solo se ven en el archivo donde fueron creados

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {

	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}
