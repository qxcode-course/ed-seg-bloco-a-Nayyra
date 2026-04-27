package main
import "fmt"

func matchingStrings(s []string, c []string) []int{
    mapa := make(map[string]int)
    for _, v := range s{
        mapa[v]++
    }

    r := []int{}
    for _, v := range c{
        r = append(r, mapa[v])
    }

    return r
}

func main() {
    var n int
    fmt.Scan(&n)
    s := make([]string, n)
    for i := 0; i < n; i++{
        fmt.Scan(&s[i])
    }

    var q int
    fmt.Scan(&q)
    c := make([]string, q)
    for i := 0; i < q; i++{
        fmt.Scan(&c[i])
    }

    eita := matchingStrings(s, c)
    for i:=0; i < len(eita); i++{
        if i == len(eita)-1{
            fmt.Print(eita[i], "\n")
        }else{
            fmt.Print(eita[i], " ")
        }
    }
}
