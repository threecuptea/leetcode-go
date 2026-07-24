// https://leetcode.com/problems/building-h2o/description/
type H2O struct {
    hydChan chan func()
    oxyChan chan func()
    done chan struct{}
}

func NewH2O() *H2O {
	h := &H2O{
        hydChan: make(chan func()),
        oxyChan: make(chan func()),
        done: make(chan struct{}),
    }
    go h.makeH2O()
	return h
}

func (h *H2O) makeH2O() {
    for {
        releaseHydrogen1 := <-h.hydChan
        releaseHydrogen1()
        releaseHydrogen2 := <-h.hydChan
        releaseHydrogen2()
        releaseOxygen := <-h.oxyChan
        releaseOxygen()

        h.done <- struct{}{}
        h.done <- struct{}{}
        h.done <- struct{}{}
    }
}

// "HOH", "OOHHHHH"
func (h *H2O) Hydrogen(releaseHydrogen func()) {
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	h.hydChan <- releaseHydrogen
    <- h.done
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	// releaseOxygen() outputs "O". Do not change or remove this line.
	h.oxyChan <- releaseOxygen
    <- h.done
}