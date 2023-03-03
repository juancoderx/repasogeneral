package main

type informacionEmpresa struct {
	nombre, presidente, pais     string
	empleados, edades, fundacion int
}

type informacionEmpleados struct {
	nombre, cargo       string
	edad, añosLaborando int
}

type Empleados []informacionEmpleados
