package main

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    l3 := &ListNode{0, nil}
    carry := 0

    currA := l1
    currB := l2
    currC := l3

    for currA != nil || currB != nil || carry > 0 {
        c := carry

        if currA != nil {
            c += currA.Val
            currA = currA.Next
        }
        if currB != nil {
            c += currB.Val
            currB = currB.Next
        }

        carry = c / 10
        currC.Val = c % 10

        if carry > 0 || currA != nil || currB != nil {
            currC.Next = &ListNode{0,nil}
            currC = currC.Next
        }
    }

    return l3
}

func ArrayToListNode(A []int) *ListNode {
    list := &ListNode{0, nil}
    curr := list

    for i := 0; i < len(A); i++ {
        curr.Val = A[i]
        if i + 1 < len(A) {
            curr.Next = &ListNode{0,nil}
            curr = curr.Next
        }
    }

    return list
}

func PrintList(l *ListNode) {
    for c := l; c != nil; c = c.Next {
        print(c.Val)
    }
    println()
}

func main() {
    a := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
    b := ListNode{4, &ListNode{5, &ListNode{6, nil}}}
    PrintList(&a)
    PrintList(&b)
    PrintList(addTwoNumbers(&a, &b))
    println()
    

    a = ListNode{1, &ListNode{2, &ListNode{3, nil}}}
    b = ListNode{4, &ListNode{5, nil}}
    PrintList(&a)
    PrintList(&b)
    PrintList(addTwoNumbers(&a, &b))
    println()
    
    a = ListNode{1, &ListNode{1, nil}}
    b = ListNode{1, &ListNode{2, &ListNode{3, nil}}}
    PrintList(&a)
    PrintList(&b)
    PrintList(addTwoNumbers(&a, &b))
    println()
    
    println("nil")
    println("nil")
    PrintList(addTwoNumbers(nil, nil))
    println()
    
    a = *ArrayToListNode([]int{9,9,9,9})
    b = *ArrayToListNode([]int{9,9,9,9,9,9})
    PrintList(&a)
    PrintList(&b)
    PrintList(addTwoNumbers(&a, &b))
    println()
}