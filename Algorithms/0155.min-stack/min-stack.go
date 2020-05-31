package problem0155

// MinStack1 是可以返回最小值的栈
type MinStack1 struct {
	stack []item
}
type item struct {
	min, x int
}

// Constructor 构造 MinStack
func Constructor1() MinStack1 {
	return MinStack1{}
}

func (this *MinStack1) Push(x int) {
	var min int
	if len(this.stack)==0{
		this.stack=append(this.stack,item{x,x})
	}else{
		min=this.stack[len(this.stack)-1].min
		if x<min {
			this.stack=append(this.stack,item{x,x})
		}else{
			this.stack=append(this.stack,item{min,x})
		}
	}
}

func (this *MinStack1) Pop(){
	this.stack=this.stack[:len(this.stack)-1]
}

func (this *MinStack1) Top() int {
	return this.stack[len(this.stack)-1].x
}

func (this *MinStack1) GetMin() int {
	return this.stack[len(this.stack)-1].min
}


type MinStack struct {
	stack []int
	min int
}

// Constructor 构造 MinStack
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack)==0{
		this.stack=append(this.stack,x)
		this.min=x
		return
	}
	if x<=this.min{
		this.stack=append(this.stack,this.min,x)
		this.min=x
		return
	}
	this.stack=append(this.stack,x)
}

func (this *MinStack) Pop(){
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if val == this.min && len(this.stack) != 0 {
		this.min = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		return
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}