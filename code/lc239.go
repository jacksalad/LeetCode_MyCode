/*
239. 滑动窗口最大值 : https://leetcode.cn/problems/sliding-window-maximum/
	手写双向链表
    单调队列算法
*/

// 链表节点
type Node struct {
	Idx   int
	Val   int
	Left  *Node
	Right *Node
}

// 双向链表
type Deque struct {
	First *Node
	Last  *Node
}

// 判断队空
func (this Deque) IsEmpty() bool {
	return this.First == nil
}

// 队首值
func (this Deque) getFirst() int {
	return this.First.Val
}

// 队尾值
func (this Deque) getLast() int {
	return this.Last.Val
}

// 队首下标
func (this Deque) FirstIndex() int {
	return this.First.Idx
}

// 队尾下标
func (this Deque) LastIndex() int {
	return this.Last.Idx
}

// 队首入队
func (this *Deque) OfferFirst(idx int, val int) bool {
	if this.IsEmpty() {
		this.First = &Node{Idx: idx, Val: val}
		this.Last = this.First
		return true
	}
	this.First.Left = &Node{Idx: idx, Val: val, Right: this.First}
    this.First = this.First.Left
	return true
}

// 队尾入队
func (this *Deque) OfferLast(idx int, val int) bool {
	if this.IsEmpty() {
		this.First = &Node{Idx: idx, Val: val}
		this.Last = this.First
		return true
	}
	this.Last.Right = &Node{Idx: idx, Val: val, Left: this.Last}
    this.Last = this.Last.Right
	return true
}

// 队首出队
func (this *Deque) PollFirst() *Node {
	if this.IsEmpty() {
		return nil
	}
	res := this.First
	if this.First == this.Last {
		this.First = nil
		this.Last = nil
		return res
	}
	this.First = this.First.Right
	this.First.Left = nil
	return res
}

// 队尾出队
func (this *Deque) PollLast() *Node {
	if this.IsEmpty() {
		return nil
	}
	res := this.Last
	if this.First == this.Last {
		this.First = nil
		this.Last = nil
		return res
	}
	this.Last = this.Last.Left
	this.Last.Right = nil
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	var ql Deque
	for i, v := range nums {
        // 维护窗口大小
		for !ql.IsEmpty() && i-ql.FirstIndex() >= k {
			ql.PollFirst()
		}
        // 维护单调队列
		for !ql.IsEmpty() && ql.getLast() < v {
			ql.PollLast()
		}
		ql.OfferLast(i, v)
		if i-k+1 >= 0 {
			ans[i-k+1] = ql.getFirst()
		}
	}
	return ans
}
