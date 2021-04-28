package main

import "fmt"

func changeValue (x int) {
	x = x * x
}
func  changeValue_1(x * int )  { // принимает в качестве параметра указатель на объект типа int.
	*x = (*x) * (*x)
}


func main (){
d := 5
	fmt.Println("d before:", d) // 5
changeValue(d)                      // изменяем значение
	fmt.Println("d after:", d)  // 5 - значение не изменилось
	fmt.Println("*********** передаем переменную в changeValue_1 ***********")
	changeValue_1(&d) // При вызове функции changeValue в нее передается адрес переменной d (changeValue_1(&d))
	fmt.Println("После передачи в changeValue_1 ",d)

}