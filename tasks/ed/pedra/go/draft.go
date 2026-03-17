package main
import "fmt"
import "math"
func main() {
    var n int
    fmt.Scan(&n)
    // pontos := make([]float64, n)
    mapeamento := make(map[int]float64)

    for i := 0; i <= n; i++{
        var pa, pb float64
        fmt.Scan(&pa, &pb)
        if (pa >= 10) && (pb  >= 10){
        
            diferenca := math.Abs(pa-pb)
            mapeamento[i] = diferenca
            
        }
    }

    menor := 10.0
    vencendor := 0


    for i, a := range mapeamento{
        if (a <= menor){
            vencendor = i
        }
    }
    
    if (len(mapeamento) == 0){
        fmt.Println("sem ganhador")
    }else {
        
        fmt.Println(vencendor)
    }

}
