package main

import
	"fmt"
                   // ============ блок 1 =============
func main(){
var x int = 4  // определяем переменную
var p *int     // определяем указатель. Переменная p указатель на обьект типа int
p = & x        // указатель получает адрес  конкретной переменной x
fmt.Println(p)
	                 // ============ блок 2 =============
var a int = 10
var b *int = &a                  // указатель получает адрес переменной
fmt.Println("Address: ",b)  // значение указателя - адрес переменной x
fmt.Println("Value: ",*b)    // значение переменной x
	           // ============ сокращенная форма =============
	      f := 2.3
	      pf := &f
	      fmt.Println("Address: ",pf)
	      fmt.Println("Address: ",*pf)
	                // ============ Функция new =============
	c := new(int)
	fmt.Println("Value 'c' : ",*c,)  // Value: 0 - значение по умолчанию
	*c = 8                               // изменяем значение
	fmt.Println("Value 'c' : ",*c)  // Value: 8
	fmt.Println("++++++ На приметре строк ++++++++")
	myStringAnton := "Anton"
	fmt.Println(myStringAnton)
	myStringJulia := &myStringAnton
	*myStringJulia = "Julia "+"Anton"
	fmt.Println(myStringJulia)



}