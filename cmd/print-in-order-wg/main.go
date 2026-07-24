// https://leetcode.com/problems/print-in-order/description
import "sync"
type Foo struct {
    wg1 sync.WaitGroup // The gate for Second
    wg2 sync.WaitGroup // The gate for Third
}

func NewFoo() *Foo {
    f := &Foo{}
    f.wg1.Add(1)
    f.wg2.Add(1)
    return f
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
    defer f.wg1.Done()
	printFirst()
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
    defer f.wg2.Done()
    f.wg1.Wait()
	printSecond()
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
    f.wg2.Wait()
	printThird()
}