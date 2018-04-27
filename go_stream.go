package main 

import "fmt"
import "reflect"

type _forEach func(interface{})

type JStream struct {
    data []interface{} 
}

func (stream JStream) forEach(f _forEach) {
    for d := range stream.data {
        f(d)
    }
}

func print(v interface{}) {
    fmt.Println(v)
}

func main() {
    arr := []int{1,2,3,4,5}

    stream, success := of(arr)

    if !success {
        print("Something happened!")
        return
    }

    stream.forEach(print)
}

func of(data interface{}) (res JStream, success bool) {
    t := reflect.ValueOf(data).Kind()

    if t != reflect.Slice {
        success = false
        return
    }

    slice := reflect.ValueOf(data)
    result := make([]interface{}, slice.Len(), slice.Len())
    
    for i := 0; i < slice.Len(); i++ {
        result[i] = slice.Index(i).Interface()
    }

    return JStream{data:result}, true
}
 
