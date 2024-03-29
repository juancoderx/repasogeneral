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
		{
			empleado: informacionEmpleados{nombre: "Carlos", cargo: "Auxiliar", edad: 21, anosLaborando: 3},
			found:    false,
		},
	}

	var listadoEmpleados Empleados

	for i := 0; i < len(tests); i++ {
		t.Run("", func(t *testing.T) {
			if tests[i].found {
				listadoEmpleados = append(listadoEmpleados, tests[i].empleado)
			}

			revision := listadoEmpleados.eliminarEmpleado(tests[i].empleado)

			if revision != tests[i].found {
				t.Errorf("Se esperaba %t y se obtuvo %t", tests[i].found, revision)
			}

			for _, e := range listadoEmpleados {
				if e.nombre == tests[i].empleado.nombre {
					t.Error("No se elimino")
				}
			}
		})
	}
}
