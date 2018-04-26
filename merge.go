package main

import "fmt"

func main() {
    var slice1 = []int{1, 3, 5, 9, 11}
    var slice2 = []int{0, 2, 4, 6, 8}
    
    var result = merge(slice1, slice2)

    fmt.Println(result)
}

func merge(a1, a2 []int) []int {
    var len1 = len(a1)
    var len2 = len(a2)
    var result = make([]int, len1+len2, len1+len2)
 
    var i1, i2 int

    for i1 < len1 || i2 < len2 {
        if i1 == len1 || (i2 < len2 && a2[i2] < a1[i1]) {
            result[i1+i2] = a2[i2]
            i2++
        } else {
            result[i1+i2] = a1[i1]
            i1++
        }
    }

    return result
}
