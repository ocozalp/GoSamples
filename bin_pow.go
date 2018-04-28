package main

import "fmt"

func binPow(base, power int) int {
    b := base
    p := power
    result := 1
    for p > 0 {
        if p % 2 == 1 {
            result = result * b
        }

        p = p/2
        b = b * b
    }

    return result
}

func binPowRecursive(base, power int) int {
    if power == 0 {
        return 1
    }
    
    elm := binPowRecursive(base, power/2)

    if power % 2 == 1 {
        return elm * elm * base
    }

    return elm * elm
}

func main() {
    fmt.Println(binPow(2, 15))
    fmt.Println(binPowRecursive(2, 15))
}
