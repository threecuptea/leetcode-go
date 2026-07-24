// https://leetcode.com/problems/print-in-order/description
type Foo struct {
    firstChan chan struct{}
    secondChan chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
        firstChan: make(chan struct{}),
        secondChan: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
    defer close(f.firstChan)
	printFirst()
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
    defer close(f.secondChan)
    <- f.firstChan
	printSecond()
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
    <-f.secondChan
	printThird()
}