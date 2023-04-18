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
			stringEsperado: "Nombre: " + "Juan" + "\n" + "Cargo: " + "Jefe" + "\n",
		},
		{
			empleado:       informacionEmpleados{nombre: "Carlos", cargo: "Operador", edad: 21, anosLaborando: 12},
			stringEsperado: "Nombre: " + "Josue" + "\n" + "Cargo: " + "Soporte" + "\n",
		},
	}

	for i := 0; i < len(tests); i++ {
		revision := informacionEmpleados.antiguedadEmpleado(tests[i].empleado)

		if revision != tests[i].stringEsperado {
			t.Error("Se esperaba:", tests[i].stringEsperado)
		}
	}
}
