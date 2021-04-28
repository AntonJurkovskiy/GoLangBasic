package main

import (
	"fmt"
)

func paintNeeded (w, h float64) {
	area := w*h
	fmt.Println("Результат работы paintNeeded." )
	fmt.Printf("%.2f Литров нужно для покраски стены %v шириной %v длинной\n", area / 10,w,h)
}

func paintNeeded1(w, h float64) float64 {
	area := w * h
	return area / 10
}

func functionB (a bool) {
	fmt.Println(!a)
}

func functionD (s string, n int) {
	for i := 0;i<n;i++ {
		fmt.Println(s)
	}
}

func status (grade float64) string {
	if grade< 60.0 {
		return "Не сдал"
	}
	return "Сдал"
}


func main() {

	fmt.Printf("%12s | %s \n", "Product", "Cost in Cents")
	fmt.Println("-------------------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clipc ", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
	paintNeeded(50.5, 78.4)
	functionB(true)
	functionD("Hello!!!!", 5)
	fmt.Println(status(55.5))
	amount := paintNeeded1(50.0, 60.0)
	fmt.Print("Результат работы paintNeeded1. Для стены надо ", amount)

}