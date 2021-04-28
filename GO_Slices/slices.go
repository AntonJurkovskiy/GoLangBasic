package main

import (
	"fmt"
	"strings"
)

func deleteSpace(s []string) {
	for i := range s {
		s[i] = strings.TrimSpace(s[i])

	}
	fmt.Println(s)
	text := s
	fmt.Println(strings.Join(text, "[:os:]"))

	// first:= func ()  {
	// 	for i := range s {
	// 		s[i] = strings.Join(s, "")

	// 	  }
	// 	      first()

	//    }
}

func main() {
	planets := []string{"Земля ", "Луна ", "Солнце ", "Марс ", "Нептун"}
	deleteSpace(planets)
	//fmt.Println(strings.Join(planets, ""))
}
