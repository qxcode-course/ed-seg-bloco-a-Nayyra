package main
import "fmt"
func main() {
   var qtd_a, qtd_b int
   figus := []int{}
   unicas := make(map[int]bool)
   repetidas := make([]int, qtd_a)
   fmt.Scan(&qtd_a)
   fmt.Scan(&qtd_b)

   for i := 0; i < qtd_b; i++{
        var f int
        fmt.Scan(&f)
        figus = append(figus, f)
   }

   for _, v := range figus{
     _, existe := unicas[v]

     if existe{
          repetidas = append(repetidas, v)
     }else{
          unicas[v] = true
     }
   }

   falta := []int{}

   for i := 1; i <= qtd_a; i++{
     if !unicas[i]{
          falta = append(falta, i)
     }
   }

     saida := fmt.Sprintf("%v", repetidas)

     if saida == "[]"{
          fmt.Println("N")
     }else {
          fmt.Println(saida[1: len(saida)-1])
     }

     saidaa := fmt.Sprintf("%v", falta)

     if saidaa == "[]"{
          fmt.Println("N")
     }else {
          fmt.Println(saidaa[1: len(saidaa)-1])
     }

//    for i :=1; i < len(figus); i++{
//         if(figus[i] == figus[i-1]){
//             repetidas = append(repetidas, figus[i])

//         }
//    } 

   
//    if len(figus) == 0{
//      fmt.Println("N")
//    }else{
//      for _, valor := range repetidas{
//           fmt.Printf("%d ", valor)
//      }
//    }

//    for i :=0; i < qtd_a; i++{
//         prox := i+1
//         if figus[i+1] != prox{
//             falta = append(falta, prox)
//         }
//    }

//    fmt.Print(falta)
   
}
