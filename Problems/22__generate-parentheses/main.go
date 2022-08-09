package main

import (
	"container/list"
	"fmt"
)

// Performance: O(2^(2n))
// Space: O(2^(2n))
// Runtime (n=13) 5.965s
// Runtime (n=14) 17.054s
func generateParenthesisRecursiveV1(n int) {
    maxSize := n*2
    var gp func(string,int,int)
    gp = func(parens string, left int, right int) {
        if left+right >= maxSize {
            fmt.Println(parens)
            return
        }
        if left < n {
            gp(parens+"(", left+1, right)
        }
        if right < left {
            gp(parens+")", left, right+1)
        }
    }

    gp("", 0, 0)
}

// Performance: O(2^(2n))
// Space: O(n)
// Runtime (n=13) 3.995s
// Runtime (n=14) 20.135s
func generateParenthesisRecursiveV2(n int) {
    maxSize := n*2
    parens := make([]byte, maxSize)

    var gp func(int,int)
    gp = func(left int, right int) {
        if left+right >= maxSize {
            fmt.Println(string(parens))
            return
        }
        if left < n {
            parens[left+right] = '('
            gp(left+1, right)
        }
        if right < left {
            parens[left+right] = ')'
            gp(left, right+1)
        }
    }

    gp(0, 0)
}


type Triple struct {
    Parens string
    Left int
    Right int
}

// Performance: O(2^(2n))
// Space: O(2^(2n))
// Runtime (n=13) 5.373s
// Runtime (n=14) 14.344s
func generateParenthesisIterative1(n int) {
    maxSize := n*2
    
    stack := list.New()
    stack.PushBack(Triple{"", 0, 0})

    for stack.Len() > 0 {
        currRef := stack.Back()
        stack.Remove(currRef)
        curr := currRef.Value.(Triple)

        if  (curr.Left + curr.Right) >= maxSize {
            fmt.Println(curr.Parens)
            continue
        }
        if curr.Left < n {
            stack.PushBack(Triple{curr.Parens+"(", curr.Left+1, curr.Right})
        }
        if curr.Right < curr.Left {
            stack.PushBack(Triple{curr.Parens+")", curr.Left, curr.Right+1})
        }
    }
}

// Performance: O(2^(2n))
// Space: O(2^(2n))
// Runtime (n=13) 5.729s
// Runtime (n=14) 20.931s
func generateParenthesisIterative2(n int) {
    maxSize := n*2
    
    stack := []Triple{}
    stack = append(stack, Triple{"", 0, 0})

    for len(stack) > 0 {
        curr := stack[len(stack) - 1]
        stack = stack [:len(stack)-1]


        if  (curr.Left + curr.Right) >= maxSize {
            fmt.Println(curr.Parens)
            continue
        }
        if curr.Left < n {
            stack = append(stack, Triple{curr.Parens+"(", curr.Left+1, curr.Right})
        }
        if curr.Right < curr.Left {
            stack = append(stack, Triple{curr.Parens+")", curr.Left, curr.Right+1})
        }
    }
}


func generateParenthesisIterative3(n int) []string {
    maxSize := n*2
    parensList := []string{}
    
    stack := list.New()
    stack.PushBack(Triple{"", 0, 0})

    for stack.Len() > 0 {
        currRef := stack.Back()
        stack.Remove(currRef)
        curr := currRef.Value.(Triple)

        if  (curr.Left + curr.Right) >= maxSize {
            parensList = append(parensList, curr.Parens)
            continue
        }
        if curr.Left < n {
            stack.PushBack(Triple{curr.Parens+"(", curr.Left+1, curr.Right})
        }
        if curr.Right < curr.Left {
            stack.PushBack(Triple{curr.Parens+")", curr.Left, curr.Right+1})
        }
    }

    return parensList
}

func main() {
    fmt.Println(generateParenthesisIterative3(4))
}