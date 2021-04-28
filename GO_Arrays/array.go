package main

import "fmt"

func main() {
	//__________________________________________________
	var planets [8]string // Вариант 1.
	planets[0] = "Меркурий, "
	planets[1] = "Венера, "
	planets[2] = "Земля, "
	//_________________________________________________
	family := [5]string{"Юля ", "Антон", "Мика ", "Лютик", "...."} // Вариант 2.
	//__________________________________________________
	familyGrischenko := [...]string{"Анжела  ", "Сережа ", "Мижа "} // Вариант 3.ЗАПЯТАЯ В КОНЦЕ ОБЯЗАТЕЛЬЬНА!!
	//__________________________________________________
	fmt.Printf("%T\n", planets)
	fmt.Println(planets)
	fmt.Println("Array Lens ", len(planets))
	fmt.Println(family)
	fmt.Println(familyGrischenko)
	fmt.Println(len(family))
}
