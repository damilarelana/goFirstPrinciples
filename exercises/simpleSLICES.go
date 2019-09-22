package main
import "fmt"

var (
  yTotal int64
  xTotal int64
)

func main() {
  initialArray := [7] int64{1,2,3,4,5,6,7}
  y := initialArray[:]
  x := make([]int64, 5)
  copy(x, y)
  for _, value := range x {
    xTotal += value
  }
  for _, value := range y {
    yTotal += value
  }
  fmt.Println("x :", x)
  fmt.Println("y :", y)
  fmt.Println("Sumy - Sumx =", yTotal - xTotal)
  fmt.Println(initialArray[2:5])
}
