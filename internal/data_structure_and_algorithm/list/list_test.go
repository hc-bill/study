package list

import (
	"testing"
)

func TestMainer(t *testing.T) {
	head := &ListNode{
		Data: 1,
		Next: &ListNode{
			Data: 2,
			Next: &ListNode{Data: 3},
		},
	}

	head.Traverse()

	d := &ListNode{
		Data: 3,
		Next: nil,
	}
	head.Delete(d)
	head.Traverse()
}
