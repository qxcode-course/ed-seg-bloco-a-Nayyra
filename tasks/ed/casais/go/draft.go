package main
import (
    "fmt"
    "slices"
    )
    
func main() {
    var n int
    fmt.Scan(&n)
    solteiros := make([]int, 0 , n)
    contador := 0

    for i := 0; i < n; i++ {
        var bixo int
        fmt.Scan(&bixo)
        par_bixo := bixo * -1
        posicao := slices.Index(solteiros, par_bixo)
        if posicao != -1{
            solteiros = append(solteiros[:posicao], solteiros[posicao+1:]...)
            contador++
        }else{
            solteiros = append(solteiros, bixo)
        }

    }

    fmt.Println(contador)
}
