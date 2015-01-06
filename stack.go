package suitecase

type stack []rune

func (this *stack) push(chr rune) {
	*this = append(*this, chr)
}
func (this *stack) top() (r rune) {
	if this.isEmpty() {
		return
	}
	return (*this)[this.depth()-1]
}
func (this *stack) pop() (r rune) {
	if this.isEmpty() {
		return
	}
	r = this.top()
	*this = (*this)[:this.depth()-1]
	return
}
func (this *stack) clear() {
	*this = make([]rune, 0)
}
func (this *stack) isEmpty() bool {
	return this.depth() == 0
}
func (this *stack) depth() int {
	return len(*this)
}
func (this *stack) String() (r string) {
	return string(*this)
}
