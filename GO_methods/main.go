package main
import "fmt"

type celsius float64        // Новый тип celsius
type kelvin float64        // Новый тип kelvin
type fahrenheit float64
               // Func_1 _________________________________
func kelvinToCelsius (k kelvin) celsius{
	return celsius (k-273.15)
}
                // Func_2 ___________________________________
func celsiusToKelvin(c celsius) kelvin{
	return kelvin(c+273.15)
}
                // Method_1___________________________________
func(k kelvin)celsius()celsius{
	return celsius(k-273.15)
}
func (f fahrenheit) celsius()celsius{
	return celsius((f-32)*5.0/9.0)
}

func main(){
var k kelvin = 294
var c celsius = 36.6
var f fahrenheit = 294

fmt.Println(c)
fmt.Println(celsiusToKelvin(127),"° K is Convert celsius To Kelvin")
fmt.Println(k.celsius()," Конвертация Кельвин в Цельсий")
fmt.Println(f.celsius(),"Конвертация Форенгейт в Цельсий")
}

