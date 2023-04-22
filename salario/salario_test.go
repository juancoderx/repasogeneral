package salario

import "testing"

func TestSalario(t *testing.T) {
	tests := []struct {
		edad, result int
		exp          bool
	}{
		{edad: 26, result: 6500, exp: true},
		{edad: 34, result: 6500, exp: true},
		{edad: 29, result: 6500, exp: true},
	}

	for i := 0; i < len(tests); i++ {
		resultado := Salario(tests[i].edad, tests[i].exp)

		if resultado != tests[i].result {
			t.Errorf("Se esperaba %d y se obtuvo %d", tests[i].result, resultado)
		}
	}
}
