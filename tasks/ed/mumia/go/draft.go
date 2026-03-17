package main
import "fmt"
func main() {
    var name, classif string
    var age int

    fmt.Scan(&name)
    fmt.Scan(&age)

            if (age < 12){
                classif = "crianca"
            }else if (age < 18 && age >= 12){
                classif = "jovem"
            }else if (age >= 18 && age < 65){
                classif = "adulto"
            }else if (age >= 65 && age < 1000){
                classif = "idoso"
            }else{
                classif = "mumia"
            }   

        fmt.Printf("%s eh %s\n", name, classif)

        }


