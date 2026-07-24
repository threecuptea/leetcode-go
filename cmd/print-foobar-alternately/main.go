// https://leetcode.com/problems/print-foobar-alternately/description/
type FooBar struct {
	n int
    foo chan struct{}
    bar chan struct{}
}

func NewFooBar(n int) *FooBar {
    // make(chan struct{}) unbuffered channel requires synchronized operation. It will lead to timeout (the first line)
    foo := make(chan struct{}, 1); foo <- struct{}{}
    bar := make(chan struct{}) // This can leave unbuffered
	return &FooBar {
        n: n,
        foo: foo,
        bar: bar,
        }
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
        <- fb.foo
        printFoo()
        fb.bar <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
        <- fb.bar
        printBar()
        fb.foo <- struct{}{}
	}
}