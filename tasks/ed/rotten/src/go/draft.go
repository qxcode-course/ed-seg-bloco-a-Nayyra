package main
import "fmt"
func main() {
    var m, n int
    var min, frescas int
    var podres [][]int
    fmt.Scan(&m)
    fmt.Scan(&n)

    laranjas := make([][]int, m)

    for i := 0; i < m; i++{
        laranjas[i] = make([]int, n)
    }

    for i :=0; i<m; i++{
        for j:=0; j<n; j++{
            fmt.Scan(&laranjas[i][j])

            if laranjas[i][j] == 1{
                frescas++
            }else if laranjas[i][j] == 2{
                podres = append(podres, []int{i,j})
            }
            
        }
    }
    
    for len(podres) >0 && frescas>0{
        for i:=0; i < len(podres); i++{
            atual := podres[0]
            podres = podres[1:]

            dire := [][]int{
                {-1, 0}, {0, -1}, {1,0}, {0,1},
            }

            for _, i:= range dire{
                newL := atual[0]+ i[0]
                newC := atual[1]+i[1]
                if newL >= 0 && newL < m && newC >= 0 && newC < n && laranjas[newL][newC] == 1 {
                    laranjas[newL][newC] = 2
                    frescas--
                    podres = append(podres, []int{newL, newC})
                }
            }
        }
        min++
    }

    if frescas == 0{
        fmt.Println(min)
    }else{
        fmt.Println("-1")
    }

}
