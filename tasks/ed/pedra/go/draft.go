package main
import "fmt"
import "math"
func main() {
    var n int
    fmt.Scan(&n)
    jogadores := make(map([int]int))

    for i := range n{
        var pa, pb float64
        fmt.Scan(&pa, &pb)

        diferenca := math.Abs(pa-pb)
        if (diferenca > 10){
            jogadores[i] = diferenca
        }
    }

    fmt.Print(jogadores)
}
