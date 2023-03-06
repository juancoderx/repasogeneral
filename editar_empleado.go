package main

func (f Empleados) editarEmpleado(infoEmpleado informacionEmpleados) (exist bool) {
	for i := 0; i < len(f); i++ {
		if (f)[i].nombre == infoEmpleado.nombre {
			f[i].nombre = infoEmpleado.nombre
			f[i].cargo = infoEmpleado.cargo
			f[i].edad = infoEmpleado.edad
			f[i].añosLaborando = infoEmpleado.añosLaborando

			return true
		}
	}

	return false
}
