package main

import (
	"testing"
)

func Test_presentarEmpresa(t *testing.T) {
	tests := []struct {
		stringEsperado string
		empresas       informacionEmpresa
	}{
		{
			empresas:       informacionEmpresa{nombre: "Twitter", presidente: "Elon Musk", pais: "NA", empleados: 15000, edades: 25, fundacion: 2000},
			stringEsperado: "Presidente: Elon Musk\nFundacion: 2000\nPais: NA\nEmpleados: 15000\nEdades: 25\n",
		},
		{
			empresas:       informacionEmpresa{},
			stringEsperado: "Presidente: \nFundacion: 0\nPais: \nEmpleados: 0\nEdades: 0\n",
		},
		{
			empresas:       informacionEmpresa{},
			stringEsperado: "",
		},
	}

	for i := 0; i < len(tests); i++ {
		empresasCreadas = []informacionEmpresa{}

		if tests[i].stringEsperado != "" {
			empresasCreadas = append(empresasCreadas, tests[i].empresas)
		}

		datos := presentarEmpresa(tests[i].empresas.nombre)

		if datos != tests[i].stringEsperado {
			t.Errorf("Se esperaba %s y se obtuvo %s", tests[i].stringEsperado, datos)
		}
	}
}

func Test_option_String(t *testing.T) {
	tests := []struct {
		stringEsperado string
		opcion         option
	}{
		{
			opcion:         opcionDatosEmpresa,
			stringEsperado: "1. Nueva Empresa",
		},
		{
			opcion:         opcionBuscarEmpresa,
			stringEsperado: "2. Buscar una Empresa",
		},
		{
			opcion:         opcionEliminarEmpresa,
			stringEsperado: "3. Eliminar una Empresa",
		},
		{
			opcion:         opcionEditarEmpresa,
			stringEsperado: "4. Editar una Empresa",
		},
		{
			opcion:         opcionPresentarEmpresa,
			stringEsperado: "5. Presentación de una Empresa",
		},
		{
			opcion:         opcionRegistarEmpleado,
			stringEsperado: "6. Registar Empleado",
		},
		{
			opcion:         opcionEditarEmpleado,
			stringEsperado: "7. Editar Empleado",
		},
		{
			opcion:         opcionEliminarEmpleado,
			stringEsperado: "8. Eliminar Empleado",
		},
		{
			opcion:         opcionJubilacion,
			stringEsperado: "9. Revisar Jubilacion",
		},
		{
			opcion:         opcionCalcularSalario,
			stringEsperado: "10. Calculo Salario",
		},
		{
			opcion:         0,
			stringEsperado: "",
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].opcion.String() != tests[i].stringEsperado {
			t.Errorf("Se esperaba %s y se obtuvo %s", tests[i].stringEsperado, tests[i].opcion.String())
		}
	}
}
