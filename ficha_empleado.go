package main

import (
	"strconv"
)

func (f informacionEmpleados) fichaEmpleado() (ficha string) {
	ficha = "Nombre: " + f.nombre + "\n"
	ficha += "Cargo: " + f.cargo + "\n"
	ficha += "Edad: " + strconv.Itoa(f.edad) + "\n"
	ficha += "Años Laborando: " + strconv.Itoa(f.anosLaborando) + "\n"

	return ficha
}
