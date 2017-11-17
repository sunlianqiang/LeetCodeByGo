package _232_Implement_Queue_using_Stacks

type MyStack struct {
	val []int
}

func ConstructorStack() MyStack  {
	return MyStack{[]int{}}
}

func (this *MyStack) Push(x int) {
	this.val = append([]int{x}, this.val...)
}

func (this *MyStack) Pop() int {
	tmp := this.val[0]
	this.val = this.val[1:]
	return tmp
}

func (this *MyStack) Peek() int {
	return this.val[0]
}
func (this *MyStack) Size() int {
	return len(this.val)
}

func (this *MyStack) Empty() bool {
	if 0 == len(this.val) {
		return true
	}

	return false
}

type MyQueue struct {
	stack1 MyStack
	stack2 MyStack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	var queue MyQueue
	queue.stack1 = ConstructorStack()
	queue.stack2 = ConstructorStack()

	return queue
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stack1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if !this.stack2.Empty() {
		return this.stack2.Pop()
	} else {
		for ;!this.stack1.Empty(); {
			tmp := this.stack1.Pop()
			this.stack2.Push(tmp)
		}
		return this.stack2.Pop()
	}
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.stack2.Empty() {
		return this.stack2.Peek()
	} else {
		for ;!this.stack1.Empty(); {
			tmp := this.stack1.Pop()
			this.stack2.Push(tmp)
		}
		return this.stack2.Peek()
	}
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack1.Empty() && this.stack2.Empty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
