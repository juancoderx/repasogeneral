package main

import "fmt"

func (f informacionEmpleados) antiguedadEmpleado() (antiguedad string) {
	antiguedad = "Nombre: " + f.nombre + "\n"
	antiguedad += "Cargo: " + f.cargo + "\n"

	if f.edad >= 55 {
		fmt.Println("Edad necesaria")
	} else {
		fmt.Println("No cumple con la edad")
	}

	if f.añosLaborando >= 25 {
		fmt.Println("Años cumplidos")
	} else {
		fmt.Println("Aun no cumple los años requeridos")
	}

	fmt.Println()

	return antiguedad
}
