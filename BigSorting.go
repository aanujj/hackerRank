package main

import (
    "fmt"
    "math/big"
)

func sort(a []float64) []float64 {

    for i := 0; i < len(a); i = i + 1 {

        for j := i + 1; j < len(a); j++ {
            //tmp := 0
           
            if a[i] > a[j] {

                a[i], a[j] = a[j], a[i]
            }

        }

    }

    return a
}

func main() {
    var test int
    fmt.Scanln(&test)
    var a float64

    slice := make([]float64, 0)
    for i := 0; i < test; i++ {
        fmt.Scanln(&a)
        big.NewFloat(a)
        slice = append(slice, a)

    }
    sortedArray := sort(slice)
    for _, v := range sortedArray {
        fmt.Println(v)
    }
}
