package main


// func main() {

// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers( l1 *ListNode, l2 *ListNode)  *ListNode {
	var head *ListNode
	var tail *ListNode
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		sum = sum % 10
		node := &ListNode{Val: sum}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}



// what are the zero values for a struct and a pointer to a struct?
// what is the zero value for a slice?
// what is the zero value for a map?
// what is the zero value for a function?
// what is the zero value for a channel?
// what is the zero value for an interface?
// what is the zero value for a string?
// what is the zero value for a bool?
// what is the zero value for a int?
// what is the zero value for a float?
// what is the zero value for a complex?
// answer:
// struct: {}
// pointer to struct: nil
// slice: nil
// map: nil
// function: nil
// channel: nil
// interface: nil
// string: ""
// bool: false
// int: 0
// float: 0
// complex: 0
