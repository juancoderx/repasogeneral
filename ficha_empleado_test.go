package main

import "testing"

func Test_informacionEmpleados_fichaEmpleado(t *testing.T) {
	tests := []struct {
		stringEsperado string
		empleado       informacionEmpleados
	}{
		{
			empleado:       informacionEmpleados{nombre: "Juan", cargo: "Jefe", edad: 23, anosLaborando: 2},
			stringEsperado: "Nombre: " + "Juan" + "\n" + "Cargo: " + "Jefe" + "\n" + "Edad: " + "23" + "\n" + "Años Laborando: " + "2" + "\n",
		},
		{
			empleado:       informacionEmpleados{nombre: "Carlos", cargo: "Operador", edad: 21, anosLaborando: 12},
			stringEsperado: "Nombre: " + "Carlos" + "\n" + "Cargo: " + "Operador" + "\n" + "Edad: " + "21" + "\n" + "Años Laborando: " + "12" + "\n",
		},
	}

	for i := 0; i < len(tests); i++ {
		revision := tests[i].empleado.fichaEmpleado()

		if revision != tests[i].stringEsperado {
			t.Error("Se esperaba:", tests[i].stringEsperado)
		}
	}
}
