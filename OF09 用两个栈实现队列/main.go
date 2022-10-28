package main

type CQueue struct {
	inStark, outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStark = append(this.inStark, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStark) == 0 {
			return -1
		}
		this.in2on()
	}
	value := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return value
}

func (this *CQueue) in2on() {
	for len(this.inStark) > 0 {
		this.outStack = append(this.outStack, this.inStark[len(this.inStark)-1])
		this.inStark = this.inStark[:len(this.inStark)-1]
	}
}

func main() {

}
