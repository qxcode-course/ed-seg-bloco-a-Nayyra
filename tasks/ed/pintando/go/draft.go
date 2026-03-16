package main
import "fmt"
import "math"
func main() {
    var a, b, c float32
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

    var p float32 = (a + b + c) / 2

    var area float32 = math.Sqrt(p*(p-a)*(p-b)*(p-c))

    fmt.printf("%.2f\n", a)
}
