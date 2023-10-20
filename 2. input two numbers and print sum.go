package main
import ("fmt")

func main(){
  var x, y,z float32
  fmt.Print("Enter the first number = ")
  fmt.Scan(&x)
  fmt.Print("Enter the second number = ")
  fmt.Scan(&y)
  z = x+y
  fmt.Println("The sum of ",x," and ",y," is = ",z)
}
  
