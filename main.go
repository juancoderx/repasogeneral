package main

import (
	"errors"
	"fmt"

	"github.com/juancoderx/repasogeneral/salario"
)

type option int8

func (o option) String() string {
	switch o {
	case opcionDatosEmpresa:
		return fmt.Sprintf("%d. %s", opcionDatosEmpresa, "Nueva empresa")

	case opcionBuscarEmpresa:
		return fmt.Sprintf("%d. %s", opcionBuscarEmpresa, "Buscar una Empresa")

	case opcionEliminarEmpresa:
		return fmt.Sprintf("%d. %s", opcionEliminarEmpresa, "Eliminar una Empresa")

	case opcionEditarEmpresa:
		return fmt.Sprintf("%d. %s", opcionEditarEmpresa, "Editar una Empresa")

	case opcionPresentarEmpresa:
		return fmt.Sprintf("%d. %s", opcionPresentarEmpresa, "Presentanción de una Empresa")

	case opcionRegistarEmpleado:
		return fmt.Sprintf("%d. %s", opcionRegistarEmpleado, "Registar Empleado")

	case opcionEditarEmpleado:
		return fmt.Sprintf("%d. %s", opcionEditarEmpleado, "Editar Empleado")

	case opcionEliminarEmpleado:
		return fmt.Sprintf("%d. %s", opcionEliminarEmpleado, "Eliminar Empleado")

	case opcionJubilacion:
		return fmt.Sprintf("%d. %s", opcionJubilacion, "Revisar Jubilacion")

	case opcionCalcularSalario:
		return fmt.Sprintf("%d. %s", opcionCalcularSalario, "Calcular Salario")

	}

	return ""
}

const (
	opcionDatosEmpresa option = iota + 1
	opcionBuscarEmpresa
	opcionEliminarEmpresa
	opcionEditarEmpresa
	opcionPresentarEmpresa
	opcionRegistarEmpleado
	opcionEditarEmpleado
	opcionEliminarEmpleado
	opcionJubilacion
	opcionCalcularSalario
)

var (
	empresasCreadas  []informacionEmpresa
	listadoEmpleados Empleados
)

