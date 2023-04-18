package main

type informacionEmpresa struct {
	nombre, presidente, pais     string
	empleados, edades, fundacion int
}

type informacionEmpleados struct {
	nombre, cargo       string
	edad, anosLaborando int
}

type Empleados []informacionEmpleados
