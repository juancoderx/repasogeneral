package main

import (
	"testing"
)

func Test_informacionEmpleados_antiguedadEmpleado(t *testing.T) {
	tests := []struct {
		stringEsperado string
		empleado       informacionEmpleados
	}{
		{
			empleado:       informacionEmpleados{nombre: "Juan", cargo: "Jefe", edad: 55, anosLaborando: 25},
			stringEsperado: "Nombre: Juan\nCargo: Jefe\nEdad: Valida\nAños: Validos\n",
		},
		{
			empleado:       informacionEmpleados{nombre: "Carlos", cargo: "Operador", edad: 21, anosLaborando: 12},
			stringEsperado: "Nombre: Carlos\nCargo: Operador\nEdad: No valida\nAños: No validos\n",
		},
	}

	for i := 0; i < len(tests); i++ {
		t.Run("", func(t *testing.T) {
			revision := tests[i].empleado.antiguedadEmpleado()
			if revision != tests[i].stringEsperado {
				t.Errorf("Se esperaba %s y se obtuvo %s", tests[i].stringEsperado, revision)
			}
		})
	}
}
