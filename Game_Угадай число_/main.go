package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("****** Генерация случайного числа *******")
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я выбрал случайное число от 1 до 100 !!")
	fmt.Println("Вы можете догадаться ?")
	//fmt.Println(target)
	fmt.Println("****** Чтение ввода с клавиатуры *******")
	reader := bufio.NewReader(os.Stdin) // создаем Reader для чтения с ввода с клавиатуры
	success := false

	for guess := 0; guess < 10; guess++ {
		fmt.Println("У вас осталось ", 10-guess, "попыток")

		fmt.Println("Угадай:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("****** Сравнение введенного числа с загаданым *******")
		if guess < target {
			fmt.Println("Ой. Ваше предположение было МЕНЬШЕ..")
		} else if guess > target {
			fmt.Println("Ой. Ваше предположение было БОЛЬШЕ..")
		} else {
			success = true
			fmt.Println("Вы угадали!!!")
			break
		}
	}
	if !success {
		fmt.Println("Извините, вы не угадали число. Это было :", target)
	}
}
