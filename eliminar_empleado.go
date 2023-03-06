package main

func (e *Empleados) eliminarEmpleado(nombreEmpleado informacionEmpleados) (exist bool) {
	for i := 0; i < len(*e); i++ {
		if (*e)[i].nombre == nombreEmpleado.nombre {

			*e = append((*e)[:i], (*e)[i+1:]...)

			return true
		}
	}

	return false
}
