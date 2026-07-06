package main

import (
	//"container/list"
	"fmt"
)

var n, soma int
var lista []int

func back(index, somaA int) bool{
    if somaA == soma{
        return true
    }

    if  index == n || somaA > soma{
        return false
    }

    if back(index+1, somaA+lista[index]){
        return true
    }

    return back(index+1, somaA)
}

func main() {
    fmt.Scan(&n, &soma)
    for i:=0; i < n; i++{
        var x int
        fmt.Scan(&x)
        lista = append(lista, x)
    }

    if back(0,0){
        fmt.Println(true)
    }else{
        fmt.Println(false)
    }
}