// https://leetcode.com/problems/print-zero-even-odd/description/
type ZeroEvenOdd struct {
	n        int
    zeroChan chan struct{}
    oddChan  chan struct{}
    evenChan chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
    zeroChan := make(chan struct{}, 1); zeroChan <- struct{}{}
    oddChan := make(chan struct{})
    evenChan := make(chan struct{})
	return &ZeroEvenOdd{
		n:          n,
        zeroChan:   zeroChan,
        oddChan:    oddChan,
        evenChan:   evenChan,
	}
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
    for i := 1; i <= z.n; i++ {
        <- z.zeroChan
        printNumber(0)
        if (i % 2 == 0) {
            z.evenChan <- struct{}{}
        } else {
            z.oddChan <- struct{}{}
        }
    }
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
    for i := 2; i <= z.n; i += 2 {
        <- z.evenChan
        printNumber(i)
        z.zeroChan <- struct{}{}
    }
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
    for i := 1; i <= z.n; i += 2 {
        <- z.oddChan
        printNumber(i)
        z.zeroChan <- struct{}{}
    }
}
