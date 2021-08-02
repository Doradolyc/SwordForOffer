package main


type Stack struct {
	i 	 int
	data [10000]int
}

func (s *Stack) Push(k int)  {
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() (ret int) {
	s.i--
	ret = s.data[s.i]
	return
}

func (s *Stack) Empty() bool {
	return s.i == 0
}

// push  1 2 3 4 5
// queue 1 2 3 4 5

// push
// stack1 5 4 3 2 1

// pop
// stack2 1 2 3 4 5


type CQueue struct {
	s1 Stack
	s2 Stack
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.s1.Push(value)
}


// 时间 o(n)
// 控件o(n)
func (this *CQueue) DeleteHead() int {
	if this.s1.Empty() {
		return -1
	}

	for !this.s1.Empty() {
		this.s2.Push(this.s1.Pop())
	}

	head := this.s2.Pop()

	for !this.s2.Empty() {
		this.s1.Push(this.s2.Pop())
	}

	return head

}
