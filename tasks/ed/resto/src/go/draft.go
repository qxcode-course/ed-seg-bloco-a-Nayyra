package main
import "fmt"

func divRest(n int){
    if n == 0{
        return
    }

    div := n / 2
    resto := n % 2
    divRest(div)
    fmt.Printf("%d %d\n", div, resto)

}


func main() {
    var n int
    fmt.Scan(&n)
    divRest(n)
}

