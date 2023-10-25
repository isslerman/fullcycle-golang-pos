package main

import "fmt"

func main() {
	salarios := map[string]int{"Marcao": 6000, "Boss": 10000, "Bekka": 14550, "Zé": 12}
	fmt.Println(salarios["Marcao"]) // 6000
	fmt.Println(salarios["Other"])  // 0
	fmt.Println(salarios["Ze"])     // 0
	fmt.Println(salarios["Zé"])     // 12
	delete(salarios, "Zé")
	fmt.Println(salarios["Zé"]) // 0

	// others type
	sal01 := make(map[string]int)
	sal02 := map[string]int{}
	sal01["Marcao"] = 18000
	sal02["Marcao"] = 22000
	fmt.Println(sal01["Marcao"]) // 18000
	fmt.Println(sal02["Marcao"]) // 22000

	for nome, salario := range salarios {
		fmt.Printf("O salário do %s é de %d.\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é de %d.\n", salario)
	}

	total := 0
	for _, salario := range salarios {
		total = total + salario
		fmt.Printf("O total está em: %d.\n", total)
	}

}
