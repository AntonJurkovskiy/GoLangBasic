package main

import (
	"fmt"
	"io/ioutil"


)

func main() {
	// Зачитываем содержимое файла
	data, err := ioutil.ReadFile("D:/Go_Lessons/ReadFile/Data.txt")
	
	// Если во время считывания файла произошла ошибка
	// выводим ее
	if err != nil {
		fmt.Println(err)
	}

	// Если чтение данных прошло успено
	// выводим их в консоль
	fmt.Print(string(data))
   

}