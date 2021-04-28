package main
import "fmt"

func main() {
       sum(5,5)
      
         fmt.Println()
}
 func sum (a, b int){
     var c int
      c = a+b
                fmt.Println(a,"Печатаю переменную с")
                fmt.Println(b, "Печатаю переменную b")
                fmt.Println(c,"Печатаю  сумму переменных")
     first:= func(){
                   fmt. Println("Переменная с в first")
                   fmt. Println(c+10)
                   fmt. Println("прибавил к с + 10 в first")
                   
     }
    
     first()
                  fmt.Println("Замкнул")
 }
