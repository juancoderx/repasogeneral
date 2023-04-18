package main

import "testing"

func TestEmpleados_eliminarEmpleado(t *testing.T) {
	tests := []struct {
		empleado informacionEmpleados
		found    bool
	}{
		{
			empleado: informacionEmpleados{nombre: "Juan", cargo: "Jefe", edad: 23, anosLaborando: 4},
			found:    true,
		},
		{
			empleado: informacionEmpleados{nombre: "Carlos", cargo: "Auxiliar", edad: 21, anosLaborando: 3},
			found:    true,
		},
	}

	for i := 0; i < len(tests); i++ {
		listadoEmpleados = append(listadoEmpleados, tests[i].empleado)
		revision := listadoEmpleados.eliminarEmpleado(tests[i].empleado)

		if revision != tests[i].found {
			t.Error("Se esperaba", tests[i].found)
		}
	}
}
