package main

import (
	"fmt"
	// "strings"
)


func valido(ss []byte, pos int, num byte, l int) bool  {
    for i := pos -1; i >= pos - l; i--{

        if i < 0{
            break
        }

        if ss[i] == num{
            return false
        }
    }

    for i := pos +1; i <= pos + l; i++{
        if i >= len(ss){
            break
        }

        if ss[i] == num{
            return false
        }
    }

    return true
}

func back(ss []byte, is []int, p int, l int) bool{
    if p == len(is){
        return true
    }

    posS := is[p]

    for i:= 0; i <= l; i++{
        //dá um jeito de fazer a conversão do numero em byte
        digB := byte(i + '0')

        if valido(ss, posS, digB, l){
            ss[posS] = digB

            if back(ss, is, p+1, l){
                return true
            }
        }

        ss[posS] = '.'
    }

    return false
}

func main() {
    var sequenciaS string
    var l int
    var is []int
    fmt.Scan(&sequenciaS)
    fmt.Scan(&l)

    sequencia := []byte(sequenciaS)

    for i:=0; i < len(sequencia); i++{
        if sequencia[i] == '.'{
            is = append(is, i)
        }
    }

    if back(sequencia, is, 0, l){
        fmt.Println(string(sequencia))
    }
}