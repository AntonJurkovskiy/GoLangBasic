package main

import (
	"errors"
	"fmt"
	"log"
)

func needPaint( w, h float64) (float64, error) {      // на выход подаются два параметра
	if w < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", w)  // обозначиваем условия ошибки
	}
    if h < 0 {
    	return 0, fmt.Errorf("a height of %0.2f is invalid", h)  // обозначиваем условия ошибки
	}
	   area := w * h
	   fmt.Println("Площадь стены = ", area)
	   return area / 10, nil // в случае если ошибки нет то возвращаютс ядва параметра
}

func divide (dividend, divisor float64)  (float64, error){
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
} // Еще вариант обьявление ошибки в функции



func main() {
	fmt.Println(      "**********************************************************************")
	amount, err := needPaint(-50.0, 60.0)   // в функцию на фход подаются два параметра

	if err != nil {              // Вариант 1. Вывода ошибки.
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f нужно литров", amount)
	}

	amount1, err := needPaint(100.0,200.0)
	if err != nil {     // Вариант 2. Вывода ошибки.
		log.Fatal(err)
	}
	fmt.Printf("%0.2f нужно литров\n", amount1)
	           // вызов функции  divide
	quotient, err := divide(50.0,6.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print( "****************************** func divide *************************************")
	fmt.Printf("%0.2f\n",quotient)

}