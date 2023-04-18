package salario

func Salario(edad int, experiencia bool) int {
	var salarioTotal int = 2000
	var edadBase int = 26

	if edad >= edadBase {
		salarioTotal += 1000
	}

	if experiencia {
		salarioTotal += 3500
	}

	return salarioTotal
}
