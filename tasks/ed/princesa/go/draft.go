package main
import "fmt"
func main() {
    var n, e int
    fmt.Scan(&n, &e)
    pessoas := make([]int, n)
    for i := 0; i < n; i++ {
        pessoas[i] = i + 1
    }

    espadaAtual := e - 1
    for len(pessoas) >= 1{

        fmt.Print("[ ")
        for i := 0; i < len(pessoas); i++{
            if ( i == espadaAtual){
                fmt.Printf("%d> ", pessoas[i])
            }else{
                fmt.Printf("%d ", pessoas[i])
            }

        }
        fmt.Println("]")

        if len(pessoas) == 1 {
            break
        }

        indiceMorto := (espadaAtual + 1) % len(pessoas)

        pessoas = append(pessoas[:indiceMorto], pessoas[indiceMorto+1:]...)


        espadaAtual = indiceMorto % len(pessoas)
    }

}
