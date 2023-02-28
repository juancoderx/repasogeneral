package salario

func Salario(edad int, experiencia bool) int {
	var salarioTotal int = 2000
	if edad >= 25 {
		salarioTotal += 1000
	}

	if experiencia {
		salarioTotal += 3500
	}

	return salarioTotal
}
