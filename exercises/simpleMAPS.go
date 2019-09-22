package main
import "fmt"

var (
  yTotal int64
  xTotal int64
)

func main() {
  outerMap := map[int64]map[string]int64 {
    1: map[string]int64{
      "Age":2,
    },
  }
  if innerMap, ok := outerMap[2]; ok{
    fmt.Println(innerMap["Age"])
  } else {
    fmt.Println("key:value pair does not exist")
  }
}
