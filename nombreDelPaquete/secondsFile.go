package nombreDelPaquete

import (
	"fmt"
)

func PublicFunc() int {
	privateFunc()
	return 0
}

func privateFunc() {
	var numero int64 = 64000
	numero01 := 1992
	numero02 := 2.2
	fmt.Println(numero, numero01, numero02)

	var boleano bool = true
	boleano02 := false
	fmt.Println(boleano, boleano02)

	a, b, c, d, e := privateFuncRetunsMult()
	fmt.Println(a, b, c, d, e)
}

func privateFuncRetunsMult() (int64, int, float64, bool, bool) {
	return 64000, 1992, 2.2, true, false
}
