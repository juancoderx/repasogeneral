package main

func (f informacionEmpleados) antiguedadEmpleado() (antiguedad string) {
	antiguedad = "Nombre: " + f.nombre + "\n"
	antiguedad += "Cargo: " + f.cargo + "\n"
	var edadBase int = 55

	if f.edad >= edadBase {
		antiguedad += "Edad: Valida" + "\n"
	} else {
		antiguedad += "Edad: No valida" + "\n"
	}
	var anosLaborandoBase int = 25

	if f.anosLaborando >= anosLaborandoBase {
		antiguedad += "Años: Validos" + "\n"
	} else {
		antiguedad += "Años: No validos" + "\n"
	}

	return antiguedad
}
