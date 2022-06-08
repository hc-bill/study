package list

/*
	链表
		头节点: 指的是链表第一个节点
		头指针: 指的是指向头节点的指针
*/

type ListNode struct {
	Data int
	Next *ListNode
}

// Traverse 遍历
func (ln *ListNode) Traverse() {
	for mln := ln; mln != nil; mln = mln.Next {
		println(mln.Data)
	}
}

// FindValue 查询链表中的值
func (ln *ListNode) FindValue(value int) *ListNode {
	mln := ln
	for mln != nil {
		if mln.Data == value {
			return mln
		}
		mln = mln.Next
	}
	return nil
}

// HeadInsert 头插
func (ln *ListNode) HeadInsert(value int) {
	newNode := &ListNode{value, nil}
	// 先将新节点的 Next 连到链表上,防止断链,再进行插入操作
	newNode.Next = ln
	ln = newNode
}

// TailInsert 尾插
func (ln *ListNode) TailInsert(value int) {
	newNode := &ListNode{value, nil}
	if ln == nil {
		ln.Next = newNode
	} else {
		mln := ln
		// find tail node
		for mln.Next != nil {
			mln = mln.Next
		}
		mln.Next = newNode
	}
}

// 尾部结点
var tail *ListNode = nil

// WithTailNodeOfTailInsert 带尾部节点的尾插
func (ln *ListNode) WithTailNodeOfTailInsert(value int) {
	newNode := &ListNode{value, nil}
	if ln == nil {
		ln.Next = newNode
	} else {
		tail.Next = newNode
	}
}

func (ln *ListNode) LocationInsert(p *ListNode, value int) {
	if p == nil || ln == nil {
		return
	} else {
		mln := p

		newNode := &ListNode{value, nil}
		newNode.Next = mln.Next
		p.Next = newNode
	}
}

// Delete 删除给定结点
func (ln *ListNode) Delete(p *ListNode) {
	if p == nil || ln == nil {
		return
	}

	var pre *ListNode = nil
	mln := ln
	for mln != nil {
		if mln == p { // 找到要删除的节点
			break
		}

		pre = mln // 记录要删除节点的前驱结点
		mln = mln.Next
	}

	if mln == nil {
		return
	}
	if pre == nil {
		ln = ln.Next
	} else {
		pre.Next = pre.Next.Next
	}
}

// DeleteNextNode 删除给定结点之后的结点
func (ln *ListNode) DeleteNextNode(p *ListNode) {
	if p == nil || p.Next == nil {
		return
	}
	p.Next = p.Next.Next
}

// Reverse 反转链表
func (head *ListNode) Reverse() {
	if head == nil {
		return
	}

	var r *ListNode
	newVirtualHeadListNode := &ListNode{Data: 0, Next: nil}
	pre := head
	head = newVirtualHeadListNode

	for pre != nil {
		r = pre
		pre = pre.Next
		r.Next = newVirtualHeadListNode.Next
		newVirtualHeadListNode.Next = pre
	}
}
