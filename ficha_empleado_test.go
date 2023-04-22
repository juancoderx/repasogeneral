package main

import "testing"

func Test_informacionEmpleados_fichaEmpleado(t *testing.T) {
	tests := []struct {
		stringEsperado string
		empleado       informacionEmpleados
	}{
		{
			empleado:       informacionEmpleados{nombre: "Juan", cargo: "Jefe", edad: 23, anosLaborando: 2},
			stringEsperado: "Nombre: Juan\nCargo: Jefe\nEdad: 23\nAños Laborando: 2\n",
		},
		{
			empleado:       informacionEmpleados{nombre: "Carlos", cargo: "Operador", edad: 21, anosLaborando: 12},
			stringEsperado: "Nombre: Carlos\nCargo: Operador\nEdad: 21\nAños Laborando: 12\n",
		},
	}

	for i := 0; i < len(tests); i++ {
		t.Run("", func(t *testing.T) {
			revision := tests[i].empleado.fichaEmpleado()

			if revision != tests[i].stringEsperado {
				t.Errorf("Se esperaba %s y se obtuvo %s", tests[i].stringEsperado, revision)
			}
		})
	}
}
