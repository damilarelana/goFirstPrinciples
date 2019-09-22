package main
import "fmt"

var degreeFarenheit float64

func main() {
  fmt.Println("Enter the degrees (in F):")
  fmt.Scanf("%f", &degreeFarenheit)
  degreeCelcius := (5 * (degreeFarenheit - 32)) / 9
  fmt.Println("this gives", degreeCelcius,"C")
}
