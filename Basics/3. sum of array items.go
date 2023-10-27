package main
import "fmt"

func main() {
    Arr := []int{5, 10, 15, 20, 25}
    Sum := 0
    for i := 0; i < len(Arr); i++ {
        Sum = Sum + Arr[i]
    }
    fmt.Println("The sum of array items = ", Sum)
}
