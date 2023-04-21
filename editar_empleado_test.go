package main

import (
	"testing"
)

func TestEmpleados_editarEmpleado(t *testing.T) {
	tests := []struct {
		empleado, edited informacionEmpleados
		found            bool
	}{
		{
			empleado: informacionEmpleados{nombre: "Juan", cargo: "Jefe", edad: 23, anosLaborando: 4},
			edited:   informacionEmpleados{nombre: "Juan", cargo: "Operador", edad: 25, anosLaborando: 2},
			found:    true,
		},
		{
			empleado: informacionEmpleados{nombre: "Carlos", cargo: "Auxiliar", edad: 21, anosLaborando: 3},
			edited:   informacionEmpleados{nombre: "Carlos", cargo: "Auxiliar", edad: 21, anosLaborando: 3},
			found:    true,
		},
		{
			empleado: informacionEmpleados{nombre: "Pedro", cargo: "Auxiliar", edad: 21, anosLaborando: 3},
			edited:   informacionEmpleados{nombre: "Josue", cargo: "Cocina", edad: 43, anosLaborando: 4},
			found:    false,
		},
	}

	for i := 0; i < len(tests); i++ {
		var empleadosLista Empleados

		empleadosLista = append(empleadosLista, tests[i].empleado)

		revision := empleadosLista.editarEmpleado(tests[i].edited)

		if revision != tests[i].found {
			t.Error("Se esperaba", tests[i].found)
		}

		if !tests[i].found {
			continue
		}

		if empleadosLista[0].nombre != tests[i].edited.nombre {
			t.Error("Se esperaba:", tests[i].edited.nombre)
		}

		if empleadosLista[0].edad != tests[i].edited.edad {
			t.Errorf("Se esperaban %d años de edad", tests[i].edited.edad)
		}

		if empleadosLista[0].cargo != tests[i].edited.cargo {
			t.Error("Se esperaba el cargo:", tests[i].edited.cargo)
		}

		if empleadosLista[0].anosLaborando != tests[i].edited.anosLaborando {
			t.Errorf("Se esperaban %d años laborados", tests[i].empleado.anosLaborando)
		}
	}
}
