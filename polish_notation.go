package main

import "fmt"
import s "strings"
import sc "strconv"

func main() {
    var expression = "* + 2 3 4"
    result := evaluate(expression)

    fmt.Println(result.eval())    
}

func evaluate(result string) Expression {
    var expressions = make([]Expression, 0)
    var operators = make([]string, 0)

    for _, token := range s.Split(result, " ") {
        if token == "+" || token == "*" {
            operators = append(operators, token)
        } else {
            v, _ := sc.Atoi(token)
            expressions = append(expressions, ConstantExpression{value:v})
        }
    }

    current := 0

    for i := len(operators) - 1; i >= 0; i-- {
        exp := operators[i]
        op1 := expressions[current]
        current++
        op2 := expressions[current]

        if exp == "+" {
            new_exp := Sum{BinaryOperator{left:op1, right:op2}}
            expressions[current] = new_exp
        } else {
            new_exp := Multiply{BinaryOperator{left:op1, right:op2}}
            expressions[current] = new_exp
        }
    }

    return expressions[current] 
}

type Expression interface {
    eval() int
}

// Binary Operators
type BinaryOperator struct {
    left, right Expression
}

type Sum struct {
    oper BinaryOperator
}

type Multiply struct{
    oper BinaryOperator
}

func (s Sum) eval() int {
    return s.oper.left.eval() + s.oper.right.eval()
}

func (m Multiply) eval() int {
    return m.oper.left.eval() * m.oper.right.eval()
}

// Constant Expression
type ConstantExpression struct {
    value int
}

func (ex ConstantExpression) eval() int {
    return ex.value
}