func main() {
	var eleccion option = 1
	err := errors.New("no hay empresas registradas")

	for eleccion >= 1 {
		fmt.Println("¡Registro de Empresas!")
		fmt.Println()

		fmt.Println("Empresas registradas:")
		fmt.Println(empresasCreadas)
		fmt.Println()

		fmt.Println("Empleados registados:")
		fmt.Println(listadoEmpleados)
		fmt.Println()

		fmt.Println(opcionDatosEmpresa)
		fmt.Println(opcionBuscarEmpresa)
		fmt.Println(opcionEliminarEmpresa)
		fmt.Println(opcionEditarEmpresa)
		fmt.Println(opcionPresentarEmpresa)
		fmt.Println(opcionRegistarEmpleado)
		fmt.Println(opcionEditarEmpleado)
		fmt.Println(opcionEliminarEmpleado)
		fmt.Println(opcionJubilacion)
		fmt.Println(opcionCalcularSalario)

		fmt.Println()

		fmt.Print(">")
		fmt.Scan(&eleccion)

		switch eleccion {
		case opcionDatosEmpresa:
			var datosIniciales informacionEmpresa

			fmt.Println("Ingresa el nombre de la Empresa")
			fmt.Print(">")
			fmt.Scan(&datosIniciales.nombre)

			fmt.Println("Ingresa el nombre del Presidente")
			fmt.Print(">")
			fmt.Scan(&datosIniciales.presidente)

			fmt.Println("Año de fundación de la Empresa")
			fmt.Print(">")
			fmt.Scan(&datosIniciales.fundacion)

			fmt.Println("Pais de la Empresa")
			fmt.Print(">")
			fmt.Scan(&datosIniciales.pais)

			fmt.Println("¿Cuantos empleados tiene su empresa?")
			fmt.Print(">")
			fmt.Scan(&datosIniciales.empleados)

			fmt.Println("Edad promedio de los empleados")
			fmt.Print(">")
			fmt.Scan(&datosIniciales.edades)

			empresasCreadas = append(empresasCreadas, datosIniciales)

		case opcionBuscarEmpresa:
			if len(empresasCreadas) == 0 {
				fmt.Println(err)

				break
			}

			var (
				buscarEmpresa string
				found         bool
			)

			fmt.Println("Ingrese nombre de la Empresa")
			fmt.Print(">")
			fmt.Scan(&buscarEmpresa)

			for i := 0; i < len(empresasCreadas); i++ {
				if empresasCreadas[i].nombre == buscarEmpresa {
					found = true

					fmt.Println("La empresa se encuentra registrada, ve a Presentación para conocerla")

					break
				}
			}

			if !found {
				fmt.Println("Empresa no registrada, intente nuevamente.")
			}

		case opcionEliminarEmpresa:
			if len(empresasCreadas) == 0 {
				fmt.Println(err)

				break
			}

			var (
				eliminarEmpresa string
				found           bool
			)

			fmt.Println("Ingrese nombre de la empresa a eliminar")
			fmt.Print(">")
			fmt.Scan(&eliminarEmpresa)

			for i := 0; i < len(empresasCreadas); i++ {
				if empresasCreadas[i].nombre == eliminarEmpresa {
					found = true
					empresasCreadas = append(empresasCreadas[:i], empresasCreadas[i+1:]...)

					break
				}
			}

			if !found {
				fmt.Println("No hay empresa con ese nombre.")
			}

		case opcionEditarEmpresa:
			if len(empresasCreadas) == 0 {
				fmt.Println(err)

				break
			}

			var (
				editarEmpresa string
				found         bool
			)

			fmt.Println("Ingrese el Nombre del usuario a editar")
			fmt.Print(">")
			fmt.Scan(&editarEmpresa)

			for i := 0; i < len(empresasCreadas); i++ {
				if empresasCreadas[i].nombre == editarEmpresa {
					found = true

					fmt.Println("Ingrese el nuevo Nombre de la Empresa")
					fmt.Print(">")
					fmt.Scan(&empresasCreadas[i].nombre)

					fmt.Println("Ingrese el nuevo Presidente")
					fmt.Print(">")
					fmt.Scan(&empresasCreadas[i].presidente)

					fmt.Println("Ingrese el año de fundación")
					fmt.Print(">")
					fmt.Scan(&empresasCreadas[i].fundacion)

					fmt.Println("Ingrese el numero de empleados")
					fmt.Print(">")
					fmt.Scan(&empresasCreadas[i].empleados)

					fmt.Println("Ingrese la edad promedio de los empleados")
					fmt.Print(">")
					fmt.Scan(&empresasCreadas[i].edades)

					break

				}
			}

			if !found {
				fmt.Println("No hemos encontrado la empresa solicitada")
			}

		case opcionPresentarEmpresa:
			if len(empresasCreadas) == 0 {
				fmt.Println(err)

				break
			}

			var nombreEmpresa string

			fmt.Println("¿Cual empresa desea visualizar?")
			fmt.Print(">")
			fmt.Scan(&nombreEmpresa)

			presentarEmpresa(nombreEmpresa)

		case opcionRegistarEmpleado:
			var datosEmpleados informacionEmpleados

			fmt.Println("Nombre del Empleado")
			fmt.Print(">")
			fmt.Scan(&datosEmpleados.nombre)

			fmt.Println("Cargo del empleado")
			fmt.Print(">")
			fmt.Scan(&datosEmpleados.cargo)

			fmt.Println("Edad del empleado")
			fmt.Print(">")
			fmt.Scan(&datosEmpleados.edad)

			fmt.Println("Años Laborando en la empresa")
			fmt.Print(">")
			fmt.Scan(&datosEmpleados.añosLaborando)

			listadoEmpleados = append(listadoEmpleados, datosEmpleados)

			fmt.Println(datosEmpleados.fichaEmpleado())

		case opcionJubilacion:
			var revisonJubilacion informacionEmpleados

			fmt.Println("Nombre del Empleado")
			fmt.Print(">")
			fmt.Scan(&revisonJubilacion.nombre)

			fmt.Println("Cargo del Empleado")
			fmt.Print(">")
			fmt.Scan(&revisonJubilacion.cargo)

			fmt.Println("Edad del empleado")
			fmt.Print(">")
			fmt.Scan(&revisonJubilacion.edad)

			fmt.Println("Años que ha laborado")
			fmt.Print(">")
			fmt.Scan(&revisonJubilacion.añosLaborando)

			fmt.Println(revisonJubilacion.antiguedadEmpleado())

		case opcionEditarEmpleado:
			var editarEmpleado informacionEmpleados

			fmt.Println("Ingrese el nombre del empleado a Editar")
			fmt.Print(">")
			fmt.Scan(&editarEmpleado.nombre)

			fmt.Println("Ingrese el cargo del empleado")
			fmt.Print(">")
			fmt.Scan(&editarEmpleado.cargo)

			fmt.Println("Ingrese la edad del empleado")
			fmt.Print(">")
			fmt.Scan(&editarEmpleado.edad)

			fmt.Println("Ingrese los años de servicio del empleado")
			fmt.Print(">")
			fmt.Scan(&editarEmpleado.añosLaborando)

			if exist := listadoEmpleados.editarEmpleado(editarEmpleado); !exist {
				fmt.Println("El empleado no fue encontrado")
			}

		case opcionEliminarEmpleado:
			var eliminarEmpleado informacionEmpleados

			fmt.Println("Ingrese el nombre del empleado a eliminar")
			fmt.Print(">")
			fmt.Scan(&eliminarEmpleado.nombre)

			if exist := listadoEmpleados.eliminarEmpleado(eliminarEmpleado); !exist {
				fmt.Println("El empleado no fue encontrado")
			}

		case opcionCalcularSalario:
			var (
				edad                 int
				experienciaVerificar string
				experiencia          bool
			)

			fmt.Println("Ingrese edad del empleado a calcular")
			fmt.Print(">")
			fmt.Scan(&edad)

			fmt.Println("¿El empleado tiene experiencia? Y o N")
			fmt.Print(">")
			fmt.Scan(&experienciaVerificar)

			if experienciaVerificar == "Y" {
				experiencia = true
			}

			fmt.Print("El salario según las caracteristicas es de: $")
			fmt.Println(salario.Salario(edad, experiencia))
			fmt.Println()

		}
	}
}

func presentarEmpresa(empresa string) {
	for i := 0; i < len(empresasCreadas); i++ {
		if empresasCreadas[i].nombre == empresa {
			fmt.Println("Su presidente:", empresasCreadas[i].presidente)
			fmt.Println("Año de su fundación:", empresasCreadas[i].fundacion)
			fmt.Println("Empresa creada en:", empresasCreadas[i].pais)
			fmt.Println("Cuenta con:", empresasCreadas[i].empleados, "empleados")
			fmt.Println("La edad promedio de los empleados es:", empresasCreadas[i].edades, "años")
			fmt.Println()

			break
		}
	}
}
