package main

import "fmt"

func (f informacionEmpleados) antiguedadEmpleado() (antiguedad string) {
	antiguedad = "Nombre: " + f.nombre + "\n"
	antiguedad += "Cargo: " + f.cargo + "\n"
	var edadBase int = 55

	if f.edad >= edadBase {
		fmt.Println("Edad necesaria")
	} else {
		fmt.Println("No cumple con la edad")
	}
	var anosLaborandoBase int = 25

	if f.anosLaborando >= anosLaborandoBase {
		fmt.Println("Años cumplidos")
	} else {
		fmt.Println("Aun no cumple los años requeridos")
	}

	fmt.Println()

	return antiguedad
}
