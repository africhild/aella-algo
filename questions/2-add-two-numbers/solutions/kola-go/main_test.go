// write test for addTwoNumbers

package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	result := addTwoNumbers(l1, l2)
	if result.Val != 7 {
		t.Errorf("expected 7, got %d", result.Val)
	}
	if result.Next.Val != 0 {
		t.Errorf("expected 0, got %d", result.Next.Val)
	}
	if result.Next.Next.Val != 8 {
		t.Errorf("expected 8, got %d", result.Next.Next.Val)
	}
}